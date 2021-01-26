// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package skip

import (
	"fmt"
	"math/rand"
)

// Value 链表元素的值类型
type Value = interface{}

// SearchMode 搜索模式
type SearchMode uint8

// 搜索模式常量
const (
	CloseClose = SearchMode(iota) //[a,b] a<=x<=b 取值包括a、b
	OpenOpen                      //(a,b) a<x<b   取值不包括a、b
	CloseOpen                     //[a,b) a<=x<b  取值包括a，不包括b
	OpenClose                     //(a,b] a<x<=b  取值不包括a，包括b
)

// Element 链表的元素类型
type Element struct {
	// 元素所在 level, <=level 的层都包含该元素
	level int

	// 当前元素在链表中的各级的下一个元素的指针。
	nexts []*Element

	// 元素所在链表
	list *List

	// 元素存储的值
	Value Value
}

// Next 返回下一个的链表元素或 nil
func (e *Element) Next() *Element {
	if e.list == nil || e.nexts == nil {
		return nil
	}

	return e.nexts[0]
}

// prev 返回上一个链表元素或 nil
// 时间复杂性 O(n)
func (e *Element) prev() *Element {
	if e.list == nil {
		return nil
	}

	for prev, p := &e.list.root, e.list.root.nexts[0]; p != nil; prev, p = p, p.nexts[0] {
		if p == e {
			return prev
		}
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

// reset 重置 e , 设置所有的指针为 nil
func (e *Element) reset() {
	e.list = nil
	for i := 0; i < len(e.nexts); i++ {
		e.nexts[i] = nil
	}
	e.nexts = nil
}

// List 跳跃链表类型
type List struct {
	root   Element               // 哨兵元素
	tail   *Element              // 尾元素指针，初始状态存储 &l.root
	len    int                   // 链表除哨兵元素外的长度
	levels int                   // 跳表的层数
	less   func(l, r Value) bool // 元素值类型比较函数 l<r
}

// New 返回被初始化的链表
func New(levels int, less func(l, r Value) bool) *List {
	if levels < 1 {
		panic("levels must > 0")
	}
	l := &List{}
	return l.init(levels, less)
}

// Init 初始化 l.
func (l *List) Init(levels int, less func(l, r Value) bool) *List {
	if levels < 1 {
		panic("levels must > 0")
	}
	return l.init(levels, less)
}

// Reset 重置 l，清空所有内容
func (l *List) Reset() *List {
	// 如果已经初始化, 清空所有指针
	if len(l.root.nexts) > 0 {
		prev := &l.root
		for p := l.root.nexts[0]; p != nil; p = p.nexts[0] {
			prev.reset()
			prev = p
		}
		prev.reset()
		prev = nil
	}

	return l.init(l.levels, l.less)
}

// Len 返回链表长度。
// 时间复杂度 O(1).
func (l *List) Len() int { return l.len }

// Front 返回首元素，如果链表为空返回 nil
func (l *List) Front() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.nexts[0]
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
	return l.remove(l.root.nexts[0])
}

// PopBack 弹出链表尾元素
func (l *List) PopBack() *Element {
	if l.len == 0 {
		return nil
	}
	return l.remove(l.tail)
}

// Remove 从链表中移除 e，并返回它存储的值 e.Value。
// 必须 e != nil
func (l *List) Remove(e *Element) Value {
	if e.list == l {
		l.remove(e)
	}
	return e.Value
}

// Insert 插入值为 v 的元素并返回。
func (l *List) Insert(v Value) *Element {
	return l.insert(&Element{
		Value: v,
		level: l.randomLevel(),
		nexts: make([]*Element, l.levels),
	})
}

// InsertList 把 other 链表插入到链表 l 。
// l 和 other 可能是同一个链表，但必须 other != nil
func (l *List) InsertList(other *List) {
	for i, e := other.Len(), other.Front(); e != nil && i > 0; i, e = i-1, e.Next() {
		l.Insert(e.Value)
	}
}

// Search 在链表中搜索 v，返回[a,b)。
// 如果没有找到，a 和 b 都是 nil。如何利用[a,b)
// 	1. 获得 == v 的区间
// 		a, b := l.Search(3)
// 		for p := a; p != nil && p != b; p=p.Next() {
// 			fmt.Println("p.Value = %v", p.Value)
// 		}
// 	2. 获得 < v 的区间
//		[l.Front,a)
//	3. 获得 <= v 的区间
//		[l.Front,b)
//	4. 获得 > v 的区间
//		[b,nil)
//	5. 获得 >= v 的区间
//		[a,nil)
func (l *List) Search(v Value) (a, b *Element) {
	return l.search(v, func(comp Value) bool {
		return !l.less(v, comp) // 当前 comp<v is false
	})
}

// SearchRange 在链表中搜索 [v1,v2)，返回[a,b)。
func (l *List) SearchRange(mode SearchMode, v1, v2 Value) (a, b *Element) {
	// search 会在 comp >= v1 的情况下调用 filter 函数
	var filter func(comp Value) bool // comp>=v1

	switch mode {
	case CloseClose: // v1<= comp <=v2
		filter = func(comp Value) bool {
			return l.less(comp, v2) ||
				(!l.less(comp, v2) && !l.less(v2, comp)) // =v2
		}
	case OpenOpen: // v1< comp <v2
		filter = func(comp Value) bool {
			return l.less(v1, comp) && // 排除 v1==comp 的情况
				l.less(comp, v2)
		}
	case CloseOpen: // v1<= comp <v2
		filter = func(comp Value) bool {
			return l.less(comp, v2)
		}
	case OpenClose: // v1< comp <=v2
		filter = func(comp Value) bool {
			return (l.less(v1, comp) && l.less(comp, v2)) ||
				(!l.less(comp, v2) && !l.less(v2, comp)) // = v2
		}
	default:
		panic(fmt.Sprintf("not support mode: %d", mode))
	}

	return l.search(v1, filter)
}

// init 初始化 l.
func (l *List) init(levels int, less func(l, r Value) bool) *List {
	l.levels = levels
	l.root.level = levels - 1
	l.root.nexts = make([]*Element, levels)
	l.tail = &l.root
	l.len = 0
	l.less = less
	return l
}

func (l *List) randomLevel() int {
	level := 0
	maxLevel := l.levels - 1
	for {
		// 理想状态：level0 -1/2 level1 - 1/4 level2:1/8...
		if rand.Intn(2) == 1 || level >= maxLevel {
			break
		}
		level++
	}

	return level
}

// insert 插入元素 e，增长 l.len 并返回 e
func (l *List) insert(e *Element) *Element {
	prev := &l.root
	// 从高层开始找
	for i := l.levels - 1; i >= 0; i-- {
		for {
			next := prev.nexts[i]
			if next != nil && l.less(next.Value, e.Value) {
				prev = next
			} else {
				break
			}
		}

		// 在 i 层已经找到插入位置，在 prev 之后插入 e
		if i <= e.level {
			e.nexts[i] = prev.nexts[i]
			prev.nexts[i] = e
		}
	}

	e.list = l
	l.len++

	if e.nexts[0] == nil { // 在尾部添加
		l.tail = e
	}
	return e
}

// remove 从链表中移除 e，缩减 l.len 并返回 e
func (l *List) remove(e *Element) *Element {
	prev := &l.root

	// 找最近的 prev
	for i := l.levels - 1; i > e.level; i-- {
		for {
			next := prev.nexts[i]
			if next != nil && l.less(next.Value, e.Value) {
				prev = next
			} else {
				break
			}
		}
	}

	// 找 e 并修改链接
	for i := e.level; i >= 0; i-- {
		for {
			next := prev.nexts[i]
			if next == e {
				break
			}
			prev = next
		}
		prev.nexts[i] = e.nexts[i]
	}
	l.len--

	if e == l.tail { // 移除的是尾元素
		l.tail = prev
	}

	e.reset() // avoid memory leaks
	return e
}

func (l *List) search(boundary Value, filter func(comp Value) bool) (a, b *Element) {
	prev := &l.root
	// 从高层开始找
	for i := l.levels - 1; i >= 0; i-- {
		for {
			next := prev.nexts[i]
			if next != nil && l.less(next.Value, boundary) {
				prev = next
			} else {
				break
			}
		}
	}

	// prev.next[0] == nil || !prev.next[0].Value < v
	for {
		next := prev.nexts[0]
		if next != nil && filter(next.Value) {
			if a == nil { // 第一个相等
				a = next
			}
			prev = next
		} else { // 不在范围
			if a != nil {
				b = next
			}
			break
		}
	}

	return
}
