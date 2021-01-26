// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package bucket

// Ints 对 a 进行桶排序
// 	a 待排序的集合；
//  k 桶索引值最大值，它表示桶索引值的范围为 [0,k]；
//	key 计算集合中元素桶索引值函数，返回值必须在 [0,k]；
//	stable 对指定桶执行排序操作的函数
func Ints(a []int, k int, key func(i int) int, stable func(bucket []int)) {
	// 初始化
	buckets := make([][]int, k+1)

	for i := 0; i < len(a); i++ {
		bi := key(i)
		buckets[bi] = append(buckets[bi], a[i])
	}

	offset := 0
	for i := 0; i <= k; i++ {
		bucket := buckets[i]
		if len(bucket) == 0 {
			continue
		}

		// 对单个桶排序
		stable(bucket)
		// 复制到初始切片,并步进下次复制的位置
		offset += copy(a[offset:], bucket)
	}
}
