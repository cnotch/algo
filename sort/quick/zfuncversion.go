// Code generated from sort.go using genzfunc.go; DO NOT EDIT.

// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package quick

// Auto-generated variant of sort.go:quickSort
func quickSortZfunc(data LessSwap, lo, hi int) {
	for hi-lo > 2 {
		midlo, midhi := partitionZfunc(data, lo, hi)
		if midlo-lo < hi-midhi {
			quickSortZfunc(data, lo, midlo)
			lo = midhi
		} else {
			quickSortZfunc(data, midhi, hi)
			hi = midlo
		}
	}
	if hi-lo == 2 {
		if data.Less(hi-1, lo) {
			data.Swap(lo, hi-1)
		}
	}
}
