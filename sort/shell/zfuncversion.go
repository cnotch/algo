// Code generated from sort.go using genzfunc.go; DO NOT EDIT.

// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package shell

// Auto-generated variant of sort.go:shellSort
func shellSortZfunc(data LessSwap, lo, hi int) {
	gap := int(uint(hi-lo) >> 1)
	for ; gap > 0; gap /= 2 {
		for i := lo + gap; i < hi; i++ {
			for j := i - gap; j >= lo && data.Less(j+gap, j); j -= gap {
				data.Swap(j, j+gap)
			}
		}
	}
}
