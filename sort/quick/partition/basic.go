// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//go:generate go run ../../genzfunc.go -i basic.go -o zfunc_basic.go

package partition

import "math/rand"

// 标准的快速排序分区
func basic(data Interface, lo, hi int) (midlo int, midhi int) {
	povit := hi - 1 // 最后元素作为基准值
	midlo = lo - 1
	for i := lo; i < povit; i++ {
		if !data.Less(povit, i) { // 和基准值比较 [povit]>=[i]
			midlo++
			data.Swap(midlo, i) // 保持小于基准值的元素索引<=midlo
		}
	}
	midlo++
	data.Swap(midlo, povit) // 基准值元素交互到最终的位置
	return midlo, midlo + 1
}

// 随机选择基准值的快速排序分区
func randomized(data Interface, lo, hi int) (midlo int, midhi int) {
	i := rand.Intn(hi-lo) + lo
	data.Swap(i, hi-1)
	return basic(data, lo, hi)
}
