package main

import (
	"fmt"
	"testing"
)

type StrIntItem struct {
	key string
	val int
}

func BenchmarkSliceLookup_100_last(b *testing.B) {
	arr := []*StrIntItem{}
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("key_%7d", i)
		arr = append(arr, &StrIntItem{key, i})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		arrLen := len(arr)
		var item *StrIntItem
		for arrIndex := 0; arrIndex < arrLen; arrIndex++ {
			if arr[arrIndex].key == "key_999" {
				item = arr[arrIndex]
			}
		}
		_ = item
	}
}

func BenchmarkSliceLookup_1000_last(b *testing.B) {
	arr := []*StrIntItem{}
	for i := 0; i < 1000; i++ {
		key := fmt.Sprintf("key_%7d", i)
		arr = append(arr, &StrIntItem{key, i})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		arrLen := len(arr)
		var item *StrIntItem
		for arrIndex := 0; arrIndex < arrLen; arrIndex++ {
			if arr[arrIndex].key == "key_999" {
				item = arr[arrIndex]
			}
		}
		_ = item
	}
}

type Struct1 struct {
}

func BenchmarkConstructEmptyStruct(b *testing.B) {
	var p interface{}
	for i := 0; i < b.N; i++ {
		p = &Struct1{}
		_ = p
	}
}

func BenchmarkConstructSliceWithNew(b *testing.B) {
	var arr *[]string
	for i := 0; i < b.N; i++ {
		arr = new([]string)
		_ = arr
	}
}

func BenchmarkConstructSliceWithMake(b *testing.B) {
	var arr []string
	for i := 0; i < b.N; i++ {
		arr = make([]string, 1)
		_ = arr
	}
}

func BenchmarkConstructSliceWithInitializer(b *testing.B) {
	var arr []string
	s := ""
	for i := 0; i < b.N; i++ {
		arr = []string{s}
		_ = arr
	}
}

func BenchmarkConstructSlice10000WithSize(b *testing.B) {
	const length = 10000
	arr := make([]string, length)
	s := ""
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < length; j++ {
			arr = append(arr, s)
			_ = s
		}
	}
}

func BenchmarkConstructSlice10000WithoutSize(b *testing.B) {
	const length = 10000
	arr := make([]string, 0)
	s := ""
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < length; j++ {
			arr = append(arr, s)
			_ = s
		}
	}
}

type Struct2 struct {
	name  string
	code1 int
	code2 int
}

func BenchmarkAlloc64KSlice(b *testing.B) {
	var arr []byte
	//b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		arr = make([]byte, 65536)
		_ = arr
	}
}

func BenchmarkReadByteFrom64KSlice(b *testing.B) {
	const length = 65536
	arr := make([]byte, length)
	for i := 0; i < b.N; i++ {
		value := arr[i%length]
		_ = value
	}
}

func BenchmarkWriteByteTo64KSlice(b *testing.B) {
	const length = 65536
	arr := make([]byte, length)
	for i := 0; i < b.N; i++ {
		arr[i%length] = byte(i % 256)
	}
}
