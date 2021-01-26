// Code generated from threeway.go using genzfunc.go; DO NOT EDIT.

// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package partition

// Auto-generated variant of threeway.go:threeWay
func threeWayZfunc(data LessSwap, lo, hi int) (midlo, midhi int) {
	m := int(uint(lo+hi) >> 1)
	if hi-lo > 40 {
		s := (hi - lo) / 8
		medianOfThreeZfunc(data, lo, lo+s, lo+2*s)
		medianOfThreeZfunc(data, m, m-s, m+s)
		medianOfThreeZfunc(data, hi-1, hi-1-s, hi-1-2*s)
	}
	medianOfThreeZfunc(data, lo, m, hi-1)
	pivot := lo
	a, c := lo+1, hi-1
	for ; a < c && data.Less(a, pivot); a++ {
	}
	b := a
	for {
		for ; b < c && !data.Less(pivot, b); b++ {
		}
		for ; b < c && data.Less(pivot, c-1); c-- {
		}
		if b >= c {
			break
		}
		data.Swap(b, c-1)
		b++
		c--
	}
	protect := hi-c < 5
	if !protect && hi-c < (hi-lo)/4 {
		dups := 0
		if !data.Less(pivot, hi-1) {
			data.Swap(c, hi-1)
			c++
			dups++
		}
		if !data.Less(b-1, pivot) {
			b--
			dups++
		}
		if !data.Less(m, pivot) {
			data.Swap(m, b-1)
			b--
			dups++
		}
		protect = dups > 1
	}
	if protect {
		for {
			for ; a < b && !data.Less(b-1, pivot); b-- {
			}
			for ; a < b && data.Less(a, pivot); a++ {
			}
			if a >= b {
				break
			}
			data.Swap(a, b-1)
			a++
			b--
		}
	}
	data.Swap(pivot, b-1)
	return b - 1, c
}

// Auto-generated variant of threeway.go:strictThreeWay
func strictThreeWayZfunc(data LessSwap, lo, hi int) (midlo, midhi int) {
	m := int(uint(lo+hi) >> 1)
	if hi-lo > 40 {
		s := (hi - lo) / 8
		medianOfThreeZfunc(data, lo, lo+s, lo+2*s)
		medianOfThreeZfunc(data, m, m-s, m+s)
		medianOfThreeZfunc(data, hi-1, hi-1-s, hi-1-2*s)
	}
	medianOfThreeZfunc(data, lo, m, hi-1)
	pivot := lo
	a, c := lo+1, hi-1
	for ; a < c && data.Less(a, pivot); a++ {
	}
	b := a
	for {
		for ; b < c && !data.Less(pivot, b); b++ {
		}
		for ; b < c && data.Less(pivot, c-1); c-- {
		}
		if b >= c {
			break
		}
		data.Swap(b, c-1)
		b++
		c--
	}
	for {
		for ; a < b && !data.Less(b-1, pivot); b-- {
		}
		for ; a < b && data.Less(a, pivot); a++ {
		}
		if a >= b {
			break
		}
		data.Swap(a, b-1)
		a++
		b--
	}
	if !data.Less(pivot, hi-1) {
		data.Swap(c, hi-1)
		c++
	}
	data.Swap(pivot, b-1)
	return b - 1, c
}

// Auto-generated variant of threeway.go:medianOfThree
func medianOfThreeZfunc(data LessSwap, m1, m0, m2 int) {
	if data.Less(m1, m0) {
		data.Swap(m1, m0)
	}
	if data.Less(m2, m1) {
		data.Swap(m2, m1)
		if data.Less(m1, m0) {
			data.Swap(m1, m0)
		}
	}
}
