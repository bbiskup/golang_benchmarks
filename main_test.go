package main

import (
	"bytes"
	"fmt"
	"github.com/bbiskup/golang_benchmarks/pkg1"
	"strconv"
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

type StrIntItem struct {
	key string
	val int
}

func BenchmarkArrayLookup_100_last(b *testing.B) {
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

func BenchmarkArrayLookup_1000_last(b *testing.B) {
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

type Struct2 struct {
	name  string
	code1 int
	code2 int
}

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

func BenchmarkSprintfIntString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%d", i)
	}
}

/* deadlock
func BenchmarkGoChannelSendReceive(b *testing.B) {
	var ch1 = make(chan int)
	var ch2 = make(chan int)

	go func() {
		fmt.Printf("goroutine: Read\n")
		tmp := <-ch1
		fmt.Printf("goroutine: Write\n")
		ch2 <- tmp
	}()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		fmt.Printf("main: Write\n")
		ch1 <- i
		fmt.Printf("goroutine: Read\n")
		var tmp = <-ch2
		_ = tmp
	}
}*/

func produce(n int, msgs chan<- int, done chan<- bool) {
	for i := 0; i < n; i++ {
		msgs <- i
	}
	done <- true
}

func consume(msgs <-chan int) {
	var last int
	for msg := range msgs {
		//msg := <-msgs
		last = msg
		//fmt.Println(msg)
	}
	fmt.Printf("Last: %d\n", last)
}

func BenchmarkProducerConsumer(b *testing.B) {
	var done = make(chan bool)
	var msgs = make(chan int)
	b.ResetTimer()
	go produce(b.N, msgs, done)
	go consume(msgs)
	<-done
}

func BenchmarkAlloc64KArray(b *testing.B) {
	var arr []byte
	//b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		arr = make([]byte, 65536)
		_ = arr
	}
}

func BenchmarkReadByteFrom64KArray(b *testing.B) {
	const length = 65536
	arr := make([]byte, length)
	for i := 0; i < b.N; i++ {
		value := arr[i%length]
		_ = value
	}
}

func BenchmarkWriteByteTo64KArray(b *testing.B) {
	const length = 65536
	arr := make([]byte, length)
	for i := 0; i < b.N; i++ {
		arr[i%length] = byte(i % 256)
	}
}

type IF interface {
	Name() string
}

type IFImpl1 struct {
}

func (i *IFImpl1) Name() string {
	return "this is IFImpl 1"
}

type IFImpl2 struct {
}

func (i *IFImpl2) Name() string {
	return "this is IFImpl 2"
}

func BenchmarkTypeSwitch(b *testing.B) {
	instance := IFImpl1{}
	var iface IF = &instance
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		switch iface := iface.(type) {
		case *IFImpl1:
			_ = iface.Name()
		case *IFImpl2:
			_ = iface.Name()
		default:
			panic("must not happen")
		}
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
