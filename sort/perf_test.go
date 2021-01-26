package sort_test

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/cnotch/algo/sort/bubble"
	"github.com/cnotch/algo/sort/heap"
	"github.com/cnotch/algo/sort/insertion"
	"github.com/cnotch/algo/sort/intro"
	"github.com/cnotch/algo/sort/merge"
	"github.com/cnotch/algo/sort/quick"
	"github.com/cnotch/algo/sort/selection"
	"github.com/cnotch/algo/sort/shell"
)

var (
	randints    []int
	randstrings []string
	rptints     []int
)

const (
	count  = 1e4
	keymax = 99
)

func init() {
	randints = make([]int, count)
	rptints = make([]int, count)
	randstrings = make([]string, count)

	d := 0
	for i := keymax; i > 0; i /= 10 {
		d++
	}

	for i := 0; i < count; i++ {
		randints[i] = rand.Intn(count)
		rptints[i] = rand.Intn(keymax + 1)

		var bs []byte
		chars := rand.Intn(d + 1)
		for j := 0; j < chars; j++ {
			bs = append(bs, byte(rand.Intn(26)+97))
		}
		randstrings[i] = string(bs)
	}
}

func withBenchmarkSort(b *testing.B, ints []int, sort func(slice []int)) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		slice := append([]int(nil), ints...)
		b.StartTimer()
		sort(slice)
	}
}

func withBenchmarkParallelSort(b *testing.B, ints []int, sort func(slice []int)) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			slice := append([]int(nil), ints...)
			sort(slice)
		}
	})
}

func BenchmarkStdSort(b *testing.B) {
	withBenchmarkSort(b, randints, func(slice []int) {
		sort.Sort(sort.IntSlice(slice))
	})
}

func BenchmarkStdSlice(b *testing.B) {
	withBenchmarkSort(b, randints, func(slice []int) {
		sort.Slice(slice, func(i, j int) bool { return slice[i] < slice[j] })
	})
}
func BenchmarkIntroSort(b *testing.B) {
	withBenchmarkSort(b, randints, func(slice []int) {
		intro.Sort(sort.IntSlice(slice))
	})
}

func BenchmarkIntroSlice(b *testing.B) {
	withBenchmarkSort(b, randints, func(slice []int) {
		intro.Slice(slice, func(i, j int) bool { return slice[i] < slice[j] })
	})
}

func BenchmarkQuickSort(b *testing.B) {
	withBenchmarkSort(b, randints, func(slice []int) {
		quick.Sort(sort.IntSlice(slice))
	})
}

func BenchmarkQuickSlice(b *testing.B) {
	withBenchmarkSort(b, randints, func(slice []int) {
		quick.Slice(slice, func(i, j int) bool { return slice[i] < slice[j] })
	})
}

func BenchmarkBubbleSort(b *testing.B) {
	withBenchmarkSort(b, randints, func(slice []int) {
		bubble.Sort(sort.IntSlice(slice))
	})
}

func BenchmarkBubbleSlice(b *testing.B) {
	withBenchmarkSort(b, randints, func(slice []int) {
		bubble.Slice(slice, func(i, j int) bool { return slice[i] < slice[j] })
	})
}

func BenchmarkSelectionSort(b *testing.B) {
	withBenchmarkSort(b, randints, func(slice []int) {
		selection.Sort(sort.IntSlice(slice))
	})
}

func BenchmarkSelectionSlice(b *testing.B) {
	withBenchmarkSort(b, randints, func(slice []int) {
		selection.Slice(slice, func(i, j int) bool { return slice[i] < slice[j] })
	})
}

func BenchmarkInsertionSort(b *testing.B) {
	withBenchmarkSort(b, randints, func(slice []int) {
		insertion.Sort(sort.IntSlice(slice))
	})
}

func BenchmarkInsertionSlice(b *testing.B) {
	withBenchmarkSort(b, randints, func(slice []int) {
		insertion.Slice(slice, func(i, j int) bool { return slice[i] < slice[j] })
	})
}
func BenchmarkShellSort(b *testing.B) {
	withBenchmarkSort(b, randints, func(slice []int) {
		shell.Sort(sort.IntSlice(slice))
	})
}
func BenchmarkShellSlice(b *testing.B) {
	withBenchmarkSort(b, randints, func(slice []int) {
		shell.Slice(slice, func(i, j int) bool { return slice[i] < slice[j] })
	})
}
func BenchmarkMergeSort(b *testing.B) {
	withBenchmarkSort(b, randints, func(slice []int) {
		merge.Sort(sort.IntSlice(slice))
	})
}

func BenchmarkMergeSlice(b *testing.B) {
	withBenchmarkSort(b, randints, func(slice []int) {
		merge.Slice(slice, func(i, j int) bool { return slice[i] < slice[j] })
	})
}
func BenchmarkHeapSort(b *testing.B) {
	withBenchmarkSort(b, randints, func(slice []int) {
		heap.Sort(sort.IntSlice(slice))
	})
}

func BenchmarkHeapSlice(b *testing.B) {
	withBenchmarkSort(b, randints, func(slice []int) {
		heap.Slice(slice, func(i, j int) bool { return slice[i] < slice[j] })
	})
}
