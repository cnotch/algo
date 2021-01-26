// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package skip

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func checkListLen(t *testing.T, l *List, len int) bool {
	if n := l.Len(); n != len {
		t.Errorf("l.Len() = %d, want %d", n, len)
		return false
	}
	return true
}

func checkListPointers(t *testing.T, l *List, es []*Element) {
	if !checkListLen(t, l, len(es)) {
		return
	}

	// zero length lists must be the zero value or properly initialized (sentinel circle)
	if len(es) == 0 {
		if l.Front() != nil || l.Back() != nil {
			t.Errorf("l.front = %p, l.back = %p; both should both be nil", l.Front(), l.Back())
		}
		return
	}

	for i, p := 0, l.Front(); p != nil && i < len(es); i, p = i+1, p.Next() {
		if p != es[i] {
			t.Errorf("l[%d]: p = %p,want = %p", i, p, es[i])
		}
		if i == len(es)-1 && p.Next() != nil {
			t.Error("tail elem's next isn't nil")
		}
	}
}

func TestList(t *testing.T) {
	l := New(4, func(l, r Value) bool {
		return l.(int) < r.(int)
	})
	checkListPointers(t, l, []*Element{})

	// Single E list
	e := l.Insert(1)
	checkListPointers(t, l, []*Element{e})
	l.Remove(e)
	checkListPointers(t, l, []*Element{})

	// Bigger list
	e2 := l.Insert(2)
	e1 := l.Insert(1)
	e4 := l.Insert(4)
	e3 := l.Insert(3)
	checkListPointers(t, l, []*Element{e1, e2, e3, e4})

	l.Remove(e2)
	checkListPointers(t, l, []*Element{e1, e3, e4})

	e2 = l.Insert(2)
	checkListPointers(t, l, []*Element{e1, e2, e3, e4})
	l.Remove(e2)
	e5 := l.Insert(5)
	checkListPointers(t, l, []*Element{e1, e3, e4, e5})
	l.Remove(e5)
	checkListPointers(t, l, []*Element{e1, e3, e4})
	l.Remove(e1)
	checkListPointers(t, l, []*Element{e3, e4})
	e1 = l.Insert(1)
	checkListPointers(t, l, []*Element{e1, e3, e4})

	// Check standard iteration.
	sum := 0
	for e := l.Front(); e != nil; e = e.Next() {
		if i, ok := e.Value.(int); ok {
			sum += i
		}
	}
	if sum != 8 {
		t.Errorf("sum over l = %d, want 8", sum)
	}

	// Clear all Es by iterating
	var next *Element
	for e := l.Front(); e != nil; e = next {
		next = e.Next()
		l.Remove(e)
	}
	checkListPointers(t, l, []*Element{})
}

func checkList(t *testing.T, l *List, es []interface{}) {
	if !checkListLen(t, l, len(es)) {
		return
	}

	i := 0
	for e := l.Front(); e != nil; e = e.Next() {
		le := e.Value.(int)
		if le != es[i] {
			t.Errorf("l[%d].Value = %v, want %v", i, le, es[i])
		}
		i++
	}
}

func TestExtending(t *testing.T) {
	l1 := New(4, func(l, r Value) bool {
		return l.(int) < r.(int)
	})
	l2 := New(6, func(l, r Value) bool {
		return l.(int) < r.(int)
	})

	l1.Insert(1)
	l1.Insert(2)
	l1.Insert(3)

	l2.Insert(4)
	l2.Insert(5)

	l3 := New(3, func(l, r Value) bool {
		return l.(int) < r.(int)
	})

	l3.InsertList(l1)
	checkList(t, l3, []interface{}{1, 2, 3})
	l3.InsertList(l2)
	checkList(t, l3, []interface{}{1, 2, 3, 4, 5})

	l3 = New(3, func(l, r Value) bool {
		return l.(int) < r.(int)
	})
	l3.InsertList(l2)
	checkList(t, l3, []interface{}{4, 5})
	l3.InsertList(l1)
	checkList(t, l3, []interface{}{1, 2, 3, 4, 5})

	checkList(t, l1, []interface{}{1, 2, 3})
	checkList(t, l2, []interface{}{4, 5})

	l3 = New(3, func(l, r Value) bool {
		return l.(int) < r.(int)
	})
	l3.InsertList(l1)
	checkList(t, l3, []interface{}{1, 2, 3})
	l3.InsertList(l3)
	checkList(t, l3, []interface{}{1, 1, 2, 2, 3, 3})
}

