// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//go:generate go run ../genzfunc.go

package insertion

import (
	"sort"

	"github.com/cnotch/algo/sort/quick/partition"
)

// Interface 标准库 sort.Interface 别名
type Interface = sort.Interface

// LessSwap partition.LessSwap 别名
type LessSwap = partition.LessSwap

// Sort 对 data 执行插入排序
func Sort(data Interface) {
	insertionSort(data, 0, data.Len())
}

// Get .
func Get() func(data Interface, lo, hi int) {
	return insertionSort
}

func insertionSort(data Interface, lo, hi int) {
	for i := lo + 1; i < hi; i++ {
		// [lo:i)已经是有序，将下一个元素插入到有序集合的合适位置
		for j := i; j > lo && data.Less(j, j-1); j-- {
			data.Swap(j, j-1)
		}
	}
}
