// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package trie_test

import (
	"sort"
	"testing"

	"github.com/cnotch/algo/container/trie"
	"github.com/cnotch/algo/container/trie/compress"
	"github.com/cnotch/algo/container/trie/critbit"
	"github.com/cnotch/algo/container/trie/simple"
	"github.com/cnotch/algo/container/trie/ternary"
	"github.com/stretchr/testify/assert"
)

type Value = trie.Value

var tries = []struct {
	name    string
	newFunc func() trie.Interface
}{
	{"simplicity", func() trie.Interface { return simple.New() }},
	{"compress", func() trie.Interface { return compress.New() }},
	{"critbit", func() trie.Interface { return critbit.New() }},
	{"ternary", func() trie.Interface { return ternary.New() }},
}

func TestTrieStore(t *testing.T) {
	for _, tt := range tries {
		t.Run(tt.name+"/Store", func(t *testing.T) {
			tree := tt.newFunc()
			key := "key"

			ov, updated := tree.Store(key, "value")
			assert.Nil(t, ov)
			assert.False(t, updated)
			assert.Equal(t, 1, tree.Size())

			v, found := tree.Load(key)
			assert.True(t, found)
			assert.Equal(t, "value", v)

			ov, updated = tree.Store(key, "otherValue")
			assert.Equal(t, "value", ov)
			assert.True(t, updated)
			assert.Equal(t, 1, tree.Size())

			v, found = tree.Load(key)
			assert.True(t, found)
			assert.Equal(t, "otherValue", v)

			// nil or empty key
			tree.Clear()

			v, found = tree.Load("")
			assert.False(t, found)
			assert.Nil(t, v)

			ov, updated = tree.Store("", "value")
			assert.Nil(t, ov)
			assert.False(t, updated)
			assert.Equal(t, 1, tree.Size())

			v, found = tree.Load("")
			assert.True(t, found)
			assert.Equal(t, "value", v)

			// same prefix
			words := []string{"elector", "electibles", "elect", "electible"}

			for _, k := range words {
				tree.Store(k, k)
			}

			for _, k := range words {
				v, found = tree.Load(k)
				assert.True(t, found)
				assert.Equal(t, k, v)
			}
		})
	}
}

func TestTrieLoad(t *testing.T) {
	for _, tt := range tries {
		t.Run(tt.name+"/Load", func(t *testing.T) {
			tree := tt.newFunc()
			searchTerms := []string{"河北", "A", "a", "aa", "河南", "天津"}

			for _, term := range searchTerms {
				tree.Store(term, term)
			}

			for _, term := range searchTerms {
				v, found := tree.Load(term)
				assert.True(t, found)
				assert.Equal(t, term, v)
			}

			// shorter
			v, found := tree.Load("河")
			assert.Nil(t, v)
			assert.False(t, found)

			// longer
			v, found = tree.Load("Aaa")
			assert.Nil(t, v)
			assert.False(t, found)
		})
	}
}

func TestTrieLoadOrStore(t *testing.T) {
	for _, tt := range tries {
		t.Run(tt.name+"/LoadOrStore", func(t *testing.T) {
			tree := tt.newFunc()
			key := "key"

			av, loaded := tree.LoadOrStore(key, "value")
			assert.False(t, loaded)
			assert.Equal(t, "value", av)
			assert.Equal(t, 1, tree.Size())

			av, loaded = tree.LoadOrStore(key, "othervalue")
			assert.True(t, loaded)
			assert.Equal(t, "value", av)
			assert.Equal(t, 1, tree.Size())
		})
	}
}

