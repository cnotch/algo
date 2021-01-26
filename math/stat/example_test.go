// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package stat_test

import (
	"fmt"
	"sort"

	"github.com/cnotch/algo/math/stat"
)

func ExampleSelect() {
	var ints = []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

	for k := 0; k < len(ints); k++ {
		list := append([]int(nil), ints...)
		stat.Select(sort.IntSlice(list), k)
		sorted := append([]int(nil), list...)
		sort.Ints(sorted)
		want := sorted[k]
		fmt.Println(want)
	}
	// Output:
	// -5467984
	// -784
	// 0
	// 0
	// 42
	// 59
	// 74
	// 238
	// 905
	// 959
	// 7586
	// 7586
	// 9845
}

func ExampleMinMax() {
	var a = []int{95, 45, 15, 78, 84, 51, 24, 12}
	min, max := stat.MinMax(len(a), func(i, j int) bool { return a[i] < a[j] })
	fmt.Println(a[min])
	fmt.Println(a[max])

	// Output:
	// 12
	// 95
}
