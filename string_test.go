package main

import (
	"bytes"
	"fmt"
	"testing"

	"strconv"
	"strings"
)

func BenchmarkSprintfIntString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%d", i)
	}
}

func BenchmarkConcatString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := "a"
		s += "b"
	}
}

func BenchmarkConcat10Strings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := "a"
		s += "b1" + "b1" + "b1" + "b1" + "b1" + "b1" + "b1" + "b1" + "b1" + "b1"
	}
}

func BenchmarkConcatToLongString(b *testing.B) {
	s := strings.Repeat(" ", 10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s2 := s + "x"
		_ = s2
	}
}

func BenchmarkSprintfLongStr(b *testing.B) {
	s := strings.Repeat(" ", 10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s2 := fmt.Sprintf("%s%s", s, "x")
		_ = s2
	}
}

func BenchmarkSprintfStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := "a"
		s2 := fmt.Sprintf("%s%s", s, "b")
		_ = s2
	}
}

func BenchmarkSprintfInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := "a"
		s2 := fmt.Sprintf("%s%d", s, 10)
		_ = s2
	}
}

func BenchmarkSprintf10Ints(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := "a"
		s2 := fmt.Sprintf(
			"%s%d %d %d %d %d %d %d %d %d %d",
			s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		_ = s2
	}
}

func BenchmarkFormatInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s2 := " " + strconv.FormatInt(100, 10)
		_ = s2
	}
}

func BenchmarkBytesBuf(b *testing.B) {
	s := strings.Repeat(" ", 10000)
	var buf bytes.Buffer
	buf.WriteString(s)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.WriteString(" ")
	}
}

func BenchmarkBytesBufString(b *testing.B) {
	s := strings.Repeat(" ", 10000)
	var buf bytes.Buffer
	buf.WriteString(s)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := buf.String()
		_ = s
	}
}
