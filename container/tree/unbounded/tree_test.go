// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package unbounded

import "testing"

func checkSubTreesCount(t *testing.T, tree *Tree, count int, includeSubTrees bool) bool {
	if n := tree.Count(includeSubTrees); n != count {
		t.Errorf("tree.Count(%v) = %d, want %d", includeSubTrees, n, count)
		return false
	}
	return true
}

func checkTreePointers(t *testing.T, tree *Tree, ts []*Tree) {
	if !checkSubTreesCount(t, tree, len(ts), false) {
		return
	}

	if len(ts) == 0 {
		if tree.FirstChild() != nil || tree.LastChild() != nil {
			t.Errorf("tree.FirstChild = %p, tree.LastChild = %p; both should both be nil", tree.FirstChild(), tree.LastChild())
		}
		return
	}

	for i, p := 0, tree.FirstChild(); p != nil && i < len(ts); i, p = i+1, p.Next() {
		if p != ts[i] {
			t.Errorf("tree[%d]: p = %p,want = %p", i, p, ts[i])
		}
		if i == len(ts)-1 && p.Next() != nil {
			t.Error("LastChild's next isn't nil")
			break
		}
	}
}

func TestTree(t *testing.T) {
	root := New("root")
	checkTreePointers(t, root, []*Tree{})

	n := root.AddFirst("a")
	checkTreePointers(t, root, []*Tree{n})
	root.MoveToFirst(n)
	checkTreePointers(t, root, []*Tree{n})
	root.MoveToLast(n)
	checkTreePointers(t, root, []*Tree{n})
	root.Remove(n)
	checkTreePointers(t, root, []*Tree{})

	n2 := root.AddFirst(2)
	n1 := root.AddFirst(1)
	n3 := root.AddLast(3)
	n4 := root.AddLast("banana")
	checkTreePointers(t, root, []*Tree{n1, n2, n3, n4})

	root.Remove(n2)
	n2 = &Tree{Value: 2}
	root.InsertFirst(n2)
	checkTreePointers(t, root, []*Tree{n2, n1, n3, n4})

	root.Remove(n2)
	n2 = &Tree{Value: 2}
	root.InsertLast(n2)
	checkTreePointers(t, root, []*Tree{n1, n3, n4, n2})

	root.Remove(n2)
	n2 = &Tree{Value: 2}
	root.Insert(n2, n1)
	checkTreePointers(t, root, []*Tree{n1, n2, n3, n4})

	root.Insert(root, n1) // no change
	checkTreePointers(t, root, []*Tree{n1, n2, n3, n4})

	root.Remove(n2)
	checkTreePointers(t, root, []*Tree{n1, n3, n4})

	root.MoveToFirst(n3) // move from middle
	checkTreePointers(t, root, []*Tree{n3, n1, n4})

	root.MoveToFirst(n1)
	root.MoveToLast(n3) // move from middle
	checkTreePointers(t, root, []*Tree{n1, n4, n3})

	root.MoveToFirst(n3) // move from back
	checkTreePointers(t, root, []*Tree{n3, n1, n4})
	root.MoveToFirst(n3) // should be no-op
	checkTreePointers(t, root, []*Tree{n3, n1, n4})

	root.MoveToLast(n3) // move from front
	checkTreePointers(t, root, []*Tree{n1, n4, n3})
	root.MoveToLast(n3) // should be no-op
	checkTreePointers(t, root, []*Tree{n1, n4, n3})

	root.Remove(n2)
	n2 = root.Add(2, n1) // insert after front
	checkTreePointers(t, root, []*Tree{n1, n2, n4, n3})
	root.Remove(n2)
	n2 = root.Add(2, n4) // insert after middle
	checkTreePointers(t, root, []*Tree{n1, n4, n2, n3})
	root.Remove(n2)
	n2 = root.Add(2, n3) // insert after back
	checkTreePointers(t, root, []*Tree{n1, n4, n3, n2})
	root.Remove(n2)

	// Check standard iteration.
	sum := 0
	for n := root.FirstChild(); n != nil; n = n.Next() {
		if i, ok := n.Value.(int); ok {
			sum += i
		}
	}
	if sum != 4 {
		t.Errorf("sum over l = %d, want 4", sum)
	}

	// Clear all Es by iterating
	var next *Tree
	for n := root.FirstChild(); n != nil; n = next {
		next = n.Next()
		root.Remove(n)
	}
	checkTreePointers(t, root, []*Tree{})
}

