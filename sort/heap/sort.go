// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//go:generate go run ../genzfunc.go

package heap

import (
	"sort"

	"github.com/cnotch/algo/sort/quick/partition"
)

// Interface 标准库 sort.Interface 别名
type Interface = sort.Interface

// LessSwap partition.LessSwap 别名
type LessSwap = partition.LessSwap

// Sort 对 data 执行堆排序
func Sort(data Interface) {
	heapSort(data, 0, data.Len())
}

// Get .
func Get() func(data Interface, lo, hi int) {
	return heapSort
}

// heapSort 堆排序算法
func heapSort(data Interface, a, b int) {
	first := a
	lo := 0
	hi := b - a

	// 构建初始大顶堆
	for i := (hi - 1) / 2; i >= 0; i-- {
		maxHeapify(data, i, hi, first)
	}

	// 1. 将顶部元素置换到最后一个（当前范围最大的）
	// 2. 重新为第一个元素在大顶堆中找到合适的位置
	// 3. 重复1和2，直达第一个元素是最小的元素，此时i=0
	for i := hi - 1; i >= 0; i-- {
		data.Swap(first, first+i)
		maxHeapify(data, lo, i, first)
	}
}

// 指定范围构建大顶堆（二叉树），非递归版本
func maxHeapify(data Interface, lo, hi, first int) {
	root := lo
	child := 2*root + 1 // 左子

	for child < hi {
		// 确定子元素中较大的一个
		if child+1 < hi && data.Less(first+child, first+child+1) {
			child++
		}
		// 如果都小于父元素
		if data.Less(first+child, first+root) {
			return
		}
		// 交换父子元素
		data.Swap(first+root, first+child)
		// 继续子树构建
		root = child
		child = 2*root + 1
	}
}

// maxHeapify 递归版本
// func maxHeapify(data Interface, lo, hi, first int) {
// 	root := lo
// 	child := 2*root + 1 // 左子
//
// 	if child < hi {
// 		// 确定子元素中较大的一个
// 		if child+1 < hi && data.Less(first+child, first+child+1) {
// 			child++
// 		}
// 		// 如果都小于父元素
// 		if data.Less(first+child, first+root) {
// 			return
// 		}
// 		// 交换父子元素
// 		data.Swap(first+root, first+child)
// 		// 继续子树构建
// 		maxHeapify(data, child, hi, first)
// 	}
// }
