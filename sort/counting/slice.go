// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package counting

import (
	"reflect"

	"github.com/cnotch/algo/reflect/iface"
)

// Slice 对 Slice 执行计数排序。
// 	slice 待排序的集合；
//  k 计数参考值范围 [0,k]；
//	key 计算集合中元素计数参照值函数，返回值必须在 [0,k]；
func Slice(slice interface{}, k int, key func(i int) int) {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		panic("slice isn't slice type")
	}

	n := rv.Len()
	temp := reflect.MakeSlice(rv.Type(), n, n)
	assign := iface.Assigner(temp.Interface(), slice)

	countingSort(n, k, key, assign)

	// 保证 slice 已排序
	reflect.Copy(rv, temp)
}
