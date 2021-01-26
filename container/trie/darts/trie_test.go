// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package darts_test

import (
	"bufio"
	"os"
	"sort"
	"testing"

	"github.com/cnotch/algo/container/trie"
	"github.com/cnotch/algo/container/trie/darts"
	"github.com/stretchr/testify/assert"
)

func loadTestFile(path string) []string {
	// path = "test/assets/hsk_words.txt"
	file, err := os.Open(path)
	if err != nil {
		panic("Couldn't open " + path)
	}
	defer file.Close()

	var words []string
	reader := bufio.NewReader(file)
	for {
		if line, err := reader.ReadString(byte('\n')); err != nil {
			break
		} else {
			if len(line) > 0 {
				words = append(words, line[:len(line)-1])
			}
		}
	}
	return words
}

type memStats interface {
	MemStats() (mallocs, frees int)
}

func TestDartsMemStats(t *testing.T) {
	var words = loadTestFile("../test/assets/words.txt")
	tree := darts.New(darts.Entries(words, func(i int) darts.Value {
		return words[i]
	}))

	t.Run("MemStats", func(t *testing.T) {
		// ms := tree.(memStats)
		m, f := tree.MemStats()
		_, _ = m, f

		for i := 0; i < len(words); i++ {
			val, ok := tree.Load(words[i])
			assert.True(t, ok)
			assert.Equal(t, words[i], val)
		}
	})
}

func TestDartsLoad(t *testing.T) {
	t.Run("Load", func(t *testing.T) {
		var words = []string{"河北", "A", "a", "aa", "河南", "天津"}
		tree := darts.New(darts.Entries(words, func(i int) darts.Value {
			return words[i]
		}))

		for _, term := range words {
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

func TestDartsRange(t *testing.T) {
	t.Run("Range", func(t *testing.T) {
		var words = []string{}
		tree := darts.New(darts.Entries(words, func(i int) darts.Value {
			return words[i]
		}))

		getks := []string{}
		getvs := []string{}
		tree.Range(func(key string, value trie.Value) bool {
			getks = append(getks, key)
			getvs = append(getvs, value.(string))
			return true
		})
		assert.Equal(t, []string{}, getks)
		assert.Equal(t, []string{}, getvs)

		keys := []string{"ae", "abc", "abd", "a", "ab", "ac", "ad", "abe"}
		tree = darts.New(darts.Entries(keys, func(i int) darts.Value {
			return keys[i]
		}))
		sort.Strings(keys)

		tree.Range(func(key string, value trie.Value) bool {
			getks = append(getks, key)
			getvs = append(getvs, value.(string))
			return true
		})
		assert.Equal(t, keys, getks)
		assert.Equal(t, keys, getvs)
	})
}

func TestDartsRangeCancel(t *testing.T) {
	t.Run("RangeCancel", func(t *testing.T) {
		keys := []string{"ac", "a", "ab", "ad", "ae", "abc", "abd", "abe"}
		tree := darts.New(darts.Entries(keys, func(i int) darts.Value {
			return keys[i]
		}))

		getks := []string{}
		getn := 0
		tree.Range(func(key string, value trie.Value) bool {
			getn++
			getks = append(getks, key)
			return getn < 5
		})
		assert.Equal(t, 5, len(getks))
	})
}

func TestDartsRangePrefix(t *testing.T) {
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

	for _, d := range dataSet {
		t.Run("RangePrefix/"+d.keyPrefix, func(t *testing.T) {
			tree := darts.New(darts.Entries(d.keys, func(i int) darts.Value {
				return d.keys[i]
			}))

			actual := []string{}
			tree.RangePrefix(d.keyPrefix, func(key string, value trie.Value) bool {
				actual = append(actual, key)
				return true
			})

			sort.Strings(d.expected)
			sort.Strings(actual)
			assert.Equal(t, d.expected, actual, d.keyPrefix)
		})
	}
}

func TestDartsMinMax(t *testing.T) {
	t.Run("MinMax", func(t *testing.T) {
		keys := []string{}
		tree := darts.New(darts.Entries(keys, func(i int) darts.Value {
			return keys[i]
		}))

		min, found := tree.Minimum()
		assert.False(t, found)
		assert.Nil(t, min)
		max, found := tree.Maximum()
		assert.False(t, found)
		assert.Nil(t, max)

		keys = []string{"3", "4", "1", "2", "5"}
		tree = darts.New(darts.Entries(keys, func(i int) darts.Value {
			return keys[i]
		}))

		min, found = tree.Minimum()
		assert.True(t, found)
		assert.Equal(t, "1", min)
		max, found = tree.Maximum()
		assert.True(t, found)
		assert.Equal(t, "5", max)
	})
}

func TestDartsPrefix(t *testing.T) {
	t.Run("Prefix", func(t *testing.T) {
		keys := []string{"elector", "electibles", "elect", "electible"}
		tree := darts.New(darts.Entries(keys, func(i int) darts.Value {
			return keys[i]
		}))

		ok := tree.HasPrefix("ele")
		assert.True(t, ok)
		ok = tree.HasPrefix("elet")
		assert.False(t, ok)
	})
}

func TestTrieShortestPrefixOf(t *testing.T) {
	words := []string{
		"a",
		"aa",
		"b",
		"bb",
		"ab",
		"ba",
		"aba",
		"bab",
		"foo",
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

	t.Run("ShortestPrefixOf", func(t *testing.T) {
		tree := darts.New(darts.Entries(words, func(i int) darts.Value {
			return words[i]
		}))

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

func TestTrieLongestPrefixOf(t *testing.T) {
	words := []string{
		"a",
		"aa",
		"b",
		"bb",
		"ab",
		"ba",
		"aba",
		"bab",
		"foo",
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

	t.Run("LongestPrefixOf", func(t *testing.T) {
		tree := darts.New(darts.Entries(words, func(i int) darts.Value {
			return words[i]
		}))

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

func BenchmarkDartsNew(b *testing.B) {
	words := loadTestFile("../test/assets/words.txt")

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		tree := darts.New(darts.Entries(words, nil))
		_ = tree
	}
}
