// Code generated from sort.go using genzfunc.go; DO NOT EDIT.

// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package merge

// Auto-generated variant of sort.go:mergeSort
func mergeSortZfunc(data LessSwap, lo, hi int) {
	if hi-lo <= 1 {
		return
	}
	mid := int(uint(lo+hi) >> 1)
	mergeSortZfunc(data, lo, mid)
	mergeSortZfunc(data, mid, hi)
	symMergeZfunc(data, lo, mid, hi)
}

// Auto-generated variant of sort.go:symMerge
func symMergeZfunc(data LessSwap, a, m, b int) {
	if m-a == 1 {
		i := m
		j := b
		for i < j {
			h := int(uint(i+j) >> 1)
			if data.Less(h, a) {
				i = h + 1
			} else {
				j = h
			}
		}
		for k := a; k < i-1; k++ {
			data.Swap(k, k+1)
		}
		return
	}
	if b-m == 1 {
		i := a
		j := m
		for i < j {
			h := int(uint(i+j) >> 1)
			if !data.Less(m, h) {
				i = h + 1
			} else {
				j = h
			}
		}
		for k := m; k > i; k-- {
			data.Swap(k, k-1)
		}
		return
	}
	mid := int(uint(a+b) >> 1)
	n := mid + m
	var start, r int
	if m > mid {
		start = n - b
		r = mid
	} else {
		start = a
		r = m
	}
	p := n - 1
	for start < r {
		c := int(uint(start+r) >> 1)
		if !data.Less(p-c, c) {
			start = c + 1
		} else {
			r = c
		}
	}
	end := n - start
	if start < m && m < end {
		rotateZfunc(data, start, m, end)
	}
	if a < start && start < mid {
		symMergeZfunc(data, a, start, mid)
	}
	if mid < end && end < b {
		symMergeZfunc(data, mid, end, b)
	}
}

// Auto-generated variant of sort.go:rotate
func rotateZfunc(data LessSwap, a, m, b int) {
	i := m - a
	j := b - m
	for i != j {
		if i > j {
			swapRangeZfunc(data, m-i, m, j)
			i -= j
		} else {
			swapRangeZfunc(data, m-i, m+j-i, i)
			j -= i
		}
	}
	swapRangeZfunc(data, m-i, m, i)
}

// Auto-generated variant of sort.go:swapRange
func swapRangeZfunc(data LessSwap, a, b, n int) {
	for i := 0; i < n; i++ {
		data.Swap(a+i, b+i)
	}
}
