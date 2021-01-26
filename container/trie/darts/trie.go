// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package darts

import (
	"sort"
	"sync"

	"github.com/cnotch/algo/container/queue"
)

const (
	rootStatus           = 1    // 根状态；status 0 不可用，否则 check 无法处理
	minBase              = 2    // 最小偏移基址
	endCode              = 0    // key 关键字结束符编码
	codeOffset           = 1    // key 字节编码偏移量
	codeCount            = 257  // key 字节编码总数
	growMultiple         = 1.5  // 缓冲增长倍数
	percentageOfNonempty = 0.95 // 非空比率
)

func code(c byte) int {
	return int(c) + codeOffset
}

// Value 前缀树值类型
type Value = interface{}

type entry struct {
	key   string
	value Value
}

// Trie double-array 前缀树
type Trie struct {
	// base 偏移基址
	//	>0 当前状态已被占用，子节点的偏移基址
	//	=0 空闲
	//	<0 entries 的索引号
	base []int

	// check 检测上一个状态
	check []int

	// es 实体条目
	es []entry

	// nextCheckPos 查找未使用的状态检测位置
	nextCheckPos int
}

// New 从选项创建 double-array 前缀树
func New(options ...Option) *Trie {
	t := &Trie{}
	for _, opt := range options {
		opt.apply(t)
	}

	t.rearrange()
	// 初始化状态空间
	t.grow((len(t.es)+1)*2 + codeCount)
	t.nextCheckPos = minBase
	t.build()
	return t
}

// Store  Unsupported operations.
func (t *Trie) Store(key string, value Value) (old Value, updated bool) {
	panic("Unsupported operations")
}

// Load 加载指定 key 的值。
// 如果键存在 ok 返回 true，value 返回键对应的值。
func (t *Trie) Load(key string) (value Value, ok bool) {
	prefixStatus := t.matchPrefix(key)
	if prefixStatus < rootStatus {
		return
	}

	status := t.base[prefixStatus] + endCode
	currbase := t.base[status]
	if t.check[status] == prefixStatus && currbase < 0 {
		return t.es[-currbase-1].value, true
	}
	return
}

// LoadOrStore Unsupported operations.
func (t *Trie) LoadOrStore(key string, value Value) (actual Value, loaded bool) {
	panic("Unsupported operations")
}

// Delete Unsupported operations.
func (t *Trie) Delete(key string) (value Value, ok bool) {
	panic("Unsupported operations")
}

// Clear Unsupported operations.
func (t *Trie) Clear() {
	panic("Unsupported operations")
}

// Range 用前缀树中每个键值对迭代调用函数 f。
// 如果 f 返回 false，终止迭代。
func (t *Trie) Range(f func(key string, value Value) bool) {
	for _, e := range t.es {
		if !f(e.key, e.value) {
			return
		}
	}
}

// RangePrefix 用前缀树中每个前缀为 prefix 的键值对迭代调用函数 f。
// 如果 f 返回 false，终止迭代。
func (t *Trie) RangePrefix(prefix string, f func(key string, value Value) bool) {
	if len(prefix) == 0 {
		t.Range(f)
		return
	}

	// 获取包含前缀的状态
	prefixStatus := t.matchPrefix(prefix)
	if prefixStatus < rootStatus {
		return
	}
	t.travel(prefixStatus, f)
}

// HasPrefix 检测前缀树中是否有前缀为 prefix 的键。
func (t *Trie) HasPrefix(prefix string) bool {
	status := t.matchPrefix(prefix)
	return status > rootStatus || (status == rootStatus && len(t.es) > 0)
}

