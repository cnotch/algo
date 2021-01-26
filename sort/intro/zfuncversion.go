// Code generated from sort.go using genzfunc.go; DO NOT EDIT.

// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package intro

// Auto-generated variant of sort.go:introSort
func introSortZfunc(data LessSwap, lo, hi, maxDepth int) {
	for hi-lo > 16 {
		if maxDepth == 0 {
			heapSortZfunc(data, lo, hi)
			return
		}
		maxDepth--
		midlo, midhi := partitionZfunc(data, lo, hi)
		if midlo-lo < hi-midhi {
			introSortZfunc(data, lo, midlo, maxDepth)
			lo = midhi
		} else {
			introSortZfunc(data, midhi, hi, maxDepth)
			hi = midlo
		}
	}
	if hi-lo > 1 {
		shellSortZfunc(data, lo, hi)
	}
}