func TestRemove(t *testing.T) {
	l := New(3, func(l, r Value) bool {
		return l.(int) < r.(int)
	})
	e1 := l.Insert(1)
	e2 := l.Insert(2)
	checkListPointers(t, l, []*Element{e1, e2})
	e := l.Front()
	l.Remove(e)
	checkListPointers(t, l, []*Element{e2})
	l.Remove(e)
	checkListPointers(t, l, []*Element{e2})
	e1 = l.Insert(1)
	e3 := l.Insert(3)
	e4 := l.Insert(4)
	checkListPointers(t, l, []*Element{e1, e2, e3, e4})
	l.PopFront()
	l.PopBack()
	checkListPointers(t, l, []*Element{e2, e3})
	l.PopFront()
	l.PopBack()
	checkListPointers(t, l, []*Element{})
	l.PopFront()
	l.PopBack()
	checkListPointers(t, l, []*Element{})

	l = New(4, func(l, r Value) bool {
		return l.(int) < r.(int)
	})

	e = nil
	ev := 0
	comp := []interface{}{}
	for i := 0; i < 32; i++ {
		temp := l.Insert(i)
		tempv := i
		comp = append(comp, i)
		if temp.level > 1 {
			e = temp
			ev = tempv
		}
	}
	checkList(t, l, comp)

	if e != nil {
		l.Remove(e)
		comp = append(comp[0:ev], comp[ev+1:]...)
		checkList(t, l, comp)
	}
}

func TestRemoveOtherListElem(t *testing.T) {
	l1 := New(3, func(l, r Value) bool {
		return l.(int) < r.(int)
	})
	l1.Insert(1)
	l1.Insert(2)

	l2 := New(3, func(l, r Value) bool {
		return l.(int) < r.(int)
	})
	l2.Insert(3)
	l2.Insert(4)

	e := l1.Front()
	l2.Remove(e) // l2 should not change because e is not an E of l2
	if n := l2.Len(); n != 2 {
		t.Errorf("l2.Len() = %d, want 2", n)
	}

	l1.Insert(8)
	if n := l1.Len(); n != 3 {
		t.Errorf("l1.Len() = %d, want 3", n)
	}
}

func TestDetachedElem(t *testing.T) {
	l := New(3, func(l, r Value) bool {
		return l.(int) < r.(int)
	})
	l.Insert(1)
	l.Insert(2)

	e := l.Front()
	l.Remove(e)
	if e.Value != 1 {
		t.Errorf("e.value = %d, want 1", e.Value)
	}
	if e.Next() != nil {
		t.Errorf("e.Next() != nil")
	}
	if e.prev() != nil {
		t.Errorf("e.Prev() != nil")
	}
}

func TestRandomLevel(t *testing.T) {
	l := New(4, func(l, r Value) bool {
		return l.(int) < r.(int)
	})
	three := 0
	for i := 0; i < 32; i++ {
		e := l.Insert(i)
		if e.level == 3 {
			three++
		}
	}
	if three == 0 {
		t.Errorf("three ==0, want >0")
	}
}

func checkRange(t *testing.T, begin, end *Element, es []interface{}) {
	i := 0
	for e := begin; e != nil && e != end; e = e.Next() {
		le := e.Value.(int)
		if le != es[i] {
			t.Errorf("l[%d].Value = %v, want %v", i, le, es[i])
		}
		i++
	}
}

func TestSearch(t *testing.T) {
	l := New(4, func(l, r Value) bool {
		return l.(int) < r.(int)
	})

	for i := 0; i < 32; i++ {
		l.Insert(i)
	}
	a, b := l.Search(0)
	checkRange(t, a, b, []interface{}{0})
	a, b = l.Search(31)
	checkRange(t, a, b, []interface{}{31})
	assert.Nil(t, b, nil)
	a, b = l.Search(2)
	checkRange(t, a, b, []interface{}{2})
	l.Insert(2)
	l.Insert(2)
	l.Insert(2)
	a, b = l.Search(2)
	checkRange(t, a, b, []interface{}{2, 2, 2, 2})
	a.Remove()
	a, b = l.Search(2)
	checkRange(t, a, b, []interface{}{2, 2, 2})

	a, b = l.Search(1)
	l.Remove(a)
	a, b = l.Search(1) // not exist
	assert.Nil(t, a, nil)
	assert.Nil(t, b, nil)

}

func TestSearchRange(t *testing.T) {
	l := New(4, func(l, r Value) bool {
		return l.(int) < r.(int)
	})

	for i := 0; i < 32; i++ {
		l.Insert(i)
	}

	a, b := l.SearchRange(CloseOpen, 0, 2)
	checkRange(t, a, b, []interface{}{0, 1})
	a, b = l.SearchRange(CloseOpen, 28, 32)
	checkRange(t, a, b, []interface{}{28, 29, 30, 31})
	assert.Nil(t, b, nil)

	a, b = l.SearchRange(CloseOpen, 0, 3)
	checkRange(t, a, b, []interface{}{0, 1, 2})

	a, b = l.SearchRange(CloseClose, 1, 4)
	checkRange(t, a, b, []interface{}{1, 2, 3, 4})
	a, b = l.SearchRange(OpenOpen, 1, 4)
	checkRange(t, a, b, []interface{}{2, 3})
	a, b = l.SearchRange(CloseOpen, 1, 4)
	checkRange(t, a, b, []interface{}{1, 2, 3})
	a, b = l.SearchRange(OpenClose, 1, 4)
	checkRange(t, a, b, []interface{}{2, 3, 4})

	a, b = l.Search(1)
	l.Remove(a)
	a, b = l.SearchRange(CloseOpen, 1, 3) // left boundary no exist
	checkRange(t, a, b, []interface{}{2, 2, 2, 2})
}
