package datastructures

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := NewStack[string]()

	s.Push("this")
	s.Push("is")
	s.Push("a")
	s.Push("test")

	for !s.IsEmpty() {
		t.Log(*s.Pop())
	}
}

func TestEmptyStack(t *testing.T) {
	s := NewStack[int]()
	t.Log("Stack is empty:", s.IsEmpty())
}
