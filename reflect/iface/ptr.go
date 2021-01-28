// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package iface

import (
	"unsafe"
)

// interface{} 类型的替身
type interfaceSubstitute struct {
	typ  *struct{}      // 接口实际类型的占位符
	word unsafe.Pointer // 接口指向的真实类型指针
}

// Ptr 获得 i 指向的指针。
// 有如下两种可能：
// 	1. i 是指针类型返回 i 本身所代表的指针(Chan Func Map  Ptr UnsafePointer)
//  2. i 是 slice 返回其复制品 SliceHeader 指针
//  3. i 是 string 返回其复制品 StringHeader 指针
// 	4. 其他，返回其复制品所在地址的指针
func Ptr(i interface{}) unsafe.Pointer {
	if i == nil {
		return nil
	}
	is := *(*interfaceSubstitute)(unsafe.Pointer(&i))
	return is.word
}
