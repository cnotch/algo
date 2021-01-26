// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go run ../../genzfunc.go -i threeway.go -o zfunc_threeway.go

package partition

// 来自 go 标准库，它是一种非严格的三向切分实现。
func threeWay(data Interface, lo, hi int) (midlo, midhi int) {
	// 选择合适的中值
	m := int(uint(lo+hi) >> 1)
	if hi-lo > 40 {
		// Tukey's ``Ninther,'' median of three medians of three.
		s := (hi - lo) / 8
		medianOfThree(data, lo, lo+s, lo+2*s)
		medianOfThree(data, m, m-s, m+s)
		medianOfThree(data, hi-1, hi-1-s, hi-1-2*s)
	}
	medianOfThree(data, lo, m, hi-1)
	// Now: data[m] <= data[lo] <= data[hi-1]

	pivot := lo
	a, c := lo+1, hi-1

	for ; a < c && data.Less(a, pivot); a++ {
	}
	// Now:
	//  data[lo] == pivot
	//  data[lo < i < a] < pivot
	//  data[a <= i < c] unexamined
	//  data[hi-1] >= pivot

	b := a
	for {
		for ; b < c && !data.Less(pivot, b); b++ { // data[b] <= pivot
		}
		for ; b < c && data.Less(pivot, c-1); c-- { // data[c-1] > pivot
		}
		if b >= c {
			break
		}
		// data[b] > pivot; data[c-1] <= pivot
		data.Swap(b, c-1)
		b++
		c--
	}
	// Now:
	//  data[lo] == pivot
	//  data[lo < i < a] < pivot
	//	data[a <= i < b] <= pivot
	//	data[b <= i < c] 0 elem( b==c ), position for == pivot
	//	data[c <= i < hi-1] > pivot
	//	data[hi-1] >= pivot

	// protect 用来指示拥有多个重复 pivot
	// 它是个估计值，以下两个 if 语句块是否被执行不会影响排序的正确性
	// 严格的三向切分，应该直接执行最后一个 if 语句块中的代码

	// 假设前面执行了 Tukey 中值算法：
	//  pivot 是9个数的中值，那么9个数中至少有3个 >=pivot
	//  如果 >pivot 的数少于3个，那么表示其他的数 =pivot
	protect := hi-c < 5

	// 如果 >pivot 的数超过5个并且 <=pivot 的元素超过总数 3/4，
	// 此时存在多个 pivot 可能性较大，执行多点测试进一步验证
	if !protect && hi-c < (hi-lo)/4 {
		// Lets test some points for equality to pivot
		dups := 0
		if !data.Less(pivot, hi-1) { // data[hi-1] = pivot
			data.Swap(c, hi-1)
			c++
			dups++
		}
		if !data.Less(b-1, pivot) { // data[b-1] = pivot
			b--
			dups++
		}
		// m-lo = (hi-lo)/2 > 6
		// b-lo > (hi-lo)*3/4-1 > 8
		// ==> m < b ==> data[m] <= pivot
		if !data.Less(m, pivot) { // data[m] = pivot
			data.Swap(m, b-1)
			b--
			dups++
		}
		// if at least 2 points are equal to pivot, assume skewed distribution
		protect = dups > 1
	}

	// 估计 data[a <= i < b] 中有多个 pivot 值
	if protect {
		// Protect against a lot of duplicates
		// Add invariant:
		//	data[a <= i < b] unexamined
		//	data[b <= i < c] = pivot
		for {
			for ; a < b && !data.Less(b-1, pivot); b-- { // data[b] == pivot
			}
			for ; a < b && data.Less(a, pivot); a++ { // data[a] < pivot
			}
			if a >= b {
				break
			}
			// data[a] == pivot; data[b-1] < pivot
			data.Swap(a, b-1)
			a++
			b--
		}
	}
	// Now:
	//  data[lo] == pivot
	//  data[lo < i < b] < pivot
	//	data[b <= i < c] == pivot
	//	data[c <= i < hi-1] > pivot
	//	data[hi-1] >= pivot

	// Swap pivot into middle
	data.Swap(pivot, b-1)
	return b - 1, c
}

// 三向切分严格版实现。
func strictThreeWay(data Interface, lo, hi int) (midlo, midhi int) {
	// 选择合适的中值
	m := int(uint(lo+hi) >> 1)
	if hi-lo > 40 {
		// Tukey's ``Ninther,'' median of three medians of three.
		s := (hi - lo) / 8
		medianOfThree(data, lo, lo+s, lo+2*s)
		medianOfThree(data, m, m-s, m+s)
		medianOfThree(data, hi-1, hi-1-s, hi-1-2*s)
	}
	medianOfThree(data, lo, m, hi-1)
	// Now: data[m] <= data[lo] <= data[hi-1]

	pivot := lo
	a, c := lo+1, hi-1

	for ; a < c && data.Less(a, pivot); a++ {
	}
	// Now:
	//  data[lo] == pivot
	//  data[lo < i < a] < pivot
	//  data[a <= i < c] unexamined
	//  data[hi-1] >= pivot

	b := a
	for {
		for ; b < c && !data.Less(pivot, b); b++ { // data[b] <= pivot
		}
		for ; b < c && data.Less(pivot, c-1); c-- { // data[c-1] > pivot
		}
		if b >= c {
			break
		}
		// data[b] > pivot; data[c-1] <= pivot
		data.Swap(b, c-1)
		b++
		c--
	}
	// Now:
	//  data[lo] == pivot
	//  data[lo < i < a] < pivot
	//	data[a <= i < b] <= pivot
	//	data[b <= i < c] 0 elem( b==c ), position for == pivot
	//	data[c <= i < hi-1] > pivot
	//	data[hi-1] >= pivot

	// Add invariant:
	//	data[a <= i < b] unexamined
	//	data[b <= i < c] = pivot
	for {
		for ; a < b && !data.Less(b-1, pivot); b-- { // data[b] == pivot
		}
		for ; a < b && data.Less(a, pivot); a++ { // data[a] < pivot
		}
		if a >= b {
			break
		}
		// data[a] == pivot; data[b-1] < pivot
		data.Swap(a, b-1)
		a++
		b--
	}

	// Now:
	//  data[lo] == pivot
	//  data[lo < i < b] < pivot
	//	data[b <= i < c] == pivot
	//	data[c <= i < hi-1] > pivot
	//	data[hi-1] >= pivot

	if !data.Less(pivot, hi-1) { // data[hi-1] = pivot
		data.Swap(c, hi-1)
		c++
	}

	// Swap pivot into middle
	data.Swap(pivot, b-1)
	return b - 1, c
}

// medianOfThree moves the median of the three values data[m0], data[m1], data[m2] into data[m1].
func medianOfThree(data Interface, m1, m0, m2 int) {
	// sort 3 elements
	if data.Less(m1, m0) {
		data.Swap(m1, m0)
	}
	// data[m0] <= data[m1]
	if data.Less(m2, m1) {
		data.Swap(m2, m1)
		// data[m0] <= data[m2] && data[m1] < data[m2]
		if data.Less(m1, m0) {
			data.Swap(m1, m0)
		}
	}
	// now data[m0] <= data[m1] <= data[m2]
}
