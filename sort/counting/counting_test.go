// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package counting_test

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/cnotch/algo/sort/counting"
)

const k4test = 11

var testSlice []int

func init() {
	const count = 64
	data := make([]int, count)
	for i := 0; i < count; i++ {
		data[i] = rand.Intn(k4test + 1)
	}
	testSlice = data
}

func TestCountingSort(t *testing.T) {
	t.Run("counting.Sort", func(t *testing.T) {
		ints := make([]int, len(testSlice))
		counting.Sort(len(testSlice), k4test,
			func(r int) int { return testSlice[r] },
			func(l, r int) { ints[l] = testSlice[r] })
		if !sort.IntsAreSorted(ints) {
			t.Errorf("sorted %v", testSlice)
			t.Errorf("   got %v", ints)
		}
	})

}

func TestCountingSlice(t *testing.T) {
	t.Run("counting.Slice", func(t *testing.T) {
		ints := make([]int, len(testSlice))
		copy(ints, testSlice)
		counting.Slice(ints, k4test,
			func(i int) int { return ints[i] })
		if !sort.IntsAreSorted(ints) {
			t.Errorf("sorted %v", testSlice)
			t.Errorf("   got %v", ints)
		}
	})
}

func TestCountingUnstable(t *testing.T) {
	t.Run("counting.Unstable", func(t *testing.T) {
		ints := append([]int(nil), testSlice...)
		counting.Unstable(len(ints), k4test,
			func(i int) int { return ints[i] },
			func(i, j int) { ints[i], ints[j] = ints[j], ints[i] })
		if !sort.IntsAreSorted(ints) {
			t.Errorf("sorted %v", testSlice)
			t.Errorf("   got %v", ints)
		}
	})
}

func TestCountingSliceUnstable(t *testing.T) {
	t.Run("counting.SliceUnstable", func(t *testing.T) {
		ints := append([]int(nil), testSlice...)
		counting.SliceUnstable(ints, k4test, func(i int) int { return ints[i] })
		if !sort.IntsAreSorted(ints) {
			t.Errorf("sorted %v", testSlice)
			t.Errorf("   got %v", ints)
		}
	})
}
