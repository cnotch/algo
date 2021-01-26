// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ternary

// Trie 三分前缀树
type Trie struct {
	size     int
	nilnode  node
	root     *node
	maxDepth int
}

// New 创建三叉前缀树
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
	if len(key) == 0 {
		return t.nilnode.Value()
	}

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
	if len(key) > t.maxDepth {
		return
	}

	if len(key) == 0 { // nil 处理
		value, ok = t.nilnode.clearValue()
		if ok {
			t.size--
		}
		return
	}

	var nilp *node
	var maybeShrink **node
	maybeShrink = &nilp

	i := 0
	keyChar := key[i]
	n := &t.root
Match:
	for *n != nil {
		switch {
		case keyChar < (*n).keyChar:
			n = &(*n).lt
		case keyChar > (*n).keyChar:
			n = &(*n).gt
		case i == len(key)-1:
			break Match
		default:
			if !(*n).HasValue() {
				if *maybeShrink == nil || (*n).lt != nil || (*n).gt != nil {
					maybeShrink = n
				}
			} else {
				maybeShrink = &nilp
			}

			n = &(*n).eq
			i++
			keyChar = key[i]
		}
	}

	if *n == nil {
		return
	}

	if value, ok = (*n).clearValue(); !ok { // 匹配节点未存储值，只是前缀
		return
	}
	t.size--

	if cont := t.shrinkNode(n, false); !cont {
		return
	}

	// 其他处理
	if (*maybeShrink) != nil {
		t.shrinkNode(maybeShrink, true)
	}
	return
}

// Clear 清空前缀树。
func (t *Trie) Clear() {
	*t = Trie{
		size:     0,
		maxDepth: 0,
	}
}

// Range 用前缀树中每个键值对迭代调用函数 f。
// 如果 f 返回 false，终止迭代。
//
// 注意：函数 f 的 key 参数在整个迭代过程使用同一内存，仅能作为临时使用。
// 如果需要保存留作后用，请先复制。
func (t *Trie) Range(f func(key string, value Value) bool) {
	if t.nilnode.HasValue() {
		if !f("", t.nilnode.value) {
			return
		}
	}

	buf := make([]byte, 0, t.maxDepth)
	t.travel(t.root, buf, f)
}

// RangePrefix 用前缀树中每个前缀为 prefix 的键值对迭代调用函数 f。
// 如果 f 返回 false，终止迭代。
//
// 注意：函数 f 的 key 参数在整个迭代过程使用同一内存，仅能作为临时使用。
// 如果需要保存留作后用，请先复制。
func (t *Trie) RangePrefix(prefix string, f func(key string, value Value) bool) {
	if len(prefix) == 0 {
		t.Range(f)
		return
	}

	prefixNode := t.matchPrefix(prefix)
	if prefixNode == nil {
		return
	}

	buf := make([]byte, len(prefix), t.maxDepth)
	copy(buf, prefix)
	if prefixNode.HasValue() {
		if !f(string(buf), prefixNode.value) {
			return
		}
	}

	t.travel(prefixNode.eq, buf, f)
}

// HasPrefix 检测前缀树中是否有前缀为 prefix 的键。
func (t *Trie) HasPrefix(prefix string) bool {
	if len(prefix) == 0 {
		return t.size > 0
	}

	return t.matchPrefix(prefix) != nil
}

// ShortestPrefixOf 获取 s 在前缀树中的最短前缀。
func (t *Trie) ShortestPrefixOf(s string) (key string, value Value, ok bool) {
	if t.nilnode.HasValue() {
		return "", t.nilnode.value, true
	}

	if len(s) == 0 {
		return
	}

	for i, n := 0, t.root; n != nil && i < len(s); {
		keyChar := s[i]
		switch {
		case keyChar < n.keyChar:
			n = n.lt
		case keyChar > n.keyChar:
			n = n.gt
		default:
			if n.HasValue() {
				return s[:i+1], n.value, true
			}
			n = n.eq
			i++
		}
	}

	return
}

