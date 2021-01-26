// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package iface

import "unsafe"

// interface{} 类型的替身
type interfaceSubstitute struct {
	typ  *struct{}      // 接口实际类型的占位符
	word unsafe.Pointer // 接口指向的真实类型指针
}

// Ptr i 指向的真实类型指针。
// 注意：原值的指针不等于复制品的指针。
func Ptr(i interface{}) unsafe.Pointer {
	is := *(*interfaceSubstitute)(unsafe.Pointer(&i))
	return is.word
}
