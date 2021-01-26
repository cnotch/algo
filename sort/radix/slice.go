// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package radix

import (
	"reflect"

	"github.com/cnotch/algo/reflect/iface"
	"github.com/cnotch/algo/sort/counting"
)

// Slice 对 slice 执行升序排序，以计数排序为稳定排序算法
// 	d	位数
//	keyMaker 用于创建计数排序所需的 k 值和 key 函数；slice 当前被计数排序的切片，p 第几位
func Slice(slice interface{}, d int, keyMaker func(slice interface{}, p int) (k int, key func(i int) int)) {
	Sort(d, stable4Slice(slice, d, keyMaker))
}

// Ints 对 []int 执行升序排序，以计数排序为稳定排序算法
func Ints(ints []int, d int) {
	Sort(d, stable4Ints(ints, d))
}

// Strings 对 []Strings 执行升序排序，以计数排序为稳定排序算法
func Strings(strings []string, d int) {
	Sort(d, stable4Strings(strings, d))
}

// stable4Ints 创建整数 Slice 的计数排序函数
func stable4Ints(ints []int, d int) func(p int) {
	n := len(ints)
	// 计数排序所需的额外空间
	temp := make([]int, n)
	inslice, outslice := temp, ints

	return func(p int) {
		// 循环使用临时空间 Slice，
		// 保证 inslice 是待排序 Slice，outslice 是排序结果
		inslice, outslice = outslice, inslice

		counting.Sort(n, 9,
			func(i int) int {
				return DecDigit(inslice[i], p)
			},
			func(l, r int) {
				outslice[l] = inslice[r]
			})

		// 如果最终排序结果不是最初 Slice
		if p == d-1 && p%2 == 0 {
			copy(ints, outslice)
			temp, inslice, outslice = nil, nil, nil
		}
	}
}

// stable4Strings 创建字串 Slice 的计数排序函数
func stable4Strings(strings []string, d int) func(p int) {
	n := len(strings)
	// 计数排序所需的额外空间
	temp := make([]string, n)
	inslice, outslice := temp, strings

	return func(p int) {
		// 循环使用临时空间 Slice，
		// 保证 inslice 是待排序 Slice，outslice 是排序结果
		inslice, outslice = outslice, inslice

		counting.Sort(n, 255,
			func(i int) int {
				return CharDigit(inslice[i], d, p)
			},
			func(l, r int) {
				outslice[l] = inslice[r]
			})

		// 如果最终排序结果不是最初 Slice
		if p == d-1 && p%2 == 0 {
			copy(strings, outslice)
			temp, inslice, outslice = nil, nil, nil
		}
	}
}

// stable4Slice 创建 Slice 的计数排序函数
// 	d	位数
//	keyMaker 用于创建计数排序所需的 k 值和 key 函数；slice 当前被计数排序的切片，p 第几位
func stable4Slice(slice interface{}, d int, keyMaker func(slice interface{}, p int) (k int, key func(i int) int)) func(p int) {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		panic("Isn't slice type")
	}

	n := rv.Len()
	// 计数排序所需的额外空间
	temp := reflect.MakeSlice(rv.Type(), n, n)

	inslice, outslice := temp.Interface(), slice
	assigns := [2]func(l, r int){
		iface.Assigner(inslice, outslice),
		iface.Assigner(outslice, inslice)}

	return func(p int) {
		// 循环使用临时空间 Slice，
		// 保证 inslice 是待排序 Slice，outslice 是排序结果
		inslice, outslice = outslice, inslice
		k, key := keyMaker(inslice, p)

		counting.Sort(n,
			k, key,
			assigns[p%2])

		// 如果最终排序结果不是最初 Slice
		if p == d-1 && p%2 == 0 {
			reflect.Copy(rv, temp)
			inslice, outslice = nil, nil
		}
	}
}
