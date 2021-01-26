// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package merge

import "reflect"

// Slice 切片排序
func Slice(slice interface{}, less func(i, j int) bool) {
	rv := reflect.ValueOf(slice)
	swap := reflect.Swapper(slice)
	length := rv.Len()
	mergeSortZfunc(LessSwap{Less: less, Swap: swap}, 0, length)
}
