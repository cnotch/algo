// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package counting

// Sort 计数排序。
// 	n 集合包含元素的个数；
//  k 计数参考值范围 [0,k]；
//	key 计算集合中元素计数参照值函数，返回值必须在 [0,k]；
//	assign 排序后元素赋值函数。
func Sort(n, k int, key func(r int) int, assign func(l, r int)) {
	countingSort(n, k, key, assign)
}

// 计数排序标准的算法逻辑
func countingSort(n, k int, key func(r int) int, assign func(l, r int)) {
	c := make([]int, k+1)

	// 初始化计数
	for i := range c {
		c[i] = 0
	}

	// [0,k] 中每个 key 出现的次数
	for i := 0; i < n; i++ {
		ek := key(i)
		c[ek]++
	}

	//[0,k] 中每个 key 累计元素个数
	for i := 1; i < len(c); i++ {
		c[i] = c[i] + c[i-1]
	}

	// 从高到低索引顺序拷贝元素
	for i := n - 1; i >= 0; i-- {
		ek := key(i)
		c[ek]--
		assign(c[ek], i)
	}
}
