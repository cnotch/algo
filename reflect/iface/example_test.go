// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package iface_test

import (
	"fmt"
	"unsafe"

	"github.com/cnotch/algo/reflect/iface"
)

func ExamplePtr() {
	a := []int{3, 4}
	aptr := unsafe.Pointer(&a) // 原值的指针
	iptr := iface.Ptr(a)       // 被封装后复制品的指针

	fmt.Println(aptr != iptr)

	// 对于引用类型，不重新分配内存的情况下的修改相互影响
	a[0] += 2
	got := (*(*[]int)(iptr))[0]
	fmt.Println(got)

	// Output:
	// true
	// 5
}
