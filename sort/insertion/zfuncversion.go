// Code generated from sort.go using genzfunc.go; DO NOT EDIT.

// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package insertion

// Auto-generated variant of sort.go:insertionSort
func insertionSortZfunc(data LessSwap, lo, hi int) {
	for i := lo + 1; i < hi; i++ {
		for j := i; j > lo && data.Less(j, j-1); j-- {
			data.Swap(j, j-1)
		}
	}
}
