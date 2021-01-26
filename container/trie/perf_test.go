// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package trie_test

import (
	"bufio"
	"os"
	"testing"

	"github.com/cnotch/algo/container/trie/compress"
	"github.com/cnotch/algo/container/trie/critbit"
	"github.com/cnotch/algo/container/trie/darts"
	"github.com/cnotch/algo/container/trie/simple"
	"github.com/cnotch/algo/container/trie/ternary"
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

func BenchmarkMapStore(b *testing.B) {
	words := loadTestFile("test/assets/words.txt")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		m := make(map[string]Value)
		for _, w := range words {
			m[w] = w
		}
	}
}
func BenchmarkMapLoad(b *testing.B) {
	words := loadTestFile("test/assets/words.txt")
	m := make(map[string]Value)
	for _, w := range words {
		m[w] = w
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for _, w := range words {
			_, _ = m[w]
		}
	}
}

func BenchmarkSimpleStore(b *testing.B) {
	words := loadTestFile("test/assets/words.txt")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		tree := simple.New()
		for _, w := range words {
			tree.Store(w, w)
		}
	}
}

func BenchmarkSimpleLoad(b *testing.B) {
	words := loadTestFile("test/assets/words.txt")
	tree := simple.New()
	for _, w := range words {
		tree.Store(w, w)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for _, w := range words {
			_, _ = tree.Load(w)
		}
	}
}

func BenchmarkCompressStore(b *testing.B) {
	words := loadTestFile("test/assets/words.txt")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		tree := compress.New()
		for _, w := range words {
			tree.Store(w, w)
		}
	}
}

func BenchmarkCompressLoad(b *testing.B) {
	words := loadTestFile("test/assets/words.txt")
	tree := compress.New()
	for _, w := range words {
		tree.Store(w, w)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for _, w := range words {
			_, _ = tree.Load(w)
		}
	}
}

func BenchmarkCritbitStore(b *testing.B) {
	words := loadTestFile("test/assets/words.txt")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		tree := critbit.New()
		for _, w := range words {
			tree.Store(w, w)
		}
	}
}
func BenchmarkCritbitLoad(b *testing.B) {
	words := loadTestFile("test/assets/words.txt")
	tree := critbit.New()
	for _, w := range words {
		tree.Store(w, w)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for _, w := range words {
			_, _ = tree.Load(w)
		}
	}
}

func BenchmarkTernaryStore(b *testing.B) {
	words := loadTestFile("test/assets/words.txt")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		tree := ternary.New()
		for _, w := range words {
			tree.Store(w, w)
		}
	}
}
func BenchmarkTernaryLoad(b *testing.B) {
	words := loadTestFile("test/assets/words.txt")
	tree := ternary.New()
	for _, w := range words {
		tree.Store(w, w)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for _, w := range words {
			_, _ = tree.Load(w)
		}
	}
}

func BenchmarkDartsStore(b *testing.B) {
	words := loadTestFile("test/assets/words.txt")
	option := darts.Entries(words, func(i int) Value {
		return words[i]
	})

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		tree := darts.New(option)
		_ = tree
	}
}

func BenchmarkDartsLoad(b *testing.B) {
	words := loadTestFile("test/assets/words.txt")
	option := darts.Entries(words, func(i int) Value {
		return words[i]
	})
	tree := darts.New(option)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for _, k := range words {
			_, _ = tree.Load(k)
		}
	}
}

func BenchmarkHSKMapStore(b *testing.B) {
	words := loadTestFile("test/assets/hsk_words.txt")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		m := make(map[string]Value)
		for _, w := range words {
			m[w] = w
		}
	}
}
func BenchmarkHSKMapLoad(b *testing.B) {
	words := loadTestFile("test/assets/hsk_words.txt")
	m := make(map[string]Value)
	for _, w := range words {
		m[w] = w
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for _, w := range words {
			_, _ = m[w]
		}
	}
}

func BenchmarkHSKSimpleStore(b *testing.B) {
	words := loadTestFile("test/assets/hsk_words.txt")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		tree := simple.New()
		for _, w := range words {
			tree.Store(w, w)
		}
	}
}

func BenchmarkHSKSimpleLoad(b *testing.B) {
	words := loadTestFile("test/assets/hsk_words.txt")
	tree := simple.New()
	for _, w := range words {
		tree.Store(w, w)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for _, w := range words {
			_, _ = tree.Load(w)
		}
	}
}

func BenchmarkHSKCompressStore(b *testing.B) {
	words := loadTestFile("test/assets/hsk_words.txt")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		tree := compress.New()
		for _, w := range words {
			tree.Store(w, w)
		}
	}
}

func BenchmarkHSKCompressLoad(b *testing.B) {
	words := loadTestFile("test/assets/hsk_words.txt")
	tree := compress.New()
	for _, w := range words {
		tree.Store(w, w)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for _, w := range words {
			_, _ = tree.Load(w)
		}
	}
}

func BenchmarkHSKCritbitStore(b *testing.B) {
	words := loadTestFile("test/assets/hsk_words.txt")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		tree := critbit.New()
		for _, w := range words {
			tree.Store(w, w)
		}
	}
}

func BenchmarkHSKCritbitLoad(b *testing.B) {
	words := loadTestFile("test/assets/hsk_words.txt")
	tree := critbit.New()
	for _, w := range words {
		tree.Store(w, w)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for _, w := range words {
			_, _ = tree.Load(w)
		}
	}
}

func BenchmarkHSKTernaryStore(b *testing.B) {
	words := loadTestFile("test/assets/hsk_words.txt")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		tree := ternary.New()
		for _, w := range words {
			tree.Store(w, w)
		}
	}
}

func BenchmarkHSKTernaryLoad(b *testing.B) {
	words := loadTestFile("test/assets/hsk_words.txt")
	tree := ternary.New()
	for _, w := range words {
		tree.Store(w, w)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for _, w := range words {
			_, _ = tree.Load(w)
		}
	}
}

func BenchmarkHSKDartsStore(b *testing.B) {
	words := loadTestFile("test/assets/hsk_words.txt")
	option := darts.Entries(words, func(i int) Value {
		return words[i]
	})

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		tree := darts.New(option)
		_ = tree
	}
}

func BenchmarkHSKDartsLoad(b *testing.B) {
	words := loadTestFile("test/assets/hsk_words.txt")
	option := darts.Entries(words, func(i int) Value {
		return words[i]
	})
	tree := darts.New(option)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for _, k := range words {
			_, _ = tree.Load(k)
		}
	}
}
