// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//go:generate go run ../genzfunc.go

package shell

import (
	"sort"

	"github.com/cnotch/algo/sort/quick/partition"
)

// Interface 标准库 sort.Interface 别名
type Interface = sort.Interface

// LessSwap partition.LessSwap 别名
type LessSwap = partition.LessSwap

// Sort 对 data 执行希尔排序
func Sort(data Interface) {
	shellSort(data, 0, data.Len())
}

// Get .
func Get() func(data Interface, lo, hi int) {
	return shellSort
}

func shellSort(data Interface, lo, hi int) {
	// 初始间隔
	gap := int(uint(hi-lo) >> 1)
	for ; gap > 0; gap /= 2 {
		// 开始间隔为 gap 插入的排序；标准的插入排序间隔为 1
		for i := lo + gap; i < hi; i++ {
			for j := i - gap; j >= lo && data.Less(j+gap, j); j -= gap {
				data.Swap(j, j+gap)
			}
		}
	}
}
