// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package compress

// Trie 压缩前缀树
type Trie struct {
	size     int
	root     node
	maxDepth int
}

// New 创建压缩前缀树
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
	curr := t.getOrCreateNode(key)
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
	if len(key) > t.maxDepth {
		return
	}

	curr := &t.root
	for i := 0; i < len(key); i += int(curr.prefixLen) {
		curr = curr.child(key[i])
		if curr == nil || !curr.IsPrefix(key[i:]) {
			return
		}
	}
	return curr.Value()
}

// LoadOrStore 加载或存储键值对，并返回操作后键对应的实际值。
// 如果键存在执行加载操作，loaded 返回 true，
// 如果键不存在执行存储操作，loaded 返回 false。
func (t *Trie) LoadOrStore(key string, value Value) (actual Value, loaded bool) {
	curr := t.getOrCreateNode(key)

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
	if len(key) > t.maxDepth {
		return
	}

	if len(key) == 0 { // nil 处理
		value, ok = t.root.clearValue()
		if ok {
			t.size--
		}
		return
	}

	curr := &t.root
	retainNode := &t.root // 应保留的节点
	retainNodeDepth := 0  // 应保留节点的深度

	var retainParentNode *node

	for i := 0; i < len(key); i += int(curr.prefixLen) {
		n := curr.child(key[i])
		if n == nil || !n.IsPrefix(key[i:]) {
			return
		}

		// 检测是否要保留
		// 1. 不是当前要删除的 key
		// 2. 有超过 1 个孩子 或 存储了值
		if i < len(key)-int(n.prefixLen) &&
			(n.Size() > 1 || n.HasValue()) {
			retainNode = n
			retainNodeDepth = i + int(n.prefixLen)
			retainParentNode = curr
		}

		curr = n
	}

	if value, ok = curr.clearValue(); !ok { // 匹配节点未存储值，只是前缀
		return
	}
	t.size--

	// 收缩处理
	cc := curr.Size()
	if cc == 0 { // 没有子节点
		// 删除
		retainNode.delChild(key[retainNodeDepth])
	} else if cc == 1 { // 1 个子节点
		// 合并
		curr.mergeUniqueChild()
	} else { // >1个子节点
		return
	}

	// 尝试合并
	if retainNode != &t.root {
		retainNode.mergeUniqueChild()
	}
	if retainParentNode != nil && retainParentNode != &t.root {
		retainParentNode.mergeUniqueChild()
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
	prefixNode, prefixMatchLen := t.matchPrefix(prefix)
	if prefixNode == nil {
		return
	}

	buf := make([]byte, t.maxDepth)
	copy(buf, prefix[0:len(prefix)-prefixMatchLen])
	depth := len(prefix) - prefixMatchLen
	prefixNode.travel(f, buf, depth)
}

// HasPrefix 检测前缀树中是否有前缀为 prefix 的键。
func (t *Trie) HasPrefix(prefix string) bool {
	prefixNode, _ := t.matchPrefix(prefix)
	return prefixNode != nil
}

// ShortestPrefixOf 获取 s 在前缀树中的最短前缀。
func (t *Trie) ShortestPrefixOf(s string) (key string, value Value, ok bool) {
	curr := &t.root
	if curr.HasValue() {
		return s[:0], curr.value, true
	}

	for i := 0; i < len(s); i += int(curr.prefixLen) {
		curr = curr.child(s[i])
		if curr == nil || !curr.IsPrefix(s[i:]) {
			break
		}

		if curr.HasValue() {
			return s[:i+int(curr.prefixLen)], curr.value, true
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
		offset += len(curr.prefix)
	}

	for i := 0; i < len(s); i += int(curr.prefixLen) {
		curr = curr.child(s[i])
		if curr == nil || !curr.IsPrefix(s[i:]) {
			break
		}

		if curr.HasValue() {
			longestPrefixNode = curr
			offset = i + int(curr.prefixLen)
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

// 获取或创建指定 key 的节点
func (t *Trie) getOrCreateNode(key string) (n *node) {
	n = &t.root
	i := 0
	for i < len(key) {
		keyChar := key[i]

		child := n.child(keyChar)
		if child == nil { // 全新节点，取最大的前缀
			break
		}

		matchLen := longestCommonPrefix(key[i:], child.prefix[:child.prefixLen]) //child.match(key, i)
		// 没有全匹配
		if matchLen < int(child.prefixLen) {
			child.split(matchLen) // 拆分当前子节点
		}
		i += matchLen
		n = child
	}

	for i < len(key) {
		prefixLen := min(len(key)-i, maxPrefixLen)

		child := &node{prefixLen: uint8(prefixLen)}
		copy(child.prefix[0:], key[i:i+prefixLen])
		n.addChild(child)

		i += prefixLen
		n = child
	}
	return
}

// 获取匹配指定前缀的节点
func (t *Trie) matchPrefix(prefix string) (match *node, matchLen int) {
	if len(prefix) > t.maxDepth {
		return
	}

	match = &t.root
	for i := 0; i < len(prefix); i += matchLen {
		match = match.child(prefix[i])
		if match == nil {
			return nil, 0
		}

		matchLen = longestCommonPrefix(prefix[i:], match.prefix[:match.prefixLen])
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 两个字串最长公共前缀，默认第一个字节相同
func longestCommonPrefix(s1 string, s2 []byte) int {
	limit := min(len(s1), len(s2))
	i := 1 // 第一个字节已经匹配
	for ; i < limit; i++ {
		if s1[i] != s2[i] {
			return i
		}
	}
	return i
}
