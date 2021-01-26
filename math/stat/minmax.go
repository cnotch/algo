// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package stat

// MinMax 获取最小、最大值的索引位置
func MinMax(n int, less func(i, j int) bool) (min, max int) {
	// 0 长度
	if n == 0 {
		return -1, -1
	}

	min = 0
	max = 0
	if n == 1 {
		return
	}

	// 假设 len 是奇数
	i := 1
	if n%2 == 0 { // 如果 len 是偶数
		i = 2
		if less(0, 1) {
			min = 0
			max = 1
		} else {
			min = 1
			max = 0
		}
	}

	// 步长为 2 循环
	for ; i < n; i += 2 {
		if less(i, i+1) {
			if less(i, min) {
				min = i
			}
			if less(max, i+1) {
				max = i + 1
			}
		} else {
			if less(i+1, min) {
				min = i + 1
			}
			if less(max, i) {
				max = i
			}
		}
	}
	return
}

// Min 获取最小值的索引位置
func Min(n int, less func(i, j int) bool) (min int) {
	// 0 长度
	if n == 0 {
		return -1
	}

	min = 0
	if n == 1 {
		return
	}

	for i := 1; i < n; i++ {
		if less(i, min) {
			min = i
		}
	}
	return
}

// Max 获取最大值的索引位置
func Max(n int, greater func(i, j int) bool) (max int) {
	// 0 长度
	if n == 0 {
		return -1
	}

	max = 0
	if n == 1 {
		return
	}

	for i := 1; i < n; i++ {
		if greater(i, max) {
			max = i
		}
	}
	return
}
