// Code generated from sort.go using genzfunc.go; DO NOT EDIT.

// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package bubble

// Auto-generated variant of sort.go:bubbleSort
func bubbleSortZfunc(data LessSwap, lo, hi int) {
	for i := lo; i < hi; i++ {
		for j := hi - 1; j > i; j-- {
			if data.Less(j, j-1) {
				data.Swap(j, j-1)
			}
		}
	}
}
