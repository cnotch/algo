// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
//
// Copyright (c) 2015 Keita Sone
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

package critbit

import "strings"

// Trie critbit 前缀树
type Trie struct {
	root     node
	size     int
	maxDepth int
}

// New 创建 critbit 前缀树
func New() *Trie {
	t := Trie{
		size:     0,
		maxDepth: 0,
	}
	return &t
}

// Store 存储键值对。
// 如果键已经存在 updated 返回 true, old 返回更新前的值。
func (t *Trie) Store(key string, value Value) (old Value, updated bool) {
	n := t.getOrCreateNode(key)
	if n.external != nil {
		old, updated = n.external.value, true
		n.external.value = value
		return
	}

	n.external = &external{key: key, value: value}
	t.size++
	if t.maxDepth < len(key) {
		t.maxDepth = len(key)
	}

	return
}

// Load 加载指定 key 的值。
// 如果键存在 ok 返回 true，value 返回键对应的值。
func (t *Trie) Load(key string) (value Value, ok bool) {
	if len(key) > t.maxDepth {
		return
	}

	if n := t.search(key); n.external != nil && n.external.key == key {
		return n.external.value, true
	}
	return
}

// LoadOrStore 加载或存储键值对，并返回操作后键对应的实际值。
// 如果键存在执行加载操作，loaded 返回 true，
// 如果键不存在执行存储操作，loaded 返回 false。
func (t *Trie) LoadOrStore(key string, value Value) (actual Value, loaded bool) {
	n := t.getOrCreateNode(key)
	if n.external != nil {
		return n.external.value, true
	}

	n.external = &external{key: key, value: value}
	t.size++
	if t.maxDepth < len(key) {
		t.maxDepth = len(key)
	}
	return value, false
}

// Delete 删除指定 key 的值。
// 如果键存在，ok 返回 true，value 返回被删除键存储的值。
func (t *Trie) Delete(key string) (value Value, ok bool) {
	// an empty tree
	if t.size == 0 || len(key) > t.maxDepth {
		return
	}

	var direction int
	var whereq *node // pointer to the grandparent
	var wherep = &t.root

	// finding the best candidate to delete
	for bra := wherep.internal; bra != nil; bra = wherep.internal {
		direction = bra.direction(key)
		whereq = wherep
		wherep = &bra.child[direction]
	}

	// checking that we have the right element
	if wherep.external.key != key {
		return
	}
	value = wherep.external.value
	ok = true

	// removing the node
	if whereq == nil {
		wherep.external = nil
	} else {
		othern := whereq.internal.child[1-direction]
		whereq.internal = othern.internal
		whereq.external = othern.external
	}
	t.size--
	return
}

// Clear 清空前缀树。
func (t *Trie) Clear() {
	t.size = 0
	t.maxDepth = 0
	t.root.internal, t.root.external = nil, nil
}

// Range 用前缀树中每个键值对迭代调用函数 f。
// 如果 f 返回 false，终止迭代。
func (t *Trie) Range(f func(key string, value Value) bool) {
	if t.size == 0 {
		return
	}

	t.root.travel(f)
}

// RangePrefix 用前缀树中每个前缀为 prefix 的键值对迭代调用函数 f。
// 如果 f 返回 false，终止迭代。
func (t *Trie) RangePrefix(prefix string, f func(key string, value Value) bool) {
	n := t.prefixNode(prefix)
	if n != nil {
		n.travel(f)
	}
}

// HasPrefix 检测前缀树中是否有前缀为 prefix 的键。
func (t *Trie) HasPrefix(prefix string) bool {
	return t.prefixNode(prefix) != nil
}

// ShortestPrefixOf 获取 s 在前缀树中的最短前缀。
func (t *Trie) ShortestPrefixOf(s string) (key string, value Value, ok bool) {
	if t.size == 0 {
		return
	}
	return t.root.shortestPrefixOf(s)
}

// LongestPrefixOf 获取 s 在前缀树中的最长前缀。
func (t *Trie) LongestPrefixOf(s string) (key string, value Value, ok bool) {
	if t.size == 0 {
		return
	}
	return t.root.longestPrefixOf(s)
}

// Minimum 返回前缀树中的最小值。
func (t *Trie) Minimum() (min Value, ok bool) {
	if t.size == 0 {
		return
	}

	if n := t.root.minimum(); n != nil {
		return n.external.value, true
	}
	return
}

// Maximum 返回前缀树中的最大值。
func (t *Trie) Maximum() (max Value, ok bool) {
	if t.size == 0 {
		return
	}

	if n := t.root.maximum(); n != nil {
		return n.external.value, true
	}
	return
}

// Size 返回前缀树中键的个数。
func (t *Trie) Size() int {
	return t.size
}

// 获取或创建指定 key 的节点
func (t *Trie) getOrCreateNode(key string) *node {
	if t.size == 0 {
		return &t.root
	}

	n := t.search(key)
	newOffset, newBitmask, newCont := n.external.critbit(key)

	// already exists in the tree
	if newOffset == -1 {
		return n
	}

	// allocate new node
	newBranch := &internal{
		offset:  newOffset,
		bitmask: newBitmask,
		cont:    newCont,
	}
	direction := newBranch.direction(key)

	// insert new node
	wherep := &t.root
	for bra := wherep.internal; bra != nil; bra = wherep.internal {
		if bra.offset > newOffset || (bra.offset == newOffset && bra.bitmask < newBitmask) {
			break
		}
		wherep = &bra.child[bra.direction(key)]
	}

	if wherep.internal != nil {
		newBranch.child[1-direction].internal = wherep.internal
	} else {
		newBranch.child[1-direction].external = wherep.external
		wherep.external = nil
	}
	wherep.internal = newBranch
	return &newBranch.child[direction]
}

// 获取最可能包含对应 key 键值对的节点
func (t *Trie) search(key string) *node {
	curr := &t.root
	for bra := curr.internal; bra != nil; bra = curr.internal {
		curr = &bra.child[bra.direction(key)]
	}
	return curr
}

// 获取指定前缀的节点
func (t *Trie) prefixNode(prefix string) *node {
	if t.size == 0 {
		return nil
	}

	curr := &t.root
	wherep := curr
	if len(prefix) > 0 {
		for bra := curr.internal; bra != nil; bra = curr.internal {
			curr = &bra.child[bra.direction(prefix)]
			if bra.offset < len(prefix) {
				wherep = curr
			}
		}

		// check prefix
		if !strings.HasPrefix(curr.external.key, prefix) {
			return nil
		}
	}
	return wherep
}
