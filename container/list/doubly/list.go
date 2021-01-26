// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package doubly 是 container.list 的复制品
package doubly

// Value 链表元素的值类型
type Value = interface{}

// Element 链表的元素类型
type Element struct {
	// 当前元素在链表中上一个、下一个元素的指针。
	// 为简化实现，链表内部实现为循环链表；
	// 其中 &l.root 被作为尾元素（l.Back()）的下一个元素
	// 和首元素（l.Front()）的上一个元素。
	prev, next *Element

	// 元素所在链表
	list *List

	// 元素存储的值
	Value Value
}

// Next 返回下一个的链表元素或 nil
func (e *Element) Next() *Element {
	if p := e.next; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

// Prev 返回上一个链表元素或 nil
func (e *Element) Prev() *Element {
	if p := e.prev; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

// Remove 从链表移除元素
func (e *Element) Remove() Value {
	if e.list != nil {
		e.list.remove(e)
	}
	return e.Value
}

// List 双向链表类型
type List struct {
	root Element // 哨兵元素，仅 &root, root.prev, and root.next 被使用
	len  int     // 链表除哨兵元素外的长度
}

// Init 初始化或清空 l.
func (l *List) Init() *List {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}

// New 返回被初始化的链表
func New() *List { return new(List).Init() }

// Len 返回链表长度。
// 时间复杂度 O(1).
func (l *List) Len() int { return l.len }

// Front 返回首元素，如果链表为空返回 nil
func (l *List) Front() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

// Back 返回链表尾元素，如果链表为空返回 nil
func (l *List) Back() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}

// Remove 从链表中移除 e，并返回它存储的值 e.Value。
// 必须 e != nil
func (l *List) Remove(e *Element) Value {
	if e.list == l {
		// if e.list == l, l must have been initialized when e was inserted
		// in l or l == nil (e is a zero Element) and l.remove will crash
		l.remove(e)
	}
	return e.Value
}

// PushFront 在链表首部插入值为 v 的元素并返回。
func (l *List) PushFront(v Value) *Element {
	l.lazyInit()
	return l.insertValue(v, &l.root)
}

// PushBack 在链表尾部插入值为 v 的元素并返回。
func (l *List) PushBack(v Value) *Element {
	l.lazyInit()
	return l.insertValue(v, l.root.prev)
}

// InsertBefore 在链表中 mark 前插入值为 v 的元素并返回。
// 如果 mark 不是链表 l 的元素，直接返回 nil；同时必须 mark != nil
func (l *List) InsertBefore(v Value, mark *Element) *Element {
	if mark.list != l {
		return nil
	}
	return l.insertValue(v, mark.prev)
}

// InsertAfter 在链表中 mark 后插入值为 v 的元素并返回。
// 如果 mark 不是链表 l 的元素，直接返回 nil；同时必须 mark != nil
func (l *List) InsertAfter(v Value, mark *Element) *Element {
	if mark.list != l {
		return nil
	}
	return l.insertValue(v, mark)
}

// MoveToFront 移动元素 e 作为链表首元素。
// 如果 e 不是链表 l 的元素，直接返回 nil；同时必须 e != nil
func (l *List) MoveToFront(e *Element) {
	if e.list != l || l.root.next == e {
		return
	}
	l.move(e, &l.root)
}

// MoveToBack 移动元素 e 作为链表尾元素。
// 如果 e 不是链表 l 的元素，直接返回 nil；同时必须 e != nil
func (l *List) MoveToBack(e *Element) {
	if e.list != l || l.root.prev == e {
		return
	}
	l.move(e, l.root.prev)
}

// MoveBefore 移动元素 e 作为 mark 的上一个元素。
// 如果 e、mark 不是链表 l 的元素或 e == mark，直接返回；同时必须 e != nil && mark != nil
func (l *List) MoveBefore(e, mark *Element) {
	if e.list != l || e == mark || mark.list != l {
		return
	}
	l.move(e, mark.prev)
}

// MoveAfter 移动元素 e 作为 mark 的上一个元素。
// 如果 e、mark 不是链表 l 的元素或 e == mark，直接返回；同时必须 e != nil && mark != nil
func (l *List) MoveAfter(e, mark *Element) {
	if e.list != l || e == mark || mark.list != l {
		return
	}
	l.move(e, mark)
}

// PushFrontList 把 other 链表插入到链表 l 的首部。
// l 和 other 可能是同一个链表，但必须 other != nil
func (l *List) PushFrontList(other *List) {
	l.lazyInit()
	for i, e := other.Len(), other.Back(); i > 0; i, e = i-1, e.Prev() {
		l.insertValue(e.Value, &l.root)
	}
}

// PushBackList 把 other 链表插入到链表 l 的尾部。
// l 和 other 可能是同一个链表，但必须 other != nil
func (l *List) PushBackList(other *List) {
	l.lazyInit()
	for i, e := other.Len(), other.Front(); i > 0; i, e = i-1, e.Next() {
		l.insertValue(e.Value, l.root.prev)
	}
}

// lazyInit 惰性初始化
func (l *List) lazyInit() {
	if l.root.next == nil {
		l.Init()
	}
}

// insert 在 at 后插入元素 e，增长 l.len 并返回 e
func (l *List) insert(e, at *Element) *Element {
	n := at.next
	at.next = e
	e.prev = at
	e.next = n
	n.prev = e
	e.list = l
	l.len++
	return e
}

// insertValue insert(&Element{Value: v}, at) 的便捷包装
func (l *List) insertValue(v Value, at *Element) *Element {
	return l.insert(&Element{Value: v}, at)
}

// remove 从链表中移除 e，缩减 l.len 并返回 e
func (l *List) remove(e *Element) *Element {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = nil // avoid memory leaks
	e.prev = nil // avoid memory leaks
	e.list = nil
	l.len--
	return e
}

// move 移动 e 到 at 的下一个，并返回 e
func (l *List) move(e, at *Element) *Element {
	if e == at {
		return e
	}
	e.prev.next = e.next
	e.next.prev = e.prev

	n := at.next
	at.next = e
	e.prev = at
	e.next = n
	n.prev = e

	return e
}
