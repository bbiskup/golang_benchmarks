package main

import (
	"testing"
)

func (s *Struct2) value(val int) int {
	return val + 1
}

func (s *Struct2) GetCode1() int { return s.code1 }

func BenchmarkConstructStructWithMembers(b *testing.B) {
	var p interface{}
	for i := 0; i < b.N; i++ {
		p = &Struct2{}
		_ = p
	}
}

func BenchmarkConstructStructWithMembers2(b *testing.B) {
	var p interface{}
	for i := 0; i < b.N; i++ {
		p = new(Struct2)
		_ = p
	}
}

func BenchmarkMethodCallNoMemberAccess(b *testing.B) {
	s2 := Struct2{}
	var value int
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		value = s2.value(value)
	}
	// fmt.Printf("Value: %d\n", value)
}

func BenchmarkAttributeAccess(b *testing.B) {
	s2 := Struct2{}
	for i := 0; i < b.N; i++ {
		_ = s2.code1
	}
}

func BenchmarkAttributeAccessViaMemberFunc(b *testing.B) {
	s2 := Struct2{}
	for i := 0; i < b.N; i++ {
		_ = s2.GetCode1()
	}
}
