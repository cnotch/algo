package sort_test

import (
	"math"
	"math/rand"
	"reflect"
	"runtime"
	"sort"
	stdstrings "strings"
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

var ints = [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
var float64s = [...]float64{74.3, 59.0, math.Inf(1), 238.2, -784.0, 2.3, math.NaN(), math.NaN(), math.Inf(-1), 9845.768, -959.7485, 905, 7.8, 7.8}
var strings = [...]string{"", "Hello", "foo", "bar", "foo", "f00", "%*&^*&^&", "***"}

type sortTestCase struct {
	name     string
	original interface{}
	data     sort.Interface
}
type sortTest struct {
	sortFunc func(data sort.Interface)
	cases    []sortTestCase
}

func newSortTest(sortFunc func(data sort.Interface)) sortTest {
	st := sortTest{
		sortFunc: sortFunc,
	}
	// 获取被测试的排序函数名
	fname := runtime.FuncForPC(reflect.ValueOf(sortFunc).Pointer()).Name()
	sep := stdstrings.LastIndex(fname, "/")
	fname = fname[sep+1:]

	iclone := ints
	st.cases = append(st.cases, sortTestCase{
		fname + "(ints)",
		ints,
		sort.IntSlice(iclone[0:]),
	})

	fclone := float64s
	st.cases = append(st.cases, sortTestCase{
		fname + "(float64s)",
		float64s,
		sort.Float64Slice(fclone[0:]),
	})

	sclone := strings
	st.cases = append(st.cases, sortTestCase{
		fname + "(strings)",
		strings,
		sort.StringSlice(sclone[0:]),
	})
	return st
}

func TestSort(t *testing.T) {
	tests := []sortTest{
		newSortTest(bubble.Sort),
		newSortTest(insertion.Sort),
		newSortTest(merge.Sort),
		newSortTest(heap.Sort),
		newSortTest(quick.Sort),
		newSortTest(selection.Sort),
		newSortTest(shell.Sort),
		newSortTest(intro.Sort),
	}
	for _, test := range tests {
		for _, tt := range test.cases {
			t.Run(tt.name, func(t *testing.T) {
				test.sortFunc(tt.data)
				if !sort.IsSorted(tt.data) {
					t.Errorf("sorted %v", tt.original)
					t.Errorf("   got %v", tt.data)
				}
			})
		}
	}
}

type sliceTestCase struct {
	name     string
	original interface{}
	slice    interface{}
	less     func(i, j int) bool
}

type sliceTest struct {
	sliceFunc func(data interface{}, less func(i, j int) bool)
	cases     []sliceTestCase
}

func newSliceTest(sliceFunc func(data interface{}, less func(i, j int) bool)) sliceTest {
	st := sliceTest{
		sliceFunc: sliceFunc,
	}
	// 获取被测试的排序函数名
	fname := runtime.FuncForPC(reflect.ValueOf(sliceFunc).Pointer()).Name()
	sep := stdstrings.LastIndex(fname, "/")
	fname = fname[sep+1:]

	iclone := ints
	islice := iclone[0:]
	st.cases = append(st.cases, sliceTestCase{
		fname + "(ints)",
		ints,
		islice,
		func(i, j int) bool { return islice[j] < islice[i] },
	})

	fclone := float64s
	fslice := fclone[0:]
	st.cases = append(st.cases, sliceTestCase{
		fname + "(float64s)",
		float64s,
		fslice,
		func(i, j int) bool { return fslice[j] < fslice[i] || math.IsNaN(fslice[j]) && !math.IsNaN(fslice[i]) },
	})

	sclone := strings
	sslice := sclone[0:]
	st.cases = append(st.cases, sliceTestCase{
		fname + "(strings)",
		strings,
		sslice,
		func(i, j int) bool { return sslice[j] < sslice[i] },
	})
	return st
}

func TestSlice(t *testing.T) {
	tests := []sliceTest{
		newSliceTest(bubble.Slice),
		newSliceTest(insertion.Slice),
		newSliceTest(merge.Slice),
		newSliceTest(heap.Slice),
		newSliceTest(quick.Slice),
		newSliceTest(selection.Slice),
		newSliceTest(shell.Slice),
		newSliceTest(intro.Slice),
	}
	for _, test := range tests {
		for _, tt := range test.cases {
			t.Run(tt.name, func(t *testing.T) {
				test.sliceFunc(tt.slice, tt.less)
				if !sort.SliceIsSorted(tt.slice, tt.less) {
					t.Errorf("sorted %v", tt.original)
					t.Errorf("   got %v", tt.slice)
				}
			})
		}
	}
}

func TestSortInts(t *testing.T) {
	sorts := []func(data sort.Interface){
		bubble.Sort,
		insertion.Sort,
		merge.Sort,
		heap.Sort,
		quick.Sort,
		selection.Sort,
		shell.Sort,
		intro.Sort,
	}

	const count = 2121
	rnd := rand.New(rand.NewSource(1))
	for _, sortf := range sorts {
		list := make([]int, count)
		for i := range list {
			list[i] = rnd.Int()
		}
		sortf(sort.IntSlice(list))
		if !sort.IntsAreSorted(list) {
			t.Errorf("%s sort fail", reflect.TypeOf(sortf).Name())
		}
	}
}
