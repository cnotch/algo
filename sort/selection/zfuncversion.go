// Code generated from sort.go using genzfunc.go; DO NOT EDIT.

// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package selection

// Auto-generated variant of sort.go:selectionSort
func selectionSortZfunc(data LessSwap, lo, hi int) {
	for i := lo; i < hi-1; i++ {
		min := i
		for j := i + 1; j < hi; j++ {
			if data.Less(j, min) {
				min = j
			}
		}
		data.Swap(i, min)
	}
}
