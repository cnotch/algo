// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package stat_test

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/cnotch/algo/math/stat"
)

var ints []int

const (
	count = 1e4
	k     = count / 2
)

func init() {

	ints = make([]int, count)
	for i := 0; i < count; i++ {
		ints[i] = rand.Intn(count)
	}
}

func BenchmarkSelect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		list := append([]int(nil), ints...)
		b.StartTimer()
		stat.Select(sort.IntSlice(list), k)
	}
}

func BenchmarkSliceSelect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		list := append([]int(nil), ints...)
		b.StartTimer()
		stat.SliceSelect(list, k, func(i, j int) bool { return list[i] < list[j] })
	}
}

func sortSelection(list []int, k int) int {
	sort.Ints(list)
	return list[k]
}

func TestSelect(t *testing.T) {
	rnd := rand.New(rand.NewSource(1))

	const count = 2121
	for k := 0; k < count; k++ {
		list := make([]int, count)
		for i := range list {
			list[i] = rnd.Intn(1000)
		}
		stat.Select(sort.IntSlice(list), k)
		sorted := append([]int(nil), list...)
		want := sortSelection(sorted, k)
		if list[k] != want {
			t.Errorf("unexpected result from Select(..., %d): got:%v want:%d", k, list[k], want)
		}
	}
}

func TestSliceSelect(t *testing.T) {
	rnd := rand.New(rand.NewSource(1))

	const count = 2121
	for k := 0; k < count; k++ {
		list := make([]int, count)
		for i := range list {
			list[i] = rnd.Intn(1000)
		}
		stat.SliceSelect(list, k, func(i, j int) bool { return list[i] < list[j] })
		sorted := append([]int(nil), list...)
		want := sortSelection(sorted, k)
		if list[k] != want {
			t.Errorf("unexpected result from Select(..., %d): got:%v want:%d", k, list[k], want)
		}
	}
}

func TestMinMax(t *testing.T) {
	rnd := rand.New(rand.NewSource(1))

	const count = 2121
	for k := 0; k < count; k++ {
		list := make([]int, count)
		for i := range list {
			list[i] = rnd.Intn(1000)
		}

		min, max := stat.MinMax(count, func(i, j int) bool { return list[i] < list[j] })
		sorted := append([]int(nil), list...)
		sort.Ints(sorted)
		if list[min] != sorted[0] || list[max] != sorted[count-1] {
			t.Errorf("unexpected result from MinMax(): gotMin:%d wantMin:%d;gotMax:%d wantMax:%d", list[min], sorted[0], list[max], sorted[count-1])
		}
	}
}
