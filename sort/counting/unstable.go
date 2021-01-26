// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package counting

import (
	"reflect"
)

// Unstable 原地交换计数排序
func Unstable(n, k int, key func(i int) int, swap func(i, j int)) {
	unstable(n, k, key, swap)
}

// SliceUnstable 对 Slice 原地交换计数排序
func SliceUnstable(slice interface{}, k int, key func(i int) int) {
	rv := reflect.ValueOf(slice)
	swap := reflect.Swapper(slice)
	n := rv.Len()
	unstable(n, k, key, swap)
}

// 计数排序原地交换一种实现，性能损失不大，但不属于稳定的排序
// 空间 O(2k)，
func unstable(n, k int, key func(i int) int, swap func(i, j int)) {
	c := make([]int, k+1)

	// 初始化计数起点
	for i := 0; i < k+1; i++ {
		c[i] = 0
	}

	// [0,k] 中每个 key 出现的次数
	for i := 0; i < n; i++ {
		ck := key(i)
		c[ck]++
	}

	//[0,k] 中每个 key 累计元素个数
	for i := 1; i < len(c); i++ {
		c[i] = c[i] + c[i-1]
	}

	// 最后交换位置
	lastp := make([]int, k+1)
	copy(lastp, c)

	// 将元素交换到合适的位置
	for i := k; i > 0; i-- {
		for j := lastp[i] - 1; j >= c[i-1]; {
			ek := key(j)
			lastp[ek]--
			if ek == i { // 已经被排序
				j--
				continue
			}

			swap(lastp[ek], j)
		}
	}
}
