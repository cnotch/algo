// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package iface_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/cnotch/algo/reflect/iface"
	"github.com/stretchr/testify/assert"
)

const size = 64

func BenchmarkAssignerMode(b *testing.B) {
	a1 := [][size]byte{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}}
	a2 := [][size]byte{{4, 5, 6}, {4, 5, 6}, {4, 5, 6}}

	// 使用赋值模式
	f := func(lptr, rptr unsafe.Pointer) func(l, r int) {
		left := *(*[][size]byte)(lptr)
		right := *(*[][size]byte)(rptr)
		return func(l, r int) { left[l] = right[r] }
	}
	assigner := f(unsafe.Pointer(&a2), unsafe.Pointer(&a1))
	for i := 0; i < b.N; i++ {
		assigner(0, 0)
	}
}

func BenchmarkCopyMode(b *testing.B) {
	a1 := [][size]byte{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}}
	a2 := [][size]byte{{4, 5, 6}, {4, 5, 6}, {4, 5, 6}}

	// 使用 Copy 模式
	var right, left []byte // 替身
	rightHeader := (*reflect.SliceHeader)(unsafe.Pointer(&right))
	leftHeader := (*reflect.SliceHeader)(unsafe.Pointer(&left))
	rheader := (*reflect.SliceHeader)(unsafe.Pointer(&a1))
	lheader := (*reflect.SliceHeader)(unsafe.Pointer(&a2))

	rightHeader.Data = rheader.Data
	rightHeader.Len = size * rheader.Len
	rightHeader.Cap = size * rheader.Cap

	leftHeader.Data = lheader.Data
	leftHeader.Len = size * lheader.Len
	leftHeader.Cap = size * lheader.Cap

	assigner := func(l, r int) {
		loffset := l * size
		roffset := r * size
		copy(left[loffset:loffset+size], right[roffset:roffset+size])
	}

	for i := 0; i < b.N; i++ {
		assigner(0, 0)
	}
}

func TestIntAssigner(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	assigner := iface.Assigner(b, a)
	assigner(0, 0)

	assert.Equal(t, a[0], b[0])
}

func TestArray(t *testing.T) {
	a := [][110]byte{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}}
	b := [][110]byte{{4, 5, 6}, {4, 5, 6}, {4, 5, 6}}
	assigner := iface.Assigner(b, a)
	assigner(0, 0)
	
	assert.Equal(t, a[0], b[0])
}

func TestSliceAssigner(t *testing.T) {
	a := [][]int{{1, 2, 3}, {4, 5, 6}}
	b := [][]int{{10, 20, 30}, {40, 50, 60}}
	assigner := iface.Assigner(b, a)
	assigner(0, 0)

	assert.Equal(t, a[0], b[0])
}

func TestStringAssigner(t *testing.T) {
	a := []string{"123", "456"}
	b := []string{"102030", "405060"}
	assigner := iface.Assigner(b, a)
	assigner(0, 0)

	assert.Equal(t, a[0], b[0])
}

func TestInterfaceAssigner(t *testing.T) {
	a := []interface{}{"123", "456"}
	b := []interface{}{"102030", "405060"}
	assigner := iface.Assigner(b, a)
	assigner(0, 0)

	assert.Equal(t, a[0], b[0])
}

func TestUintptrAssigner(t *testing.T) {
	a := []uintptr{uintptr(123), uintptr(456)}
	b := []uintptr{uintptr(102030), uintptr(405060)}
	assigner := iface.Assigner(b, a)
	assigner(0, 0)

	assert.Equal(t, a[0], b[0])
}

type sv struct {
	i  int
	s  string
	bs []byte
}

func TestStructAssigner(t *testing.T) {
	a := []sv{{1, "124", []byte{1, 2, 3}}, {2, "234", []byte{2, 3, 4}}}
	b := []sv{{10, "102040", []byte{10, 20, 30}}, {20, "203040", []byte{20, 30, 40}}}
	assigner := iface.Assigner(b, a)
	assigner(0, 0)

	assert.Equal(t, a[0], b[0])
}
