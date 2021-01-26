// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package unbounded

import (
	"github.com/cnotch/algo/container/queue"
)

// Value 树存储的值类型
type Value = interface{}

// Tree 分支无限制的有根树
type Tree struct {
	// 双亲树
	parent *Tree
	// 左孩子
	leftChild *Tree
	// 右兄弟
	rightBrother *Tree
	// 树节点存储的值
	Value Value
}

// New 实例化 Tree
func New(v Value) *Tree {
	return &Tree{Value: v}
}

// IsRoot 是否是跟节点.
func (t *Tree) IsRoot() bool {
	return t.parent == nil
}

// IsParentOf 判断是否是 x 的双亲.
func (t *Tree) IsParentOf(x *Tree) bool {
	return x.parent != nil && t == x.parent
}

// IsChildOf 判断是否是 x 的孩子.
func (t *Tree) IsChildOf(x *Tree) bool {
	return t.parent != nil && t.parent == x
}

// IsAncestorOf 判断是否是 x 的祖先.
func (t *Tree) IsAncestorOf(x *Tree) bool {
	for p := x.parent; p != nil; p = p.parent {
		if p == t {
			return true
		}
	}
	return false
}

// IsDescendantOf 判断是否是 x 的子孙.
func (t *Tree) IsDescendantOf(x *Tree) bool {
	return x.IsAncestorOf(t)
}

// IsBrotherOf 判断是否是 x 的兄弟.
func (t *Tree) IsBrotherOf(x *Tree) bool {
	return t.parent != nil && x.parent == t.parent
}

// Level 返回层次（0: 表示是跟）.
func (t *Tree) Level() int {
	lvl := -1
	p := t
	for p != nil {
		lvl++
		p = p.parent
	}
	return lvl
}

// Count Returns the number of child tree children.
func (t *Tree) Count(includeSubTrees bool) int {
	num := 0
	if includeSubTrees {
		t.count(&num)
		return num
	}

	for p := t.leftChild; p != nil; p = p.rightBrother {
		num++
	}
	return num
}

func (t *Tree) count(num *int) {
	for p := t.leftChild; p != nil; p = p.rightBrother {
		*num++
		p.count(num)
	}
}

// Root 返回跟.
func (t *Tree) Root() *Tree {
	r := t
	for p := t.parent; p != nil; p = p.parent {
		r = p
	}
	return r
}

// Parent 返回双亲.
func (t *Tree) Parent() *Tree {
	return t.parent
}

// Next 返回下一个兄弟分支或 nil.
// 时间复杂性 O(1).
func (t *Tree) Next() *Tree {
	if p := t.rightBrother; t.parent != nil {
		return p
	}
	return nil
}

// Prev 返回上一个兄弟分支或 nil.
// 时间复杂性 O(n), n 为 t 的兄弟数量.
func (t *Tree) Prev() *Tree {
	if t.parent == nil {
		return nil
	}

	var prev *Tree
	for p := t.parent.leftChild; p != nil; prev, p = p, p.rightBrother {
		if p == t {
			break
		}
	}
	return prev
}

// FirstChild 返回 first 孩子分支.
// 时间复杂性 O(1).
func (t *Tree) FirstChild() *Tree {
	return t.leftChild
}

// LastChild 返回 last 孩子分支.
// 时间复杂性 O(n), n 为 t 的孩子数量.
func (t *Tree) LastChild() *Tree {
	var last *Tree
	for p := t.leftChild; p != nil; last, p = p, p.rightBrother {
	}
	return last
}

// Remove 移除子孙分支 x,并返回它存储的值 x.Value.
// 必须 x != nil;
// 该操作仅将 x 和双亲脱钩, x 本身包含的值和分支仍然有效; 可继续操作或插入到其他树中;
// 时间复杂性 O(n), n 为 x 的兄弟数量.
func (t *Tree) Remove(x *Tree) Value {
	if t.IsAncestorOf(x) {
		x.parent.remove(x) // 在其双亲中做最终的移除操作
	}
	return x.Value
}

// AddFirst 添加值为 v 的 first 孩子并返回新分支.
// 时间复杂性 O(1).
func (t *Tree) AddFirst(v Value) *Tree {
	return t.insertValue(v, nil)
}

// AddLast 添加值为 v 的 last 孩子并返回新分支.
// 时间复杂性 O(n), n 为 t 的孩子数量.
func (t *Tree) AddLast(v Value) *Tree {
	return t.insertValue(v, t.LastChild())
}

// Add 在 at 后插入值为 v 的兄弟分支并返回。
// 必须 at != nil;
// 如果 at 不是 t 的子孙分支，直接返回 nil
// 时间复杂性 O(1).
func (t *Tree) Add(v Value, at *Tree) *Tree {
	if t.IsAncestorOf(at) {
		return at.parent.insertValue(v, at)
	}
	return nil
}

// InsertFirst 插入 x 作为 t 的 first 孩子并返回 x 或 nil.
// x 必须不为 nil;
// 如果 x.parent != nil 或 x == t.Root(), 返回 nil;
// 时间复杂性 O(1).
func (t *Tree) InsertFirst(x *Tree) *Tree {
	if x.parent != nil || x == t.Root() {
		return nil
	}
	return t.insert(x, nil)
}