func TestTrieDelete(t *testing.T) {
	var data = []*struct {
		message      string
		insert       []string
		delete       []string
		size         int
		deleteStatus bool
	}{
		{
			"Insert 1 Delete 0",
			[]string{"test"},
			[]string{""},
			1,
			false},
		{
			"Insert nil Delete nil",
			[]string{""},
			[]string{""},
			0,
			true},
		{
			"Insert 1 Delete 1",
			[]string{"test"},
			[]string{"test"},
			0,
			true},
		{
			"Insert 2 Delete 1",
			[]string{"test2", "test1"},
			[]string{"test2"},
			1,
			true},
		{
			"Insert 2 Delete 2",
			[]string{"test1", "test2"},
			[]string{"test1", "test2"},
			0,
			true},
		{
			"Insert 3 Try to delete 1 shorter prefix",
			[]string{"keyb::", "keyb::2", "keyb::3"},
			[]string{"keyb:", "kb"},
			3,
			false},
		{
			"Insert 5 Delete 1",
			[]string{"3", "1", "2", "4", "5"},
			[]string{"1"},
			4,
			true},
		{
			"Insert 5 Try to delete 1 wrong",
			[]string{"6", "2", "3", "4", "5"},
			[]string{"123"},
			5,
			false},
		{
			"Insert 5 Delete 5",
			[]string{"1", "2", "3", "4", "5"},
			[]string{"1", "2", "3", "4", "5"},
			0,
			true},
		{
			"Insert 17 Delete 1",
			[]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17"},
			[]string{"2"},
			16,
			true},
		{
			"Insert 17 Delete 17",
			[]string{"5", "6", "7", "8", "1", "2", "3", "4", "9", "10", "11", "12", "13", "14", "15", "16", "17"},
			[]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17"},
			0,
			true},
	}

	for _, tt := range tries {
		t.Run(tt.name+"/Delete", func(t *testing.T) {
			tree := tt.newFunc()
			for _, ds := range data {
				tree.Clear()
				// build
				for _, term := range ds.insert {
					tree.Store(term, term)
				}

				// process
				for _, term := range ds.delete {
					val, deleted := tree.Delete(term)
					assert.Equal(t, ds.deleteStatus, deleted, ds.message)

					if s, ok := val.(string); ok {
						assert.Equal(t, term, s, ds.message)
					}

					_, found := tree.Load(term)
					assert.False(t, found, ds.message)
				}

				assert.Equal(t, ds.size, tree.Size(), ds.message)
			}
		})
	}
}

func TestTrieDeleteForPrefixIsKey(t *testing.T) {
	for _, tt := range tries {
		t.Run(tt.name+"/DeleteForPrefixIsKey", func(t *testing.T) {
			tree := tt.newFunc()

			keys := []string{"a", "ab", "abc", "abcd", "abcde"}
			for _, key := range keys {
				tree.Store(key, key)
			}
			for i := 0; i < len(keys); i++ {
				tree.Delete(keys[i])
				for j := i + 1; j < len(keys); j++ {
					v, loaded := tree.Load(keys[j])
					assert.True(t, loaded)
					assert.Equal(t, keys[j], v)
				}
			}
			assert.Equal(t, 0, tree.Size())

			// 逆序删除
			for _, key := range keys {
				tree.Store(key, key)
			}
			for i := len(keys) - 1; i >= 0; i-- {
				tree.Delete(keys[i])
				for j := 0; j < i; j++ {
					v, loaded := tree.Load(keys[j])
					assert.True(t, loaded)
					assert.Equal(t, keys[j], v)
				}
			}
			assert.Equal(t, 0, tree.Size())
		})
	}
}

func TestTrieRange(t *testing.T) {
	for _, tt := range tries {
		t.Run(tt.name+"/Range", func(t *testing.T) {
			tree := tt.newFunc()

			getks := []string{}
			getvs := []string{}
			tree.Range(func(key string, value Value) bool {
				getks = append(getks, key)
				getvs = append(getvs, value.(string))
				return true
			})
			assert.Equal(t, []string{}, getks)
			assert.Equal(t, []string{}, getvs)

			keys := []string{"a", "ab", "ac", "ad", "ae", "abc", "abd", "abe"}
			sort.Strings(keys)
			for _, key := range keys {
				tree.Store(key, key)
			}

			tree.Range(func(key string, value Value) bool {
				getks = append(getks, key)
				getvs = append(getvs, value.(string))
				return true
			})
			assert.Equal(t, keys, getks)
			assert.Equal(t, keys, getvs)
		})
	}
}

