// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package partition_test

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/cnotch/algo/sort/quick/partition"
)

var ints []int

func init() {
	const count = 1e4
	rnd := rand.New(rand.NewSource(1))

	ints = make([]int, count)
	for i := 0; i < count; i++ {
		ints[i] = rnd.Int()
	}
}

func TestPartition(t *testing.T) {
	ps := []partition.Partition{
		partition.Get(partition.Basic),
		partition.Get(partition.Randomized),
		partition.Get(partition.ThreeWay),
		partition.Get(partition.StrictThreeWay),
	}

	rnd := rand.New(rand.NewSource(1))
	for _, partition := range ps {
		list := append([]int(nil), ints...)
		for i := range list {
			list[i] = rnd.Int()
		}
		midlo, midhi := partition(sort.IntSlice(list), 0, len(list))
		for i := 0; i < midlo; i++ {
			if list[i] > list[midlo] {
				t.Errorf("unexpected partition sort order p[%d] > p[%d]: %d > %d", i, midlo, list[i], list[midlo])
			}
		}
		for i := midhi; i < len(list); i++ {
			if list[i] < list[midlo] {
				t.Errorf("unexpected partition sort order p[%d] < p[%d]: %d < %d", i, midlo, list[i], list[midlo])
			}
		}
	}
}

func withBenchmarkPartition(b *testing.B, partition partition.Partition) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		list := append([]int(nil), ints...)
		b.StartTimer()
		_, _ = partition(sort.IntSlice(list), 0, len(list))
	}
}

func BenchmarkPartitionBasic(b *testing.B) {
	withBenchmarkPartition(b, partition.Get(partition.Basic))
}

func BenchmarkPartitionRandomized(b *testing.B) {
	withBenchmarkPartition(b, partition.Get(partition.Randomized))
}

func BenchmarkPartitionThreeWay(b *testing.B) {
	withBenchmarkPartition(b, partition.Get(partition.ThreeWay))
}

func BenchmarkPartitionStrictThreeWay(b *testing.B) {
	withBenchmarkPartition(b, partition.Get(partition.StrictThreeWay))
}
