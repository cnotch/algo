// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package iface

import (
	"reflect"
	"unsafe"
)

// Assigner 通过反射获取 Slice 元素间赋值函数。
// 	left[l] = right[r]
func Assigner(left, right interface{}) func(l, r int) {
	lv := reflect.ValueOf(left)
	if lv.Kind() != reflect.Slice {
		panic("Left isn't slice")
	}

	rv := reflect.ValueOf(right)
	if rv.Kind() != reflect.Slice {
		panic("Right isn't slicee")
	}

	ltyp := lv.Type().Elem()
	rtyp := rv.Type().Elem()
	if ltyp != rtyp {
		panic(ltyp.String() + " != " + rtyp.String())
	}

	// 切片元素大小
	size := int(ltyp.Size())

	// slice 指针
	lptr := Ptr(left)
	rptr := Ptr(right)

	// 使用 Copy 模式
	rslice := sliceToBytes(rptr, size)
	lslice := sliceToBytes(lptr, size)
	return func(l, r int) {
		loffset := l * size
		roffset := r * size
		copy(lslice[loffset:loffset+size], rslice[roffset:roffset+size])
	}
}

func sliceToBytes(ptr unsafe.Pointer, elemSize int) []byte {
	var bytes []byte // 替身
	bytesHeader := (*reflect.SliceHeader)(unsafe.Pointer(&bytes))
	header := (*reflect.SliceHeader)(ptr)

	bytesHeader.Data = header.Data
	bytesHeader.Len = elemSize * header.Len
	bytesHeader.Cap = elemSize * header.Cap

	return bytes
}
