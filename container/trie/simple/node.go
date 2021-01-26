// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package simple

import (
	"sort"
)

const growMultiple = 4 // 缓冲增长倍数

// Value 前缀树值类型
type Value = interface{}

// node 前缀树节点
type node struct {
	keyChar  byte
	hasValue bool
	value    Value
	children []*node
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
	min = n
	for {
		if min.HasValue() {
			return
		}
		if len(min.children) == 0 {
			return nil
		}
		min = min.children[0]
	}
}

func (n *node) maximum() (max *node) {
	curr := n
	for {
		if curr.HasValue() {
			max = curr
		}

		if len(curr.children) == 0 {
			break
		}
		curr = curr.children[len(curr.children)-1]
	}
	return max
}

func (n *node) child(keyChar byte) *node {
	cc := n.Size() // 子节点数

	if cc <= 64 {
		for i := 0; i < cc; i++ {
			if n.children[i].keyChar == keyChar {
				return n.children[i]
			}
		}
	} else {
		i := sort.Search(cc, func(i int) bool {
			return n.children[i].keyChar >= keyChar
		})

		if i < cc && n.children[i].keyChar == keyChar {
			return n.children[i]
		}
	}

	return nil
}

func (n *node) addChild(child *node) {
	keyChar := child.keyChar

	if n.children == nil {
		n.children = make([]*node, 1, growMultiple)
		n.children[0] = child
		return
	}

	cc := n.Size() // 子节点数
	i := sort.Search(cc, func(i int) bool {
		return n.children[i].keyChar >= keyChar
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
		return n.children[i].keyChar >= keyChar
	})

	if i < cc && n.children[i].keyChar == keyChar {
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
	if n.HasValue() {
		if !visit(string(buf[:depth]), n.value) {
			return false
		}
	}

	if len(n.children) == 0 {
		return true
	}

	cont := true
	for _, child := range n.children {
		buf[depth] = child.keyChar
		if cont = child.travel(visit, buf, depth+1); !cont {
			break
		}
	}
	return cont
}
