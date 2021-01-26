// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package bucket_test

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/cnotch/algo/sort/bucket"
	"github.com/cnotch/algo/sort/insertion"
)

var testSlice []int

func init() {
	const count = 64
	data := make([]int, count)
	for i := 0; i < count; i++ {
		data[i] = rand.Intn(100)
	}
	testSlice = data
}

func TestBucketInts(t *testing.T) {
	t.Run("bucket.Ints", func(t *testing.T) {
		ints := append([]int(nil), testSlice...)
		bucket.Ints(ints, 10,
			func(i int) int { return ints[i] / 10 },
			func(bucket []int) {
				sort.SliceStable(bucket, func(i, j int) bool { return bucket[i] < bucket[j] })
			})
		if !sort.IntsAreSorted(ints) {
			t.Errorf("sorted %v", testSlice)
			t.Errorf("   got %v", ints)
		}
	})
}

func TestBucketSlice(t *testing.T) {
	t.Run("bucket.Slice", func(t *testing.T) {
		ints := append([]int(nil), testSlice...)
		bucket.Slice(ints, 10,
			func(i int) int { return ints[i] / 10 },
			func(slice interface{}) {
				temp := slice.([]int)
				insertion.Slice(temp, func(i, j int) bool {
					return temp[i] < temp[j]
				})
			})

		if !sort.IntsAreSorted(ints) {
			t.Errorf("sorted %v", testSlice)
			t.Errorf("   got %v", ints)
		}
	})
}
