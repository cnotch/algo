// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//go:generate go run ../genzfunc.go

package intro

import (
	"sort"

	"github.com/cnotch/algo/sort/heap"
	pf "github.com/cnotch/algo/sort/quick/partition"
	"github.com/cnotch/algo/sort/shell"
)

// Interface 标准库 sort.Interface 别名
type Interface = sort.Interface

// LessSwap partition.LessSwap 别名
type LessSwap = pf.LessSwap

var (
	partition      = pf.Get(pf.ThreeWay)
	partitionZfunc = pf.GetZfunc(pf.ThreeWay)
	heapSort       = heap.Get()
	heapSortZfunc  = heap.GetZfunc()
	shellSort      = shell.Get()
	shellSortZfunc = shell.GetZfunc()
)

// Sort 对 data 执行快速排序
func Sort(data Interface) {
	introSort(data, 0, data.Len(), maxDepth(data.Len()))
}

// 内省排序实现
func introSort(data Interface, lo, hi, maxDepth int) {
	for hi-lo > 16 {
		if maxDepth == 0 { // 超过深度使用堆排序
			heapSort(data, lo, hi)
			return
		}

		maxDepth--
		midlo, midhi := partition(data, lo, hi)
		// 选择元素少的一侧执行递归，多的一侧用循环代替递归;
		// 最大限度减少递归深度，最大递归深度lg(hi-lo)
		if midlo-lo < hi-midhi {
			introSort(data, lo, midlo, maxDepth)
			lo = midhi // i.e., introSort(data, midhi, hi, maxDepth)
		} else {
			introSort(data, midhi, hi, maxDepth)
			hi = midlo // i.e., introSort(data, lo, midlo, maxDepth)
		}
	}

	if hi-lo > 1 { // 小于16个元素直接使用希尔排序
		shellSort(data, lo, hi)
	}
}

// 快速排序的最大深度：lg(n+)+1
func maxDepth(n int) int {
	var depth int
	for i := n; i > 0; i >>= 1 {
		depth++
	}
	return depth + 1
}
