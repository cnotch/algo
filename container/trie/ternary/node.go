// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ternary

// Value 前缀树值类型
type Value = interface{}

// node 前缀树节点
type node struct {
	keyChar  byte
	hasValue bool
	value    Value
	lt       *node
	eq       *node
	gt       *node
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

func (n *node) minimum() (min *node) {
	if n.lt != nil {
		return n.lt.minimum()
	}
	if n.HasValue() {
		return n
	}
	if n.eq != nil {
		return n.eq.minimum()
	}
	if n.gt != nil {
		return n.gt.minimum()
	}
	return nil
}

func (n *node) maximum() (max *node) {
	if n.gt != nil {
		return n.gt.maximum()
	}
	if n.eq != nil {
		return n.eq.maximum()
	}
	if n.HasValue() {
		return n
	}
	if n.lt != nil {
		return n.lt.maximum()
	}
	return nil
}
