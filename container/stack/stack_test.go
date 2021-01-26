// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package stack_test

import (
	"testing"

	"github.com/cnotch/algo/container/stack"
)

func TestStack_Empty(t *testing.T) {
	var s stack.Stack
	if !s.Empty() {
		t.Error("Failed Stack.Empty")
	}
}

func TestStack_Push(t *testing.T) {
	var s stack.Stack
	s.Push(1)
	s.Push("test")
	if s.Len() != 2 {
		t.Error("Failed Stack.Push")
	}
}

func TestStack_Pop(t *testing.T) {
	var s stack.Stack
	if _, err := s.Pop(); err == nil {
		t.Error("Failed Stack.Top")
	}
	s.Push("test")
	s.Push(3)
	if value, _ := s.Pop(); !(value == 3 && s.Len() == 1) {
		t.Errorf("Failed Stack.Pop, value is %d, len is %d", value, s.Len())
	}
}
