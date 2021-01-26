// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package darts

// Option DARTS 前缀树选项接口
type Option interface {
	apply(*Trie)
}

type optionFunc func(*Trie)

func (f optionFunc) apply(t *Trie) {
	f(t)
}

// Entry 向 Darts 前缀树添加实体条目的选项
func Entry(key string, value Value) Option {
	return optionFunc(func(t *Trie) {
		t.es = append(t.es, entry{key, value})
	})
}

// Entries 向 Darts 前缀树添加多个实体条目的选项
func Entries(keys []string, vf func(i int) Value) Option {
	return optionFunc(func(t *Trie) {
		if cap(t.es)-len(t.es) < len(keys) {
			es := make([]entry, len(t.es), cap(t.es)+len(keys))
			copy(es, t.es)
			t.es = es
		}

		if vf == nil {
			var v Value
			for _, k := range keys {
				t.es = append(t.es, entry{k, v})
			}
		} else {
			for i, k := range keys {
				t.es = append(t.es, entry{k, vf(i)})
			}
		}
	})
}
