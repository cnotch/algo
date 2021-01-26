// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package trie

// Value 前缀树值类型
type Value = interface{}

// Interface 前缀树接口
type Interface interface {
	// Store 存储键值对。
	// 如果键已经存在 updated 返回 true, old 返回更新前的值。
	Store(key string, value Value) (old Value, updated bool)

	// Load 加载指定 key 的值。
	// 如果键存在 ok 返回 true，value 返回键对应的值。
	Load(key string) (value Value, ok bool)

	// LoadOrStore 加载或存储键值对，并返回操作后键对应的实际值。
	// 如果键存在执行加载操作，loaded 返回 true，
	// 如果键不存在执行存储操作，loaded 返回 false。
	LoadOrStore(key string, value Value) (actual Value, loaded bool)

	// Delete 删除指定 key 的值。
	// 如果键存在，ok 返回 true，value 返回被删除键存储的值。
	Delete(key string) (value Value, ok bool)

	// Clear 清空前缀树。
	Clear()

	// Range 用前缀树中每个键值对迭代调用函数 f。
	// 如果 f 返回 false，终止迭代。
	Range(f func(key string, value Value) bool)

	// RangePrefix 用前缀树中每个前缀为 prefix 的键值对迭代调用函数 f。
	// 如果 f 返回 false，终止迭代。
	RangePrefix(prefix string, f func(key string, value Value) bool)

	// HasPrefix 检测前缀树中是否有前缀为 prefix 的键。
	HasPrefix(prefix string) bool

	// ShortestPrefixOf 获取 s 在前缀树中的最短前缀。
	ShortestPrefixOf(s string) (key string, value Value, ok bool)

	// LongestPrefixOf 获取 s 在前缀树中的最长前缀。
	LongestPrefixOf(s string) (key string, value Value, ok bool)

	// Minimum 返回前缀树中的最小值。
	Minimum() (min Value, ok bool)

	// Maximum 返回前缀树中的最大值。
	Maximum() (max Value, ok bool)

	// Size 返回前缀树中键的个数。
	Size() int
}
