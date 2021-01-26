// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package compress

import (
	"sort"
)

const (
	maxPrefixLen = 14 // 节点最大前缀
	growMultiple = 4  // 缓冲增长倍数
)

// Value 前缀树值类型
type Value = interface{}

// node 前缀树节点
type node struct {
	prefixLen byte
	prefix    [maxPrefixLen]byte
	hasValue  bool
	value     Value
	children  []*node
}

func (n *node) IsPrefix(s string) bool {
	return len(s) >= int(n.prefixLen) && s[:n.prefixLen] == string(n.prefix[:n.prefixLen])
}
func (n *node) HasValue() bool {
	return n.hasValue
}

func (n *node) Value() (Value, bool) {
	return n.value, n.hasValue
}

func (n *node) setValue(value Value) (old Value, updated bool) {
	old, updated = n.value, n.hasValue
	n.value, n.hasValue = value, true
	return
}

func (n *node) clearValue() (value Value, ok bool) {
	value, n.value = n.value, value
	ok, n.hasValue = n.hasValue, false
	return
}

func (n *node) Size() int {
	return len(n.children)
}

func (n *node) minimum() (min *node) {
	curr := n
	for {
		if curr.HasValue() {
			return curr
		}
		if curr.Size() == 0 {
			return nil
		}
		curr = curr.children[0]
	}
}

func (n *node) maximum() (max *node) {
	curr := n
	for {
		if curr.HasValue() {
			max = curr
		}

		if curr.Size() == 0 {
			break
		}
		curr = curr.children[curr.Size()-1]
	}
	return max
}

func (n *node) child(keyChar byte) *node {
	cc := n.Size() // 子节点数

	if cc <= 64 {
		for i := 0; i < cc; i++ {
			if n.children[i].prefix[0] == keyChar {
				return n.children[i]
			}
		}
	} else {
		i := sort.Search(cc, func(i int) bool {
			return n.children[i].prefix[0] >= keyChar
		})

		if i < cc && n.children[i].prefix[0] == keyChar {
			return n.children[i]
		}
	}

	return nil
}

func (n *node) addChild(child *node) {
	keyChar := child.prefix[0]

	if n.children == nil {
		n.children = make([]*node, 1, growMultiple)
		n.children[0] = child
		return
	}

	cc := n.Size() // 子节点数
	i := sort.Search(cc, func(i int) bool {
		return n.children[i].prefix[0] >= keyChar
	})

	if cc < cap(n.children) { // 无需扩展 Slice
		n.children = n.children[0 : cc+1] // len+1
		copy(n.children[i+1:cc+1], n.children[i:cc])
	} else {
		nodes := make([]*node, cc+1, cap(n.children)*growMultiple)
		copy(nodes[0:i], n.children[0:i])
		copy(nodes[i+1:], n.children[i:cc])
		n.children = nodes
	}
	n.children[i] = child
}

func (n *node) delChild(keyChar byte) {
	cc := n.Size() // 子节点数
	i := sort.Search(cc, func(i int) bool {
		return n.children[i].prefix[0] >= keyChar
	})

	if i < cc && n.children[i].prefix[0] == keyChar {
		copy(n.children[i:cc-1], n.children[i+1:cc])
		n.children[cc-1] = nil
		n.children = n.children[0 : cc-1] // len-1

		// if len(n.children)*growMultiple < cap(n.children) {
		// 	nodes := make([]*node, cc-1, cap(n.children)/growMultiple)
		// 	copy(nodes, n.children)
		// 	n.children = nodes
		// }
	}
}

func (n *node) travel(visit func(key string, value Value) bool, buf []byte, depth int) bool {
	// 复制前缀
	copy(buf[depth:], n.prefix[:n.prefixLen])
	depth += int(n.prefixLen)
	if n.HasValue() {
		if !visit(string(buf[:depth]), n.value) {
			return false
		}
	}

	cont := true
	for _, child := range n.children {
		if cont = child.travel(visit, buf, depth); !cont {
			break
		}
	}
	return cont
}

// 分离节点的前缀
func (n *node) split(i int) {
	child := &node{prefixLen: n.prefixLen - uint8(i)}
	copy(child.prefix[0:], n.prefix[i:n.prefixLen])

	child.children, n.children = n.children, nil
	child.value, n.value = n.value, child.value
	child.hasValue, n.hasValue = n.hasValue, false
	n.prefixLen = uint8(i)
	n.addChild(child)
}

// 合并唯一的子节点
func (n *node) mergeUniqueChild() bool {
	if n.HasValue() || n.Size() != 1 { // 有值或不止一个子节点不允许合并
		return false
	}

	child := n.children[0]
	if n.prefixLen+child.prefixLen > maxPrefixLen { // 超过前缀限制
		return false
	}

	copy(n.prefix[n.prefixLen:], child.prefix[0:child.prefixLen])
	n.prefixLen += child.prefixLen
	n.children, child.children = child.children, nil
	n.value, child.value = child.value, n.value
	n.hasValue, child.hasValue = child.hasValue, false
	return true
}