func TestTrieRangeCancel(t *testing.T) {
	for _, tt := range tries {
		t.Run(tt.name+"/RangeCancel", func(t *testing.T) {
			tree := tt.newFunc()

			keys := []string{"ac", "a", "ab", "ad", "ae", "abc", "abd", "abe"}
			for _, key := range keys {
				tree.Store(key, key)
			}

			getks := []string{}
			getn := 0
			tree.Range(func(key string, value Value) bool {
				getn++
				getks = append(getks, key)
				return getn < 5
			})
			assert.Equal(t, 5, len(getks))
		})
	}
}

func TestTrieRangePrefix(t *testing.T) {
	dataSet := []struct {
		keyPrefix string
		keys      []string
		expected  []string
	}{
		{
			"",
			[]string{},
			[]string{},
		},
		{
			"api",
			[]string{"api.foo.bar", "api.foo.baz", "api.foe.fum", "abc.123.456", "api.foo", "api"},
			[]string{"api.foo.bar", "api.foo.baz", "api.foe.fum", "api.foo", "api"},
		}, {
			"a",
			[]string{"api.foo.bar", "api.foo.baz", "api.foe.fum", "abc.123.456", "api.foo", "api"},
			[]string{"api.foo.bar", "api.foo.baz", "api.foe.fum", "abc.123.456", "api.foo", "api"},
		}, {
			"b",
			[]string{"api.foo.bar", "api.foo.baz", "api.foe.fum", "abc.123.456", "api.foo", "api"},
			[]string{},
		},
		{
			"api.",
			[]string{"api.foo.bar", "api.foo.baz", "api.foe.fum", "abc.123.456", "api.foo", "api"},
			[]string{"api.foo.bar", "api.foo.baz", "api.foe.fum", "api.foo"},
		},
		{
			"api.foo.bar",
			[]string{"api.foo.bar", "api.foo.baz", "api.foe.fum", "abc.123.456", "api.foo", "api"},
			[]string{"api.foo.bar"},
		},
		{
			"api.end",
			[]string{"api.foo.bar", "api.foo.baz", "api.foe.fum", "abc.123.456", "api.foo", "api"},
			[]string{},
		}, {
			"",
			[]string{"api.foo.bar", "api.foo.baz", "api.foe.fum", "abc.123.456", "api.foo", "api"},
			[]string{"api.foo.bar", "api.foo.baz", "api.foe.fum", "abc.123.456", "api.foo", "api"},
		}, {
			"this:key:has",
			[]string{
				"this:key:has:a:long:prefix:3",
				"this:key:has:a:long:common:prefix:2",
				"this:key:has:a:long:common:prefix:1",
			},
			[]string{
				"this:key:has:a:long:prefix:3",
				"this:key:has:a:long:common:prefix:2",
				"this:key:has:a:long:common:prefix:1",
			},
		}, {
			"ele",
			[]string{"elector", "electibles", "elect", "electible"},
			[]string{"elector", "electibles", "elect", "electible"},
		},
		{
			"long.api.url.v1",
			[]string{"long.api.url.v1.foo", "long.api.url.v1.bar", "long.api.url.v2.foo"},
			[]string{"long.api.url.v1.foo", "long.api.url.v1.bar"},
		},
	}

	for _, tt := range tries {
		t.Run(tt.name+"/RangePrefix", func(t *testing.T) {
			for _, d := range dataSet {
				tree := tt.newFunc()
				for _, k := range d.keys {
					tree.Store(k, k)
				}

				actual := []string{}
				tree.RangePrefix(d.keyPrefix, func(key string, value Value) bool {
					actual = append(actual, key)
					return true
				})

				sort.Strings(d.expected)
				sort.Strings(actual)
				assert.Equal(t, d.expected, actual, d.keyPrefix)
			}
		})
	}
}

func TestTrieMinMax(t *testing.T) {
	for _, tt := range tries {
		words := []string{"3", "4", "1", "2", "5"}
		t.Run(tt.name+"/MinMax", func(t *testing.T) {
			tree := tt.newFunc()

			min, found := tree.Minimum()
			assert.False(t, found)
			assert.Nil(t, min)
			max, found := tree.Maximum()
			assert.False(t, found)
			assert.Nil(t, max)

			for _, k := range words {
				tree.Store(k, k)
			}
			min, found = tree.Minimum()
			assert.True(t, found)
			assert.Equal(t, "1", min)
			max, found = tree.Maximum()
			assert.True(t, found)
			assert.Equal(t, "5", max)

			tree.Store("", nil)
			min, found = tree.Minimum()
			assert.True(t, found)
			assert.Nil(t, min)
		})
	}
}

