package stl_test

import (
	"github.com/4kangjc/stl"
	"testing"
)

func TestStack_Empty(t *testing.T) {
	s := stl.Stack[int]{}
	if !s.Empty() {
		t.Error("stack is not empty")
	}
}

func TestStack_Push(t *testing.T) {
	s := stl.Stack[string]{}
	s.Push("Hello")
	if s.Size() != 1 {
		t.Error("stack size is not 1")
	}
	s.Push("World")
	if s.Size() != 2 {
		t.Error("stack size is not 2")
	}
}

func TestStack_Pop(t *testing.T) {
	type Person struct {
		id   int
		name string
	}
	s := stl.Stack[*Person]{}
	s.Push(&Person{
		0,
		"zero",
	})
	if s.Pop().name != "zero" {
		t.Error("stack pop error")
	}
	if !s.Empty() {
		t.Error("stack size is not empty, stack pop error")
	}
}

func TestStack_Top(t *testing.T) {
	s := stl.Stack[int]{}
	for i := 0; i < 5; i++ {
		s.Push(i)
	}
	for i := 4; !s.Empty(); i-- {
		if s.Top() != i {
			t.Error("stack top is not ", i)
		}
		s.Pop()
	}
}