// ShortestPrefixOf 获取 s 在前缀树中的最短前缀。
func (t *Trie) ShortestPrefixOf(s string) (key string, value Value, ok bool) {
	if len(s) == 0 { // 空对空
		if index, ok := t.value(rootStatus); ok {
			return t.es[index].key, t.es[index].value, ok
		}
		return
	}

	status := rootStatus
	sc := len(t.base) // 状态数量
	for i := 0; i < len(s); i++ {
		next := t.base[status] + code(s[i])
		if next < sc && status == t.check[next] {
			if index, ok := t.value(next); ok {
				return t.es[index].key, t.es[index].value, ok
			}
			status = next
		} else {
			return
		}
	}
	return
}

// LongestPrefixOf 获取 s 在前缀树中的最长前缀。
func (t *Trie) LongestPrefixOf(s string) (key string, value Value, ok bool) {
	longestPrefixIndex := -1
	status := rootStatus

	// nil key check
	if index, ok := t.value(status); ok {
		longestPrefixIndex = index
	}

	sc := len(t.base) // 状态数量
	for i := 0; i < len(s); i++ {
		next := t.base[status] + code(s[i])
		if next < sc && status == t.check[next] {
			if index, ok := t.value(next); ok {
				longestPrefixIndex = index
			}
			status = next
		} else {
			break
		}
	}

	if longestPrefixIndex != -1 {
		return t.es[longestPrefixIndex].key, t.es[longestPrefixIndex].value, true
	}
	return
}

// Minimum 返回前缀树中的最小值。
func (t *Trie) Minimum() (min Value, ok bool) {
	if len(t.es) == 0 {
		return
	}
	return t.es[0].value, true
}

// Maximum 返回前缀树中的最大值。
func (t *Trie) Maximum() (max Value, ok bool) {
	if len(t.es) == 0 {
		return
	}
	return t.es[len(t.es)-1].value, true
}

// Size 返回前缀树中键的个数。
func (t *Trie) Size() int {
	return len(t.es)
}

func (t *Trie) matchPrefix(prefix string) (status int) {
	status = rootStatus
	sc := len(t.base) // 状态数量
	for i := 0; i < len(prefix); i++ {
		next := t.base[status] + code(prefix[i])
		if next < sc && status == t.check[next] {
			status = next
		} else {
			return -1
		}
	}
	return
}

func (t *Trie) travel(status int, f func(key string, value Value) bool) bool {
	begin := t.base[status] + codeOffset
	end := begin + codeCount
	if end > len(t.base) {
		end = len(t.base)
	}

	// status 是否有叶节点
	if index, ok := t.value(status); ok {
		if !f(t.es[index].key, t.es[index].value) {
			return false
		}
	}

	// 非页节点
	for i := begin; i < end; i++ {
		if t.check[i] == status {
			if !t.travel(i, f) {
				return false
			}
		}
	}
	return true
}

func (t *Trie) value(status int) (index int, ok bool) {
	terminatorStatus := t.base[status] + endCode
	if t.check[terminatorStatus] == status && t.base[terminatorStatus] < 0 {
		return -t.base[terminatorStatus] - 1, true
	}
	return
}

func (t *Trie) grow(n int) int {
	c := cap(t.base)
	size := int(growMultiple*float64(c)) + n
	newBase := make([]int, size)
	newCheck := make([]int, size)
	copy(newBase, t.base)
	copy(newCheck, t.check)
	t.base = newBase
	t.check = newCheck
	return size
}

// rearrange 对实体条目进行整理重排，确保升序和没有重复的 key
func (t *Trie) rearrange() {
	sort.Slice(t.es, func(i, j int) bool {
		return t.es[i].key < t.es[j].key
	})

	for i := len(t.es) - 1; i > 0; i-- {
		if t.es[i].key == t.es[i-1].key {
			copy(t.es[i-1:], t.es[i:])
			t.es = t.es[:len(t.es)-1]
		}
	}
}