// InsertLast 插入 x 作为 t 的 last 孩子并返回 x 或 nil.
// x 必须不为 nil;
// 如果 x.parent != nil 或 x == t.Root(), 返回 nil;
// 时间复杂性 O(n), n 为 t 的孩子数量.
func (t *Tree) InsertLast(x *Tree) *Tree {
	if x.parent != nil || x == t.Root() {
		return nil
	}
	return t.insert(x, t.LastChild())
}

// Insert 在 at 后插入兄弟分支 x 并返回 x 或 nil.
// x != nil && at != nil;
// 如果 x.parent != nil 或 x == t.Root(),返回 nil;
// 如果 at != nil 而且 at 不是 t 的子孙分支,直接返回 nil;
// 时间复杂性 O(1).
func (t *Tree) Insert(x, at *Tree) *Tree {
	if x.parent != nil || x == t.Root() || !t.IsAncestorOf(at) {
		return nil
	}
	return at.parent.insert(x, at)
}

// MoveToFirst 移动 x 作为 t 的 first 孩子.
// 必须 x != nil;
// 如果 x 不是 t 的子孙分支,直接返回 nil;
// 时间复杂性 O(n), n 为 x 的兄弟数量.
func (t *Tree) MoveToFirst(x *Tree) {
	if !t.IsAncestorOf(x) {
		return
	}

	if t == x.parent {
		t.move(x, nil)
	} else {
		x.parent.remove(x)
		t.insert(x, nil)
	}
}

// MoveToLast 移动 x 作为 last 分支.
// 必须 x != nil;
// 如果 x 不是 t 的子孙分支,直接返回;
// 时间复杂性 O(n+m), n 为 x 的兄弟数量, m 为 t 的孩子数量.
func (t *Tree) MoveToLast(x *Tree) {
	if !t.IsAncestorOf(x) {
		return
	}

	if t == x.parent {
		t.move(x, t.LastChild())
	} else {
		x.parent.remove(x)
		t.insert(x, t.LastChild())
	}
}

// Move 移动 x 到 at 后.
// 必须 x != nil && at != nil;
// 如果 x、at 不是 t 的分支 或 x == at,直接返回;
// 时间复杂性 O(n), n 为 x 的兄弟数量.
func (t *Tree) Move(x, at *Tree) {
	if x == at || !t.IsAncestorOf(x) || !t.IsAncestorOf(at) {
		return
	}

	if x.parent == at.parent {
		at.parent.move(x, at)
	} else {
		x.parent.remove(x)
		at.parent.insert(x, at)
	}
}

// insert 在 at 后插入 x 并返回 x.
func (t *Tree) insert(x, at *Tree) *Tree {
	if at == nil {
		x.rightBrother = t.leftChild
		t.leftChild = x
	} else {
		x.rightBrother = at.rightBrother
		at.rightBrother = x
	}
	x.parent = t
	return x
}

// insert 的便捷包装.
func (t *Tree) insertValue(v Value, at *Tree) *Tree {
	return t.insert(&Tree{Value: v}, at)
}

// remove 移除 x 分支并返回 x.
func (t *Tree) remove(x *Tree) *Tree {
	prev := x.Prev()
	if prev == nil {
		t.leftChild = x.rightBrother
	} else {
		prev.rightBrother = x.rightBrother
	}

	x.rightBrother = nil // avoid memory leaks
	x.parent = nil
	return x
}

// move 移动 x 到 at 后,并返回 x.
func (t *Tree) move(x, at *Tree) *Tree {
	if x == at {
		return x
	}

	if at == nil {
		if x == t.leftChild {
			return x
		}

		prev := x.Prev()
		prev.rightBrother = x.rightBrother
		x.rightBrother = t.leftChild
		t.leftChild = x
	} else {
		prev := x.Prev()
		if prev == nil {
			t.leftChild = x.rightBrother
		} else {
			prev.rightBrother = x.rightBrother
		}
		x.rightBrother = at.rightBrother
		at.rightBrother = x
	}

	return x
}

// Clear Removes all child tree branches .
func (t *Tree) Clear() {
	if t.leftChild == nil {
		return
	}

	var next *Tree
	for p := t.leftChild; p != nil; p = next {
		next = p.rightBrother
		p.Clear()
		p.parent = nil // avoid memory leaks
		p.rightBrother = nil
	}
	t.leftChild = nil
}

// DepthTravel 深度优先遍历.
func (t *Tree) DepthTravel(visit func(t *Tree)) {
	visit(t)
	for p := t.leftChild; p != nil; p = p.rightBrother {
		p.DepthTravel(visit)
	}
}

// BreadthTravel 广度优先遍历(按层扫描).
func (t *Tree) BreadthTravel(visit func(t *Tree)) {
	var q queue.Queue
	q.Enqueue(t)
	for q.Len() > 0 {
		i, _ := q.Dequeue()
		curr := i.(*Tree)
		visit(curr)

		for p := curr.leftChild; p != nil; p = p.rightBrother {
			q.Enqueue(p)
		}
	}
}