func TestTriePrefix(t *testing.T) {
	for _, tt := range tries {
		words := []string{"elector", "electibles", "elect", "electible"}
		t.Run(tt.name+"/Prefix", func(t *testing.T) {
			tree := tt.newFunc()
			for _, k := range words {
				tree.Store(k, k)
			}
			ok := tree.HasPrefix("ele")
			assert.True(t, ok)
			ok = tree.HasPrefix("elet")
			assert.False(t, ok)
		})
	}
}

func TestTrieShortestPrefixOf(t *testing.T) {
	for _, tt := range tries {
		words := []string{
			"a", "aa", "b", "bb", "ab", "ba", "aba", "bab", "foo",
			"foobar",
			"foobarbaz",
			"foobarbazzip",
			"foozip",
		}

		expects := map[string]string{
			"a":            "a",
			"a^":           "a",
			"aaa":          "a",
			"abc":          "a",
			"bac":          "b",
			"bbb":          "b",
			"bc":           "b",
			"foo":          "foo",
			"foob":         "foo",
			"foobar":       "foo",
			"foobarba":     "foo",
			"foobarbaz":    "foo",
			"foobarbazzi":  "foo",
			"foobarbazzip": "foo",
			"foozi":        "foo",
			"foozip":       "foo",
			"foozipzap":    "foo",
		}

		t.Run(tt.name+"/ShortestPrefixOf", func(t *testing.T) {
			tree := tt.newFunc()
			for _, k := range words {
				tree.Store(k, k)
			}

			for g, k := range expects {
				if key, value, ok := tree.ShortestPrefixOf(g); !ok || key != k || value != k {
					t.Errorf("ShortestPrefixOf() - invalid result - %s not %s", key, g)
				}
			}

			if _, _, ok := tree.ShortestPrefixOf(""); ok {
				t.Error("ShortestPrefixOf() - invalid result - not empty")
			}
			if _, _, ok := tree.ShortestPrefixOf("^"); ok {
				t.Error("ShortestPrefixOf() - invalid result - not empty")
			}
			if _, _, ok := tree.ShortestPrefixOf("c"); ok {
				t.Error("ShortestPrefixOf() - invalid result - not empty")
			}
		})
	}
}

func TestTrieLongestPrefixOf(t *testing.T) {
	for _, tt := range tries {
		words := []string{
			"a", "aa", "b", "bb", "ab", "ba", "aba", "bab", "foo",
			"foobar",
			"foobarbaz",
			"foobarbazzip",
			"foozip",
		}

		expects := map[string]string{
			"a":            "a",
			"a^":           "a",
			"aaa":          "aa",
			"abc":          "ab",
			"bac":          "ba",
			"bbb":          "bb",
			"bc":           "b",
			"foo":          "foo",
			"foob":         "foo",
			"foobar":       "foobar",
			"foobarba":     "foobar",
			"foobarbaz":    "foobarbaz",
			"foobarbazzi":  "foobarbaz",
			"foobarbazzip": "foobarbazzip",
			"foozi":        "foo",
			"foozip":       "foozip",
			"foozipzap":    "foozip",
		}

		t.Run(tt.name+"/LongestPrefixOf", func(t *testing.T) {
			tree := tt.newFunc()
			for _, k := range words {
				tree.Store(k, k)
			}

			for g, k := range expects {
				if key, value, ok := tree.LongestPrefixOf(g); !ok || key != k || value != k {
					t.Errorf("LongestPrefixOf() - invalid result - %s not %s", key, g)
				}
			}

			if _, _, ok := tree.LongestPrefixOf(""); ok {
				t.Error("LongestPrefixOf() - invalid result - not empty")
			}
			if _, _, ok := tree.LongestPrefixOf("^"); ok {
				t.Error("LongestPrefixOf() - invalid result - not empty")
			}
			if _, _, ok := tree.LongestPrefixOf("c"); ok {
				t.Error("LongestPrefixOf() - invalid result - not empty")
			}
		})
	}
}
