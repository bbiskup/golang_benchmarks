package main

import (
	"github.com/bbiskup/golang_benchmarks/pkg1"
	"strings"
	"testing"
)

type S struct {
}

func (s *S) getValMethod(value int) int {
	return value + 1
}

func BenchmarkEmpty(b *testing.B) {
	for i := 0; i < b.N; i++ {

	}
}

type SI interface {
	getValMethod(value int) int
}

type S2 struct {
	S
}

func BenchmarkAddLoop(b *testing.B) {
	var value int
	for i := 0; i < b.N; i++ {
		value++
	}
	_ = value
	// fmt.Printf("Value: %d\n", value)
}

func getValPlus1(val int) int {
	return val + 1
}

func BenchmarkReuseResult(b *testing.B) {
	var value int
	for i := 0; i < b.N; i++ {
		value = getValPlus1(i)
		_ = value + 1
		_ = value + 2
		_ = value + 3
	}
	_ = value
	// fmt.Printf("Value: %d\n", value)
}

func BenchmarkFuncCallSamePkg(b *testing.B) {
	var value int
	for i := 0; i < b.N; i++ {
		value = getValPlus1(value)
	}
	_ = value
	//fmt.Printf("Value: %d\n", value)
}

func BenchmarkFuncCallByVar(b *testing.B) {
	fn := getValPlus1
	var value int
	for i := 0; i < b.N; i++ {
		value = fn(value)
	}
	_ = value
}

func BenchmarkStructMethodSamePkg(b *testing.B) {
	s := &S{}
	var value int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		value = s.getValMethod(value)
	}
	_ = value
	// fmt.Printf("Value: %d\n", value)
}

func BenchmarkInterfaceMethodSamePkg(b *testing.B) {
	var si SI = &S{}
	var value int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		value = si.getValMethod(value)
	}
	_ = value
	// fmt.Printf("Value: %d\n", value)
}

// Should take twice as long as long as method call
// on struct due to indirection (see above)
func BenchmarkStructMethodSamePkgViaEmbeddedImplicit(b *testing.B) {
	s2 := &S2{}
	var value int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		value = s2.getValMethod(value)
	}
	_ = value
	// fmt.Printf("Value: %d\n", value)
}

func BenchmarkStructMethodSamePkgViaEmbeddedExplicit(b *testing.B) {
	s2 := &S2{}
	var value int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		value = s2.S.getValMethod(value)
	}
	_ = value
	// fmt.Printf("Value: %d\n", value)
}

func BenchmarkFuncCallOtherPkg(b *testing.B) {
	var value int
	for i := 0; i < b.N; i++ {
		value = pkg1.GetValPlus1_b(value)
	}
	// fmt.Printf("Value: %d\n", value)
}

func BenchmarkClosureCall(b *testing.B) {
	fn := func(val int) int {
		return val + 1
	}
	var value int
	for i := 0; i < b.N; i++ {
		value = fn(value)
	}
	// fmt.Printf("Value: %d\n", value)
}

func BenchmarkInlineAnonymousFunctionCall(b *testing.B) {
	var value int
	for i := 0; i < b.N; i++ {
		value = func(val int) int {
			return val + 1
		}(value)
	}
	// fmt.Printf("Value: %d\n", value)
}

func BenchmarkInlineClosureCall(b *testing.B) {
	var num int = 3
	var value int
	for i := 0; i < b.N; i++ {
		value = func(val int) int {
			return val + num
		}(value)
	}
	// fmt.Printf("Value: %d\n", value)
}

type StrStruct struct {
	s string
}

func funStructByVal(s StrStruct) uint8 {
	return s.s[0]
}

func funStructByRef(s *StrStruct) uint8 {
	return s.s[0]
}

func BenchmarkCallFuncWithStructByValue(b *testing.B) {
	s := StrStruct{strings.Repeat("x", 100000)}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = funStructByVal(s)
	}
}

func BenchmarkCallFuncWithStructByRef(b *testing.B) {
	s := StrStruct{strings.Repeat("x", 100000)}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = funStructByRef(&s)
	}
}