func checkTree(t *testing.T, tree *Tree, vs []interface{}) {
	if !checkSubTreesCount(t, tree, len(vs), false) {
		return
	}

	i := 0
	for n := tree.FirstChild(); n != nil; n = n.Next() {
		ne := n.Value.(int)
		if ne != vs[i] {
			t.Errorf("tree[%d].Value = %v, want %v", i, ne, vs[i])
		}
		i++
	}
}

func TestRemove(t *testing.T) {
	root := new(Tree)
	n1 := root.AddLast(1)
	n2 := root.AddLast(2)
	checkTreePointers(t, root, []*Tree{n1, n2})
	n := root.FirstChild()
	root.Remove(n)
	checkTreePointers(t, root, []*Tree{n2})
	root.Remove(n) // no changed
	checkTreePointers(t, root, []*Tree{n2})
	n1 = root.AddFirst(1)
	n3 := root.AddLast(3)
	n4 := root.AddLast(4)
	checkTreePointers(t, root, []*Tree{n1, n2, n3, n4})
}

func TestRemoveNonDescendant(t *testing.T) {
	n1 := new(Tree)
	n1.AddLast(1)
	n1.AddLast(2)

	n2 := new(Tree)
	n2.AddLast(3)
	n2.AddLast(4)

	n := n1.FirstChild()
	n2.Remove(n) // n2 should not change because n is not an descendant of l2
	if n := n2.Count(false); n != 2 {
		t.Errorf("n2.Count(false) = %d, want 2", n)
	}

	n2.InsertFirst(n)
	if n := n2.Count(false); n != 2 {
		t.Errorf("n2.Count(false) = %d, want 2", n)
	}

	n2.InsertLast(n)
	if n := n2.Count(false); n != 2 {
		t.Errorf("n2.Count(false) = %d, want 2", n)
	}

	n2.Insert(n, n2.FirstChild())
	if n := n2.Count(false); n != 2 {
		t.Errorf("n2.Count(false) = %d, want 2", n)
	}
}

func TestDetachedTree(t *testing.T) {
	root := new(Tree)
	root.AddLast(1)
	root.AddLast(2)

	n := root.FirstChild()
	root.Remove(n)
	if n.Value != 1 {
		t.Errorf("n.value = %d, want 1", n.Value)
	}
	if n.Next() != nil {
		t.Errorf("n.Next() != nil")
	}
	if n.Prev() != nil {
		t.Errorf("n.Prev() != nil")
	}
}

func TestMove(t *testing.T) {
	root := new(Tree)
	n1 := root.AddLast(1)
	n2 := root.AddLast(2)
	n3 := root.AddLast(3)
	n4 := root.AddLast(4)

	root.Move(n3, n3)
	checkTreePointers(t, root, []*Tree{n1, n2, n3, n4})

	root.Move(n3, n2)
	checkTreePointers(t, root, []*Tree{n1, n2, n3, n4})

	root.Move(n4, n1)
	checkTreePointers(t, root, []*Tree{n1, n4, n2, n3})
	n2, n3, n4 = n4, n2, n3

	root.Move(n2, n3)
	checkTreePointers(t, root, []*Tree{n1, n3, n2, n4})
	n2, n3 = n3, n2

}

func TestInsertUnknownAt(t *testing.T) {
	var root Tree
	root.AddLast(1)
	root.AddLast(2)
	root.AddLast(3)
	root.Add(1, new(Tree))
	checkTree(t, &root, []interface{}{1, 2, 3})
}

func TestMoveUnknownAt(t *testing.T) {
	var root1 Tree
	n1 := root1.AddLast(1)

	var root2 Tree
	n2 := root2.AddLast(2)

	root1.Move(n1, n2)
	checkTree(t, &root1, []interface{}{1})
	checkTree(t, &root2, []interface{}{2})
}

func checkTreeValues(t *testing.T, wants []interface{}, gets []interface{}) {
	if len(wants) != len(gets) {
		t.Errorf("len(gets) = %v, len(wants) %v", len(gets), len(wants))
		return
	}

	for i := 0; i < len(wants); i++ {
		if wants[i].(int) != gets[i].(int) {
			t.Errorf("gets[%d] = %v, want %v", i, gets[i], wants[i])
		}
	}
}

