package main

import (
	"testing"
)

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
