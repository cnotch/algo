// Code generated from sort.go using genzfunc.go; DO NOT EDIT.

// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package heap

// Auto-generated variant of sort.go:heapSort
func heapSortZfunc(data LessSwap, a, b int) {
	first := a
	lo := 0
	hi := b - a
	for i := (hi - 1) / 2; i >= 0; i-- {
		maxHeapifyZfunc(data, i, hi, first)
	}
	for i := hi - 1; i >= 0; i-- {
		data.Swap(first, first+i)
		maxHeapifyZfunc(data, lo, i, first)
	}
}

// Auto-generated variant of sort.go:maxHeapify
func maxHeapifyZfunc(data LessSwap, lo, hi, first int) {
	root := lo
	child := 2*root + 1
	for child < hi {
		if child+1 < hi && data.Less(first+child, first+child+1) {
			child++
		}
		if data.Less(first+child, first+root) {
			return
		}
		data.Swap(first+root, first+child)
		root = child
		child = 2*root + 1
	}
}
