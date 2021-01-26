// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//go:generate go run ../genzfunc.go

package bubble

import (
	"sort"

	"github.com/cnotch/algo/sort/quick/partition"
)

// Interface 标准库 sort.Interface 别名
type Interface = sort.Interface

// LessSwap partition.LessSwap 别名
type LessSwap = partition.LessSwap

// Sort 对 data 执行冒泡排序
func Sort(data Interface) {
	bubbleSort(data, 0, data.Len())
}

// Get .
func Get() func(data Interface, lo, hi int) {
	return bubbleSort
}

// Bubble Sort 对集合索引范围[lo,hi)的元素排序
func bubbleSort(data Interface, lo, hi int) {
	for i := lo; i < hi; i++ {
		for j := hi - 1; j > i; j-- {
			// 和前一个元素比较
			if data.Less(j, j-1) {
				// 向前冒一位
				data.Swap(j, j-1)
			}
		}
	}
}
