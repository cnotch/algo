// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package simple

// Trie 朴素前缀树
type Trie struct {
	size     int
	root     node
	maxDepth int
}

// New 创建朴素前缀树
func New() *Trie {
	t := Trie{
		size: 0,
	}
	return &t
}

// Store 存储键值对。
// 如果键已经存在 updated 返回 true, old 返回更新前的值。
func (t *Trie) Store(key string, value Value) (old Value, updated bool) {
	curr := t.addOrCreateNode(key)
	old, updated = curr.setValue(value)
	if !updated { // 不是修改是插入
		t.size++
	}

	if t.maxDepth < len(key) {
		t.maxDepth = len(key)
	}
	return
}

// Load 加载指定 key 的值。
// 如果键存在 ok 返回 true，value 返回键对应的值。
func (t *Trie) Load(key string) (value Value, ok bool) {
	prefixNode := t.matchPrefix(key)
	if prefixNode == nil {
		return
	}
	return prefixNode.Value()
}

// LoadOrStore 加载或存储键值对，并返回操作后键对应的实际值。
// 如果键存在执行加载操作，loaded 返回 true，
// 如果键不存在执行存储操作，loaded 返回 false。
func (t *Trie) LoadOrStore(key string, value Value) (actual Value, loaded bool) {
	curr := t.addOrCreateNode(key)

	// 没有值，设置并返回
	if actual, loaded = curr.Value(); !loaded {
		curr.setValue(value)
		actual = value
		t.size++

		if t.maxDepth < len(key) {
			t.maxDepth = len(key)
		}
	}
	return
}

// Delete 删除指定 key 的值。
// 如果键存在，ok 返回 true，value 返回被删除键存储的值。
func (t *Trie) Delete(key string) (value Value, ok bool) {
	curr := &t.root
	retainNode := &t.root // 应保留的节点
	retainNodeDepth := -1 // 应保留节点的深度

	for i := 0; i < len(key); i++ {
		n := curr.child(key[i])
		if n == nil {
			return
		}
		curr = n

		// 检测是否要保留
		// 1. 不是当前要删除的 key
		// 2. 有超过 1 个孩子 或 存储了值
		if i < len(key)-1 &&
			(curr.Size() > 1 || curr.HasValue()) {
			retainNode = n
			retainNodeDepth = i
		}
	}

	if value, ok = curr.clearValue(); !ok { // 匹配节点未存储值，只是前缀
		return
	}
	t.size--

	// 匹配节点有子节点，无需收缩处理
	if curr.Size() > 0 {
		return
	}

	// 删除需要保留的节点后续节点
	if retainNode != &t.root {
		retainNode.delChild(key[retainNodeDepth+1])
	}
	return
}

// Clear 清空前缀树。
func (t *Trie) Clear() {
	t.size = 0
	t.maxDepth = 0
	t.root = node{}
}

// Range 用前缀树中每个键值对迭代调用函数 f。
// 如果 f 返回 false，终止迭代。
func (t *Trie) Range(f func(key string, value Value) bool) {
	depth := 0
	buf := make([]byte, t.maxDepth)
	t.root.travel(f, buf, depth)
}

// RangePrefix 用前缀树中每个前缀为 prefix 的键值对迭代调用函数 f。
// 如果 f 返回 false，终止迭代。
func (t *Trie) RangePrefix(prefix string, f func(key string, value Value) bool) {
	prefixNode := t.matchPrefix(prefix)
	if prefixNode == nil {
		return
	}

	buf := make([]byte, t.maxDepth)
	copy(buf, prefix)
	depth := len(prefix)
	prefixNode.travel(f, buf, depth)
}

// HasPrefix 检测前缀树中是否有前缀为 prefix 的键。
func (t *Trie) HasPrefix(prefix string) bool {
	return t.matchPrefix(prefix) != nil
}

// ShortestPrefixOf 获取 s 在前缀树中的最短前缀。
func (t *Trie) ShortestPrefixOf(s string) (key string, value Value, ok bool) {
	curr := &t.root
	if curr.HasValue() {
		return s[:0], curr.value, true
	}

	for i := 0; i < len(s); i++ {
		curr = curr.child(s[i])
		if curr == nil {
			break
		}

		if curr.HasValue() {
			return s[:i+1], curr.value, true
		}
	}
	return
}

// LongestPrefixOf 获取 s 在前缀树中的最长前缀。
func (t *Trie) LongestPrefixOf(s string) (key string, value Value, ok bool) {
	var longestPrefixNode *node
	offset := 0
	curr := &t.root
	if curr.HasValue() {
		longestPrefixNode = curr
	}

	for i := 0; i < len(s); i++ {
		curr = curr.child(s[i])
		if curr == nil {
			break
		}

		if curr.HasValue() {
			longestPrefixNode = curr
			offset = i + 1
		}
	}

	if longestPrefixNode != nil {
		return s[:offset], longestPrefixNode.value, true
	}
	return
}

// Minimum 返回前缀树中的最小值。
func (t *Trie) Minimum() (min Value, ok bool) {
	minNode := t.root.minimum()
	if minNode != nil {
		return minNode.Value()
	}
	return
}

// Maximum 返回前缀树中的最大值。
func (t *Trie) Maximum() (max Value, ok bool) {
	maxNode := t.root.maximum()
	if maxNode != nil {
		return maxNode.Value()
	}
	return
}

// Size 返回前缀树中键的个数。
func (t *Trie) Size() int {
	return t.size
}

func (t *Trie) addOrCreateNode(key string) (n *node) {
	n = &t.root
	for i := 0; i < len(key); i++ {
		keyChar := key[i]
		child := n.child(keyChar)
		if child == nil {
			child = &node{keyChar: keyChar}
			n.addChild(child)
		}

		n = child
	}
	return
}

func (t *Trie) matchPrefix(prefix string) (n *node) {
	if len(prefix) > t.maxDepth {
		return
	}

	n = &t.root
	for i := 0; i < len(prefix); i++ {
		n = n.child(prefix[i])
		if n == nil {
			return
		}
	}
	return
}
