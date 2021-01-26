// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//go:generate go run ../genzfunc.go

package quick

import (
	"sort"

	pf "github.com/cnotch/algo/sort/quick/partition"
)

// Interface 标准库 sort.Interface 别名
type Interface = sort.Interface

// LessSwap partition.LessSwap 别名
type LessSwap = pf.LessSwap

var (
	partition      = pf.Get(pf.ThreeWay)
	partitionZfunc = pf.GetZfunc(pf.ThreeWay)
)

// Sort 对 data 执行快速排序
func Sort(data Interface) {
	quickSort(data, 0, data.Len())
}

// 单侧递归实现：用循环减少递归
func quickSort(data Interface, lo, hi int) {
	for hi-lo > 2 {
		midlo, midhi := partition(data, lo, hi)
		// 选择元素少的一侧执行递归，多的一侧用循环代替递归;
		// 最大限度减少递归深度，最大递归深度lg(hi-lo)
		if midlo-lo < hi-midhi {
			quickSort(data, lo, midlo)
			lo = midhi // i.e., quickSort(data, midhi, hi)
		} else {
			quickSort(data, midhi, hi)
			hi = midlo // i.e., quickSort(data, lo, midlo)
		}
	}

	if hi-lo == 2 {
		if data.Less(hi-1, lo) {
			data.Swap(lo, hi-1)
		}
	}
}
