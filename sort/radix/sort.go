// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package radix

// Sort 基数排序的模板方法。
// 	d	待排序集合的最大位数
//	stable 稳定排序函数，用于第 p 位执行排序
func Sort(d int, stable func(p int)) {
	for p := 0; p < d; p++ {
		stable(p)
	}
}
