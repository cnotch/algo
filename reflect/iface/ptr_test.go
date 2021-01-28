// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package iface_test

import (
	"testing"
	"unsafe"

	"github.com/cnotch/algo/reflect/iface"
	"github.com/stretchr/testify/assert"
)

func TestNumberPtr(t *testing.T) {
	a := 3
	aptr := unsafe.Pointer(&a) // 原值的指针
	iptr := iface.Ptr(a)       // 被封装后复制品的指针
	iptr2 := iface.Ptr(&a)

	assert.NotEqual(t, aptr, iptr)
	assert.Equal(t, aptr, iptr2)

	// 对于引用类型，不重新分配内存的情况下的修改相互影响
	a += 20
	assert.Equal(t, 3, *(*int)(iptr))
	assert.Equal(t, 23, *(*int)(iptr2))
}

func TestMapPtr(t *testing.T) {
	a := map[int]string{1: "a", 2: "b"}
	aptr := *((*unsafe.Pointer)(unsafe.Pointer(&a))) // 原值的指针
	iptr := iface.Ptr(a)

	assert.Equal(t, aptr, iptr)

	// 对于引用类型，不重新分配内存的情况下的修改相互影响
	a[1] = "change"
	m1 := *(*map[int]string)(unsafe.Pointer(&iptr))
	assert.Equal(t, "change", m1[1])
}

func TestSlicePtr(t *testing.T) {
	a := []int{1,2,3,4}
	aptr := unsafe.Pointer(&a) // 原值的指针
	iptr := iface.Ptr(a)
	iptr2:=iface.Ptr(&a)
	
	assert.NotEqual(t, aptr, iptr)
	assert.Equal(t, aptr, iptr2)

	// 对于引用类型，不重新分配内存的情况下的修改相互影响
	a[0] =100
	m1 := *(*[]int)(iptr)
	assert.Equal(t, 100, m1[0])
}
