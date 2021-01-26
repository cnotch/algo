// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package stack

import "errors"

// Value 栈元素值类型
type Value = interface{}

// Stack 栈数据结构
type Stack []Value

// 错误变量
var (
	ErrEmptyStack = errors.New("Stack is empty")
)

// Len 栈长度
func (s *Stack) Len() int {
	return len(*s)
}

// Cap 栈容量
func (s *Stack) Cap() int {
	return cap(*s)
}

// Empty 栈是否为空
func (s *Stack) Empty() bool {
	return len(*s) == 0
}

// Push 入栈
func (s *Stack) Push(t Value) {
	*s = append(*s, t)
}

// Pop 出栈
func (s *Stack) Pop() (v Value, err error) {
	if s.Empty() {
		err = ErrEmptyStack
		return
	}

	top := len(*s) - 1
	v, (*s)[top] = (*s)[top], v // 更早通知GC
	*s = (*s)[:top]             // 移出
	return
}
