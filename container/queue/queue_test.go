// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package queue_test

import (
	"testing"

	"github.com/cnotch/algo/container/queue"
)

func TestQueue_Enqueue(t *testing.T) {
	var q queue.Queue
	q.Enqueue(3)
	if q.Len() != 1 {
		t.Error("Failed Queue.Enqueue")
	}
}

func TestQueue_Dequeue(t *testing.T) {
	var q queue.Queue
	if _, err := q.Dequeue(); err == nil {
		t.Error("Failed Queue.Top")
	}
	q.Enqueue("test")
	q.Enqueue(3)
	if value, _ := q.Dequeue(); !(value == "test" && q.Len() == 1) {
		t.Errorf("Failed Queue.Dequeue, value is %d, len is %d", value, q.Len())
	}
}

func TestQueue_Empty(t *testing.T) {
	var q queue.Queue
	if !q.Empty() {
		t.Error("Failed Queue.Empty")
	}
}
