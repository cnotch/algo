// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package queue

import (
	"errors"
	"io"
)

// Value 队列元素值类型
type Value = interface{}

const maxLen = int(^uint(0) >> 5)

// 错误变量
var (
	ErrEmptyQueue = errors.New("Queue is empty")
	ErrTooLarge   = errors.New("Queue too large")
)

// Queue 队列数据结构
type Queue struct {
	buf []Value // contents  buf[off : len(buf)]
	off int     // read at &buf[off], write at &buf[len(buf)]
}

// New 创建新的队列
func New(buf []Value) *Queue {
	return &Queue{buf: buf}
}

func makeSlice(n int) []Value {
	// If the make fails, give a known error.
	defer func() {
		if recover() != nil {
			panic(ErrTooLarge)
		}
	}()
	return make([]Value, n)
}

// 尽早早通知GC，避免内存长期占用
func resetSlice(sl []Value) {
	var nilValue Value
	for i := 0; i < len(sl); i++ {
		sl[i] = nilValue
	}
}

// Empty 判断队列是否为空
func (q *Queue) Empty() bool { return len(q.buf) <= q.off }

// Buffer 获取队列的缓存，它包含队列中所有的对象
func (q *Queue) Buffer() []Value { return q.buf[q.off:] }

// Len 队列长度
func (q *Queue) Len() int { return len(q.buf) - q.off }

// Cap 队列容量
func (q *Queue) Cap() int { return cap(q.buf) }

// Reset 重置队列
func (q *Queue) Reset() {
	resetSlice(q.buf[q.off:])

	q.buf = q.buf[:0]
	q.off = 0
}

// 尝试在cap范围内扩展队列buf
func (q *Queue) tryGrowByReslice(n int) (int, bool) {
	if l := len(q.buf); n <= cap(q.buf)-l {
		q.buf = q.buf[:l+n]
		return l, true
	}
	return 0, false
}

// 扩展队列buf
func (q *Queue) grow(n int) int {
	m := q.Len()
	// 队列为空
	if m == 0 && q.off != 0 {
		q.Reset()
	}

	// 不移动off，队列buf 容量足够
	if i, ok := q.tryGrowByReslice(n); ok {
		return i
	}

	c := cap(q.buf)
	if n <= c/2-m {
		copy(q.buf, q.buf[q.off:]) // 移动off->0
		resetSlice(q.buf[m:])      // 释放移动copy后剩余的指针引用
	} else if c > maxLen-c-n {
		panic(ErrTooLarge)
	} else {
		// 没有充足的空间，需分配
		buf := makeSlice(2*c + n)
		copy(buf, q.buf[q.off:])
		resetSlice(q.buf[q.off:]) // 重新分配了缓冲，释放旧缓冲的指针引用
		q.buf = buf
	}

	// 恢复 off和buf 值
	q.off = 0
	q.buf = q.buf[:m+n]
	return m
}

// Grow 扩展队列长度使之能容纳接下来n个元素
func (q *Queue) Grow(n int) {
	if n < 0 {
		panic("Queue.Grow: negative count")
	}
	m := q.grow(n)
	q.buf = q.buf[:m]
}

// EnqueueN 入列多个元素
func (q *Queue) EnqueueN(elems []Value) (n int, err error) {
	m, ok := q.tryGrowByReslice(len(elems))
	if !ok {
		m = q.grow(len(elems))
	}
	return copy(q.buf[m:], elems), nil
}

// Enqueue 入列
func (q *Queue) Enqueue(elem Value) error {
	m, ok := q.tryGrowByReslice(1)
	if !ok {
		m = q.grow(1)
	}
	q.buf[m] = elem
	return nil
}

// Skip 跳过指定个数的元素
func (q *Queue) Skip(size int) int {
	if q.Empty() {
		// Queue is empty, reset to recover space.
		q.Reset()
		return 0
	}

	len := q.Len()
	if size > len {
		size = len
	}
	resetSlice(q.buf[q.off : q.off+size]) // 释放已跳过的指针引用

	q.off += size
	return size
}

// DequeueN 出列多个元素
func (q *Queue) DequeueN(elems []Value) (n int, err error) {
	if q.Empty() {
		// Queue is empty, reset to recover space.
		q.Reset()
		if len(elems) == 0 {
			return 0, nil
		}
		return 0, io.EOF
	}

	n = copy(elems, q.buf[q.off:])
	resetSlice(q.buf[q.off : q.off+n]) // 释放已经出列的指针引用

	q.off += n
	return n, nil
}

// Dequeue 出列
func (q *Queue) Dequeue() (v Value, err error) {
	if q.Empty() {
		// Queue is empty, reset to recover space.
		q.Reset()
		err = io.EOF
		return
	}

	v = q.buf[q.off]
	resetSlice(q.buf[q.off : q.off+1]) // 释放已经出列的指针引用
	q.off++
	// TODO 如果buf空闲过大，可以考虑分配小缓冲，释放大缓冲
	return
}