// LongestPrefixOf 获取 s 在前缀树中的最长前缀。
func (t *Trie) LongestPrefixOf(s string) (key string, value Value, ok bool) {
	if len(s) == 0 {
		if t.nilnode.HasValue() {
			return s, t.nilnode.value, true
		}
		return
	}

	var longestPrefixNode *node
	offset := 0
	for i, n := 0, t.root; n != nil && i < len(s); {
		keyChar := s[i]
		switch {
		case keyChar < n.keyChar:
			n = n.lt
		case keyChar > n.keyChar:
			n = n.gt
		default:
			if n.HasValue() {
				longestPrefixNode = n
				offset = i + 1
			}
			n = n.eq
			i++
		}
	}

	if longestPrefixNode != nil {
		return s[:offset], longestPrefixNode.value, true
	}
	return
}

// Minimum 返回前缀树中的最小值。
func (t *Trie) Minimum() (min Value, ok bool) {
	if t.nilnode.HasValue() {
		return t.nilnode.value, true
	}

	if t.root != nil {
		minNode := t.root.minimum()
		if minNode != nil {
			return minNode.Value()
		}
	}
	return
}

// Maximum 返回前缀树中的最大值。
func (t *Trie) Maximum() (max Value, ok bool) {
	if t.root != nil {
		maxNode := t.root.maximum()
		if maxNode != nil {
			return maxNode.Value()
		}
	}

	return t.nilnode.Value()
}

// Size 返回前缀树中键的个数。
func (t *Trie) Size() int {
	return t.size
}

func (t *Trie) addOrCreateNode(key string) (n *node) {
	if len(key) == 0 {
		return &t.nilnode
	}

	i := 0
	next := &t.root
	keyChar := key[i]

	for *next != nil {
		switch {
		case keyChar < (*next).keyChar:
			next = &(*next).lt
		case keyChar > (*next).keyChar:
			next = &(*next).gt
		case i == len(key)-1:
			return *next
		default:
			next = &(*next).eq
			i++
			keyChar = key[i]
		}
	}

	for ; i < len(key); i++ {
		n = &node{keyChar: key[i]}
		*next, next = n, &n.eq
	}
	return
}

// must len(prefix)>0
func (t *Trie) matchPrefix(prefix string) (n *node) {
	i := 0
	n = t.root
	keyChar := prefix[i]
	for n != nil {
		switch {
		case keyChar < n.keyChar:
			n = n.lt
		case keyChar > n.keyChar:
			n = n.gt
		case i == len(prefix)-1:
			return
		default:
			n = n.eq
			i++
			keyChar = prefix[i]
		}
	}

	return
}

func (t *Trie) shrinkNode(n **node, ignoreEQ bool) (cont bool) {
	curr := *n
	if curr.lt == nil && curr.gt == nil && (curr.eq == nil || ignoreEQ) {
		*n = nil // 直接清空
		return true
	}

	if curr.eq != nil && !ignoreEQ { // 有后续 key
		return
	}

	if curr.lt == nil {
		*n = curr.gt
		return
	}

	if curr.gt == nil {
		*n = curr.lt
		return
	}

	// 同时存在 lt 和 gt
	lt, gt := curr.lt, curr.gt
	*n = gt
	if gt.lt == nil {
		gt.lt = lt
		return
	}

	wherep := &gt.lt.lt
	for *wherep != nil {
		wherep = &(*wherep).lt
	}
	*wherep = lt
	return
}

func (t *Trie) travel(n *node, key []byte, f func(key string, value Value) bool) bool {
	if n == nil {
		return true
	}

	if !t.travel(n.lt, key, f) {
		return false
	}

	currKey := append(key, n.keyChar)
	if n.HasValue() {
		if !f(string(currKey), n.value) {
			return false
		}
	}
	if !t.travel(n.eq, currKey, f) {
		return false
	}

	return t.travel(n.gt, key, f)
}
