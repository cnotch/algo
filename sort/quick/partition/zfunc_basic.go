// Code generated from basic.go using genzfunc.go; DO NOT EDIT.

// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package partition

import "math/rand"

// Auto-generated variant of basic.go:basic
func basicZfunc(data LessSwap, lo, hi int) (midlo int, midhi int) {
	povit := hi - 1
	midlo = lo - 1
	for i := lo; i < povit; i++ {
		if !data.Less(povit, i) {
			midlo++
			data.Swap(midlo, i)
		}
	}
	midlo++
	data.Swap(midlo, povit)
	return midlo, midlo + 1
}

// Auto-generated variant of basic.go:randomized
func randomizedZfunc(data LessSwap, lo, hi int) (midlo int, midhi int) {
	i := rand.Intn(hi-lo) + lo
	data.Swap(i, hi-1)
	return basicZfunc(data, lo, hi)
}