func (t *Trie) build() {
	var q queue.Queue

	// 获取 0 列数据
	rootChilds := t.getNodes(node{
		status: rootStatus,
		depth:  0,
		begin:  0,
		end:    len(t.es),
	})
	if len(rootChilds.childs) == 0 {
		return
	}

	q.Enqueue(rootChilds)
	for q.Len() > 0 {
		e, _ := q.Dequeue()
		curr := e.(*nodes)

		// 获取子节点状态的偏移地址
		base := t.getBase(curr)
		// 修改父状态的偏移地址
		t.base[curr.parentStatus] = base
		for i := 0; i < len(curr.childs); i++ {
			n := &curr.childs[i]
			n.status = base + n.code              // 更新子节点的状态值
			t.check[n.status] = curr.parentStatus // 设置父状态值

			// 处理子节点
			if n.code == endCode { // Keyword terminator code
				t.base[n.status] = -(n.begin + 1)
			} else {
				q.Enqueue(t.getNodes(*n))
			}
		}

		// 放到池中
		curr.parentStatus = 0
		curr.childs = curr.childs[:0]
		nodesPool.Put(curr)
	}
}

func (t *Trie) getBase(l *nodes) (base int) {
	var pos int
	minCode, number := l.numberOfStates()
	if minCode+minBase > t.nextCheckPos {
		pos = minCode + minBase
	} else {
		pos = t.nextCheckPos
	}

	nonZeroNum := 0
	first := true
OUTER:
	for ; ; pos++ {
		// check memory
		if pos+number > len(t.base) {
			t.grow(pos + number - len(t.base))
		}

		if t.check[pos] != 0 {
			nonZeroNum++
			continue
		} else if first {
			t.nextCheckPos = pos
			first = false
		}

		base = pos - minCode
		for i := 0; i < len(l.childs); i++ {
			n := &l.childs[i]
			if t.check[base+n.code] != 0 {
				continue OUTER
			}
		}
		break // 找到
	}

	// -- Simple heuristics --
	// if the percentage of non-empty contents in check between the
	// index
	// 'next_check_pos' and 'check' is greater than some constant value
	// (e.g. 0.9),
	// new 'next_check_pos' index is written by 'check'.
	if 1.0*float64(nonZeroNum)/float64(pos-t.nextCheckPos+1) >= percentageOfNonempty {
		t.nextCheckPos = pos
	}

	return
}

type node struct {
	code       int
	depth      int
	begin, end int
	status     int
}
type nodes struct {
	parentStatus int
	childs       []node
}

func (l *nodes) append(code, depth, begin, end int) {
	l.childs = append(l.childs, node{
		code:  code,
		depth: depth,
		begin: begin,
		end:   end,
	})
}

// 需要的状态区间大小
func (l *nodes) numberOfStates() (minCode, number int) {
	if len(l.childs) == 0 {
		return 0, 0
	}
	return l.childs[0].code, l.childs[len(l.childs)-1].code - l.childs[0].code + 1
}

var nodesPool = sync.Pool{
	New: func() interface{} {
		return new(nodes)
	},
}

func (t *Trie) getNodes(n node) *nodes {
	l := nodesPool.Get().(*nodes)
	l.parentStatus = n.status // 设置子节点的父状态

	i := n.begin
	if i < n.end && len(t.es[i].key) == n.depth { // Keyword terminator
		l.append(endCode, n.depth+1, i, i+1)
		i++
	}

	var currBegin int
	currCode := -1
	for ; i < n.end; i++ {
		code := code(t.es[i].key[n.depth])
		if currCode != code {
			if currCode != -1 {
				l.append(currCode, n.depth+1, currBegin, i)
			}
			currCode = code
			currBegin = i
		}
	}
	if currCode != -1 {
		l.append(currCode, n.depth+1, currBegin, i)
	}
	return l
}

// MemStats base 和 check 内存的状态。
// mallocs 返回切片的大小，frees 返回空闲的大小。
func (t *Trie) MemStats() (mallocs, frees int) {
	mallocs = len(t.base)
	for i := 1; i < len(t.base); i++ {
		if t.base[i] == 0 {
			frees++
		}
	}
	return
}
