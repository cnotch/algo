// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package stat

import (
	"reflect"
	"sort"

	"github.com/cnotch/algo/sort/quick/partition"
)

// Interface sort.Interface 别名
type Interface = sort.Interface

// Select 选择 data 中第 k 小的元素，返回该元素的索引号。
func Select(data Interface, k int) int {
	lo, hi := 0, data.Len()
	if k >= hi || k < 0 {
		panic("Select: k index out of range")
	}

	pf := partition.Get(partition.ThreeWay)
	for hi-lo > 2 {
		midlo, midhi := pf(data, lo, hi)
		if k >= midlo && k < midhi {
			break
		}

		if k < midlo {
			hi = midlo
		} else {
			lo = midhi
		}
	}

	if hi-lo == 2 {
		if data.Less(hi-1, lo) {
			data.Swap(lo, hi-1)
		}
	}

	return k
}

// SliceSelect 选择 data 中第 k 小的元素，返回该元素的索引号。
func SliceSelect(slice interface{}, k int, less func(i, j int) bool) int {
	rv := reflect.ValueOf(slice)
	data := partition.LessSwap{Less: less, Swap: reflect.Swapper(slice)}
	lo, hi := 0, rv.Len()

	if k >= hi || k < 0 {
		panic("SliceSelect: k index out of range")
	}

	pf := partition.GetZfunc(partition.ThreeWay)
	for hi-lo > 2 {
		midlo, midhi := pf(data, lo, hi)
		if k >= midlo && k < midhi {
			return k
		}

		if k < midlo {
			hi = midlo
		} else {
			lo = midhi
		}
	}

	if hi-lo == 2 {
		if data.Less(hi-1, lo) {
			data.Swap(lo, hi-1)
		}
	}

	return k
}