func TestTravel(t *testing.T) {
	var root Tree
	root.Value = 0
	n1 := root.AddLast(1)
	n2 := root.AddLast(2)
	n3 := root.AddLast(3)
	n4 := root.AddLast(4)
	n1.AddLast(11)
	n1.AddLast(12)
	n1.AddLast(13)
	n1.AddLast(14)
	n2.AddLast(21)
	n22 := n2.AddLast(22)
	if lvl := n22.Level(); lvl != 2 {
		t.Errorf("n22.Level = %d, want 3", lvl)
	}
	if r := n22.Root(); r != &root {
		t.Errorf("n22.Root = %p, want %p", r, &root)
	}
	if p := n22.Parent(); p != n2 {
		t.Errorf("n22.Parent = %p, want %p", p, n2)
	}

	n22.AddLast(221)
	n2.AddLast(23)
	n2.AddLast(24)
	n3.AddLast(31)
	n3.AddLast(32)
	n4.AddLast(41)

	var vals []interface{}
	root.DepthTravel(func(n *Tree) {
		vals = append(vals, n.Value)
	})
	checkTreeValues(t, []interface{}{0, 1, 11, 12, 13, 14, 2, 21, 22, 221, 23, 24, 3, 31, 32, 4, 41}, vals)

	vals = nil
	root.BreadthTravel(func(n *Tree) {
		vals = append(vals, n.Value)
	})
	checkTreeValues(t, []interface{}{0, 1, 2, 3, 4, 11, 12, 13, 14, 21, 22, 23, 24, 31, 32, 41, 221}, vals)

	if count := root.Count(true); count != len(vals)-1 {
		t.Errorf("root.Count(true) = %d, want %d", count, len(vals)-1)
	}

	root.Clear()
	if count := root.Count(true); count != 0 {
		t.Errorf("root.Count(true) = %d, want 0", count)
	}
}

func TestCrossMove(t *testing.T) {
	var root Tree
	root.Value = 0
	n1 := root.AddLast(1)
	n2 := root.AddLast(2)
	n3 := root.AddLast(3)
	n4 := root.AddLast(4)
	n1.AddLast(11)
	n1.AddLast(12)
	n1.AddLast(13)
	n1.AddLast(14)
	n2.AddLast(21)
	n22 := n2.AddLast(22)
	n2.AddLast(23)
	n2.AddLast(24)
	n3.AddLast(31)
	n3.AddLast(32)
	n4.AddLast(41)
	n42 := n4.AddLast(42)
	root.Remove(n42)
	var vals []interface{}

	vals = nil
	root.DepthTravel(func(n *Tree) {
		vals = append(vals, n.Value)
	})

	checkTreeValues(t, []interface{}{0, 1, 11, 12, 13, 14, 2, 21, 22, 23, 24, 3, 31, 32, 4, 41}, vals)
	n221 := n22.AddLast(221)
	n2.MoveToFirst(n221)
	vals = nil
	root.DepthTravel(func(n *Tree) {
		vals = append(vals, n.Value)
	})
	checkTreeValues(t, []interface{}{0, 1, 11, 12, 13, 14, 2, 221, 21, 22, 23, 24, 3, 31, 32, 4, 41}, vals)

	root.Remove(n221)
	n221 = n22.AddLast(221)
	n2.MoveToLast(n221)
	vals = nil
	root.DepthTravel(func(n *Tree) {
		vals = append(vals, n.Value)
	})
	checkTreeValues(t, []interface{}{0, 1, 11, 12, 13, 14, 2, 21, 22, 23, 24, 221, 3, 31, 32, 4, 41}, vals)

	root.Remove(n221)
	n221 = n22.AddLast(221)
	root.Move(n221, n2)
	vals = nil
	root.DepthTravel(func(n *Tree) {
		vals = append(vals, n.Value)
	})
	checkTreeValues(t, []interface{}{0, 1, 11, 12, 13, 14, 2, 21, 22, 23, 24, 221, 3, 31, 32, 4, 41}, vals)

	root.Remove(n221)
	n221 = n22.AddLast(221)
	n2.Move(n221, n22)
	vals = nil
	root.DepthTravel(func(n *Tree) {
		vals = append(vals, n.Value)
	})
	checkTreeValues(t, []interface{}{0, 1, 11, 12, 13, 14, 2, 21, 22, 221, 23, 24, 3, 31, 32, 4, 41}, vals)
	if n221.Parent() != n2 {
		t.Error("parent error")
	}
	if n221.Root() != &root {
		t.Error("root error")
	}
	if n221.IsRoot() {
		t.Error("isroot error")
	}
	if !n221.IsChildOf(n2) {
		t.Error("IsChildOf error")
	}
	if !n2.IsParentOf(n221) {
		t.Error("IsParentOf error")
	}
	if !n221.IsDescendantOf(&root) {
		t.Error("IsDescendantOf error")
	}
	if !n221.IsBrotherOf(n22) {
		t.Error("IsBrotherOf error")
	}
}
