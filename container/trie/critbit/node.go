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

import (
	"strings"
)

var msbMatrix [256]byte // 最高有效位掩码表

func init() {
	var x, y uint
	for x = 0; x <= 8; x++ {
		for ; y < (1 << x); y++ {
			msbMatrix[y] = 0x80 >> (8 - x)
		}
	}
}

// 前缀树节点
type node struct {
	internal *internal // 内部节点
	external *external // 外部节点：键值对
}

// 内部节点
type internal struct {
	offset  int  // key[offset]
	bitmask byte // key[offset] & bitmask
	cont    bool // if true, key of child[1] contains key of child[0]
	child   [2]node
}

// Value 前缀树值类型
type Value = interface{}

// 外部节点，存储键值对
type external struct {
	key   string
	value Value
}

// n.key 和 key 的差异位
func (n *external) critbit(key string) (offset int, bitmask byte, cont bool) {
	nlen := len(n.key)
	plen := len(key)
	limit := nlen
	if nlen > plen {
		limit = plen
	}

	for offset = 0; offset < limit; offset++ {
		if a, b := key[offset], n.key[offset]; a != b {
			bitmask = msbMatrix[a^b]
			return
		}
	}

	if nlen > plen {
		bitmask = msbMatrix[n.key[offset]]
	} else if plen > nlen {
		bitmask = msbMatrix[key[offset]]
	} else { // e.key == key
		offset = -1
	}
	return offset, bitmask, true
}

// 获取 key 在分支中的下一个节点的方向
func (n *internal) direction(key string) int {
	if n.offset < len(key) && (key[n.offset]&n.bitmask != 0 || n.cont) {
		return 1
	}
	return 0
}

func (n *node) travel(visit func(key string, value Value) bool) bool {
	if n.internal != nil {
		// dealing with an internal node while recursing
		for i := 0; i < 2; i++ {
			if !n.internal.child[i].travel(visit) {
				return false
			}
		}
	} else {
		return visit(n.external.key, n.external.value)
	}

	return true
}

func (n *node) shortestPrefixOf(s string) (key string, value Value, ok bool) {
	if n.external != nil {
		if strings.HasPrefix(s, n.external.key) {
			return n.external.key, n.external.value, true
		}
		return
	}

	direction := n.internal.direction(s)
	if direction == 1 {
		if key, value, ok = n.internal.child[0].shortestPrefixOf(s); ok {
			return
		}
	}
	return n.internal.child[direction].shortestPrefixOf(s)
}

func (n *node) longestPrefixOf(s string) (key string, value Value, ok bool) {
	if n.internal != nil {
		direction := n.internal.direction(s)
		if key, value, ok = n.internal.child[direction].longestPrefixOf(s); ok {
			return
		}
		if direction == 1 {
			return n.internal.child[0].longestPrefixOf(s)
		}
	} else {
		if strings.HasPrefix(s, n.external.key) {
			return n.external.key, n.external.value, true
		}
	}
	return
}

func (n *node) minimum() (min *node) {
	if n.external != nil {
		return n
	}
	return n.internal.child[0].minimum()
}

func (n *node) maximum() (max *node) {
	if n.internal != nil {
		return n.internal.child[1].maximum()
	}
	return n
}
