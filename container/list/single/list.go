// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package single

// Value 链表元素的值类型
type Value = interface{}

// Element 链表的元素类型
type Element struct {
	// 当前元素在链表中的下一个元素的指针。
	next *Element

	// 元素所在链表
	list *List

	// 元素存储的值
	Value Value
}

// Next 返回下一个的链表元素或 nil
func (e *Element) Next() *Element {
	if p := e.next; e.list != nil {
		return p
	}
	return nil
}

// prev 返回上一个链表元素或 nil
// 时间复杂性 O(n)
func (e *Element) prev() *Element {
	if e.list == nil {
		return nil
	}

	for prev, p := &e.list.root, e.list.root.next; p != nil; prev, p = p, p.next {
		if p == e {
			return prev
		}
	}
	return nil
}

// Remove 从链表移除元素
func (e *Element) Remove() Value {
	if e.list != nil {
		e.list.remove(e, e.prev())
	}
	return e.Value
}

// List 单向链表类型
type List struct {
	root Element  // 哨兵元素
	tail *Element // 尾元素指针，初始状态存储 &l.root
	len  int      // 链表除哨兵元素外的长度
}

// Init 初始化或清空 l.
func (l *List) Init() *List {
	l.root.next = nil
	l.tail = &l.root
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
	return l.tail
}

// PopFront 弹出链表首元素
func (l *List) PopFront() *Element {
	if l.len == 0 {
		return nil
	}
	return l.remove(l.root.next, &l.root)
}

// PopBack 弹出链表尾元素
func (l *List) PopBack() *Element {
	if l.len == 0 {
		return nil
	}
	return l.remove(l.tail, l.tail.prev())
}

// Remove 从链表中移除 e，并返回它存储的值 e.Value。
// 必须 e != nil
func (l *List) Remove(e *Element) Value {
	if e.list == l {
		l.remove(e, e.prev())
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
	return l.insertValue(v, l.tail)
}

// InsertBefore 在链表中 mark 前插入值为 v 的元素并返回。
// 如果 mark 不是链表 l 的元素，直接返回 nil；同时必须 mark != nil
func (l *List) InsertBefore(v Value, mark *Element) *Element {
	if mark.list != l {
		return nil
	}
	return l.insertValue(v, mark.prev())
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
	if e.list != l || l.tail == e {
		return
	}
	l.move(e, l.tail)
}

// MoveBefore 移动元素 e 作为 mark 的上一个元素。
// 如果 e、mark 不是链表 l 的元素或 e == mark，直接返回；同时必须 e != nil && mark != nil
func (l *List) MoveBefore(e, mark *Element) {
	if e.list != l || mark.list != l || e == mark {
		return
	}
	l.move(e, mark.prev())
}

// MoveAfter 移动元素 e 作为 mark 的上一个元素。
// 如果 e、mark 不是链表 l 的元素或 e == mark，直接返回；同时必须 e != nil && mark != nil
func (l *List) MoveAfter(e, mark *Element) {
	if e.list != l || mark.list != l || e == mark {
		return
	}
	l.move(e, mark)
}

// PushFrontList 把 other 链表插入到链表 l 的首部。
// l 和 other 可能是同一个链表，但必须 other != nil
func (l *List) PushFrontList(other *List) {
	l.lazyInit()
	p := &l.root
	for i, e := other.Len(), other.Front(); e != nil && i > 0; i, e = i-1, e.Next() {
		p = l.insertValue(e.Value, p)
	}
}

// PushBackList 把 other 链表插入到链表 l 的尾部。
// l 和 other 可能是同一个链表，但必须 other != nil
func (l *List) PushBackList(other *List) {
	l.lazyInit()
	for i, e := other.Len(), other.Front(); e != nil && i > 0; i, e = i-1, e.Next() {
		l.insertValue(e.Value, l.tail)
	}
}

// lazyInit 惰性初始化
func (l *List) lazyInit() {
	if l.tail == nil {
		l.Init()
	}
}

// insert 在 at 后插入元素 e，增长 l.len 并返回 e
func (l *List) insert(e, at *Element) *Element {
	n := at.next
	at.next = e
	e.next = n
	e.list = l
	l.len++

	if e.next == nil { // 在尾部添加
		l.tail = e
	}
	return e
}

// insertValue insert(&Element{Value: v}, at) 的便捷包装
func (l *List) insertValue(v Value, at *Element) *Element {
	return l.insert(&Element{Value: v}, at)
}

// remove 从链表中移除 e，缩减 l.len 并返回 e
func (l *List) remove(e, prev *Element) *Element {
	prev.next = e.next
	e.next = nil // avoid memory leaks
	e.list = nil
	l.len--

	if prev.next == nil { // 移除的是尾元素
		l.tail = prev
	}
	return e
}

// move 移动 e 到 at 的下一个，并返回 e
func (l *List) move(e, at *Element) *Element {
	if e == at {
		return e
	}

	// 取出&插入 e
	prev := e.prev()
	prev.next = e.next
	n := at.next
	at.next = e
	e.next = n

	if e.next == nil { // 被移到尾部
		l.tail = e
	} else if prev.next == nil { // 从尾部移出
		l.tail = prev
	}
	return e
}
