// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//go:generate go run ../genzfunc.go

package selection

import (
	"sort"

	"github.com/cnotch/algo/sort/quick/partition"
)

// Interface 标准库 sort.Interface 别名
type Interface = sort.Interface

// LessSwap partition.LessSwap 别名
type LessSwap = partition.LessSwap

// Sort 对 data 执行选择排序
func Sort(data Interface) {
	selectionSort(data, 0, data.Len())
}

// Get .
func Get() func(data Interface, lo, hi int) {
	return selectionSort
}

// 依次选择最小值
func selectionSort(data Interface, lo, hi int) {
	for i := lo; i < hi-1; i++ {
		min := i
		for j := i + 1; j < hi; j++ {
			if data.Less(j, min) {
				min = j
			}
		}
		// 将最小值交换到i
		data.Swap(i, min)
	}
}
