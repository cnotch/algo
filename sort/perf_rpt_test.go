package sort_test

import (
	"sort"
	"testing"

	"github.com/cnotch/algo/sort/bucket"
	"github.com/cnotch/algo/sort/counting"
	"github.com/cnotch/algo/sort/insertion"
	"github.com/cnotch/algo/sort/intro"
	"github.com/cnotch/algo/sort/quick"
	"github.com/cnotch/algo/sort/radix"
)

func BenchmarkRptStdSort(b *testing.B) {
	withBenchmarkSort(b, rptints, func(slice []int) {
		sort.Sort(sort.IntSlice(slice))
	})
}

func BenchmarkRptStdSlice(b *testing.B) {
	withBenchmarkSort(b, rptints, func(slice []int) {
		sort.Slice(slice, func(i, j int) bool { return slice[i] < slice[j] })
	})
}

func BenchmarkRptIntroSort(b *testing.B) {
	withBenchmarkSort(b, rptints, func(slice []int) {
		intro.Sort(sort.IntSlice(slice))
	})
}

func BenchmarkRptIntroSlice(b *testing.B) {
	withBenchmarkSort(b, rptints, func(slice []int) {
		intro.Slice(slice, func(i, j int) bool { return slice[i] < slice[j] })
	})
}

func BenchmarkRptQuickSort(b *testing.B) {
	withBenchmarkSort(b, rptints, func(slice []int) {
		quick.Sort(sort.IntSlice(slice))
	})
}

func BenchmarkRptQuickSlice(b *testing.B) {
	withBenchmarkSort(b, rptints, func(slice []int) {
		quick.Slice(slice, func(i, j int) bool { return slice[i] < slice[j] })
	})
}

func BenchmarkCountingSort(b *testing.B) {
	withBenchmarkSort(b, rptints, func(slice []int) {
		_ = make([]int, count)
		counting.Sort(count, keymax,
			func(r int) int { return rptints[r] },
			func(l, r int) { slice[l] = rptints[r] })
	})
}

func BenchmarkCountingSlice(b *testing.B) {
	withBenchmarkSort(b, rptints, func(slice []int) {
		counting.Slice(slice, keymax,
			func(r int) int { return slice[r] })
	})
}

func BenchmarkCountingUnstable(b *testing.B) {
	withBenchmarkSort(b, rptints, func(slice []int) {
		counting.Unstable(len(slice), keymax,
			func(i int) int { return slice[i] },
			func(i, j int) { slice[i], slice[j] = slice[j], slice[i] })
	})
}
func BenchmarkCountingSliceUnstable(b *testing.B) {
	withBenchmarkSort(b, rptints, func(slice []int) {
		counting.SliceUnstable(slice, keymax,
			func(i int) int { return slice[i] })
	})
}

func BenchmarkRadixInts(b *testing.B) {
	d := radix.MaxDecDigits(len(rptints), func(i int) int { return rptints[i] })
	withBenchmarkSort(b, rptints, func(slice []int) {
		radix.Ints(slice, d)
	})
}

func BenchmarkRadixInts_Std(b *testing.B) {
	d := radix.MaxDecDigits(len(rptints), func(i int) int { return rptints[i] })
	withBenchmarkSort(b, rptints, func(slice []int) {
		radix.Sort(d,
			func(p int) {
				sort.SliceStable(slice, func(i, j int) bool {
					return radix.DecDigit(slice[i], p) < radix.DecDigit(slice[j], p)
				})
			})
	})
}

func BenchmarkRadixInts_Reflect(b *testing.B) {
	d := radix.MaxDecDigits(len(rptints), func(i int) int { return rptints[i] })
	withBenchmarkSort(b, rptints, func(slice []int) {
		radix.Slice(slice, d,
			func(slice interface{}, p int) (k int, key func(i int) int) {
				temp := slice.([]int)
				k = 9
				key = func(i int) int {
					return radix.DecDigit(temp[i], p)
				}
				return
			})
	})
}

func BenchmarkRadixStrings(b *testing.B) {
	d := radix.MaxDigits(len(randstrings), func(i int) int { return len(randstrings[i]) })
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		slice := make([]string, count)
		copy(slice, randstrings)
		b.StartTimer()
		radix.Strings(slice, d)
	}
}

func BenchmarkStdStrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		slice := make([]string, count)
		copy(slice, randstrings)
		b.StartTimer()
		sort.Strings(slice)
	}
}

func BenchmarkBucketInts(b *testing.B) {
	withBenchmarkSort(b, rptints, func(slice []int) {
		bucket.Ints(slice, keymax,
			func(i int) int { return slice[i] },
			func(bucket []int) {
				// sort.SliceStable(bucket, func(i, j int) bool { return bucket[i] < bucket[j] })
				insertion.Slice(bucket, func(i, j int) bool { return bucket[i] < bucket[j] })
			})
	})
}

func BenchmarkBucketSlice(b *testing.B) {
	withBenchmarkSort(b, rptints, func(slice []int) {
		bucket.Slice(slice, keymax,
			func(i int) int { return slice[i] },
			func(slice interface{}) {
				temp := slice.([]int)
				// sort.SliceStable(temp, func(i, j int) bool { return temp[i] < temp[j]})
				insertion.Slice(temp, func(i, j int) bool { return temp[i] < temp[j] })
			})
	})
}
