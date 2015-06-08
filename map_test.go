package main

import (
	"fmt"
	"strconv"
	"testing"
)

func BenchmarkInsertIntoMap_intkey_newkey(b *testing.B) {
	m := map[int]int{}
	for i := 0; i < b.N; i++ {
		m[i] = i
	}
	// fmt.Printf("Length of map: %d\n", len(m))
}

func BenchmarkInsertIntoMap_intkey_singlekey(b *testing.B) {
	m := map[int]int{}
	for i := 0; i < b.N; i++ {
		m[0] = i
	}
	// fmt.Printf("Length of map: %d\n", len(m))
}

func BenchmarkInsertIntoMap_stringkey_singlekey(b *testing.B) {
	m := map[string]int{}
	for i := 0; i < b.N; i++ {
		m["key"] = i
	}
	// fmt.Printf("Length of map: %d\n", len(m))
}

func BenchmarkMapLookup_1000_first(b *testing.B) {
	m := map[string]int{}
	for i := 0; i < 1000; i++ {
		m[fmt.Sprintf("key_%7d", i)] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m["key_000000"]
	}
}

func BenchmarkMapLookup_1000_last(b *testing.B) {
	m := map[string]int{}
	for i := 0; i < 1000; i++ {
		m[fmt.Sprintf("key_%7d", i)] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m["key_000999"]
	}
}

func BenchmarkMapIteration(b *testing.B) {
	m := map[string]int{}

	for i := 0; i < 1000; i++ {
		m[strconv.FormatInt(int64(i), 10)] = i
	}

	b.ResetTimer()
	var sum int
	for i := 0; i < b.N; i++ {
		for _, v := range m {
			sum += v
		}
		_ = sum
	}
}
