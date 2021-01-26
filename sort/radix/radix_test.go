package radix_test

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/cnotch/algo/sort/radix"
)

var ints []int
var strings []string
var students []student
var pstudents []*student

type student struct {
	name  string
	score int
}

func init() {
	const count = 100
	const maxval = 99
	maxd := maxval
	d := 0
	for maxd > 0 {
		d++
		maxd /= 10
	}

	ints = make([]int, count)
	strings = make([]string, count)
	students = make([]student, count)
	pstudents = make([]*student, count)

	for i := 0; i < len(ints); i++ {
		ints[i] = rand.Intn(maxval + 1)
		// 随机字串
		var bs []byte
		chars := rand.Intn(d + 1)
		for j := 0; j < chars; j++ {
			bs = append(bs, byte(rand.Intn(26)+97))
		}
		strings[i] = string(bs)
		students[i] = student{strings[i], ints[i]}
		pstudents[i] = &student{strings[i], ints[i]}
	}
}

func TestSort(t *testing.T) {
	t.Run("radix.Sort(ints)-sort.Stable", func(t *testing.T) {
		data := append([]int(nil), ints...)
		d := radix.MaxDecDigits(len(data), func(i int) int { return data[i] })
		radix.Sort(d,
			func(p int) {
				sort.SliceStable(data, func(i, j int) bool {
					return radix.DecDigit(data[i], p) < radix.DecDigit(data[j], p)
				})
			})
		if !sort.IntsAreSorted(data) {
			t.Errorf("sorted %v", ints)
			t.Errorf("   got %v", data)
		}
	})

	t.Run("radix.Sort(strings)-sort.Stable", func(t *testing.T) {
		data := append([]string(nil), strings...)
		d := radix.MaxDigits(len(data), func(i int) int { return len(data[i]) })
		radix.Sort(d,
			func(p int) {
				sort.SliceStable(data, func(i, j int) bool {
					return radix.CharDigit(data[i], d, p) < radix.CharDigit(data[j], d, p)
				})
			})
		if !sort.StringsAreSorted(data) {
			t.Errorf("sorted %v", strings)
			t.Errorf("   got %v", data)
		}
	})
}

func TestInts(t *testing.T) {
	t.Run("radix.Ints", func(t *testing.T) {
		data := append([]int(nil), ints...)
		d := radix.MaxDecDigits(len(data), func(i int) int { return data[i] })
		radix.Ints(data, d)
		if !sort.IntsAreSorted(data) {
			t.Errorf("sorted %v", ints)
			t.Errorf("   got %v", data)
		}
	})
}

func TestStrings(t *testing.T) {
	t.Run("radix.Strings", func(t *testing.T) {
		data := append([]string(nil), strings...)
		d := radix.MaxDigits(len(data), func(i int) int { return len(data[i]) })
		radix.Strings(data, d)
		if !sort.StringsAreSorted(data) {
			t.Errorf("sorted %v", strings)
			t.Errorf("   got %v", data)
		}
	})
}

func TestSlice(t *testing.T) {
	t.Run("radix.Slice(ints)", func(t *testing.T) {
		data := append([]int(nil), ints...)
		d := radix.MaxDecDigits(len(data), func(i int) int { return data[i] })
		radix.Slice(data, d,
			func(slice interface{}, p int) (k int, key func(i int) int) {
				temp := slice.([]int)
				k = 9
				key = func(i int) int {
					return radix.DecDigit(temp[i], p)
				}
				return
			})
		if !sort.IntsAreSorted(data) {
			t.Errorf("sorted %v", ints)
			t.Errorf("   got %v", data)
		}
	})

	t.Run("radix.Slice(strings)", func(t *testing.T) {
		data := append([]string(nil), strings...)
		d := radix.MaxDigits(len(data), func(i int) int { return len(data[i]) })

		radix.Slice(data, d,
			func(slice interface{}, p int) (k int, key func(i int) int) {
				temp := slice.([]string)
				k = 255
				key = func(i int) int {
					return radix.CharDigit(temp[i], d, p)
				}
				return
			})
		if !sort.StringsAreSorted(data) {
			t.Errorf("sorted %v", strings)
			t.Errorf("   got %v", data)
		}
	})

	t.Run("radix.Slice(students)", func(t *testing.T) {
		data := append([]student(nil), students...)
		d := radix.MaxDecDigits(len(data), func(i int) int { return data[i].score })
		radix.Slice(data, d,
			func(slice interface{}, p int) (k int, key func(i int) int) {
				temp := slice.([]student)
				k = 9
				key = func(i int) int {
					return radix.DecDigit(temp[i].score, p)
				}
				return
			})
		if !sort.SliceIsSorted(data, func(i, j int) bool { return data[i].score < data[j].score }) {
			t.Errorf("sorted %v", ints)
			t.Errorf("   got %v", data)
		}
	})

	t.Run("radix.Slice(pstudents)", func(t *testing.T) {
		data := append([]*student(nil), pstudents...)
		d := radix.MaxDecDigits(len(data), func(i int) int { return data[i].score })
		radix.Slice(data, d,
			func(slice interface{}, p int) (k int, key func(i int) int) {
				temp := slice.([]*student)
				k = 9
				key = func(i int) int {
					return radix.DecDigit(temp[i].score, p)
				}
				return
			})
		if !sort.SliceIsSorted(data, func(i, j int) bool { return data[i].score < data[j].score }) {
			t.Errorf("sorted %v", ints)
			t.Errorf("   got %v", data)
		}
	})
}
