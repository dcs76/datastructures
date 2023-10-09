package datastructures

import (
	"testing"
)

func TestSet(t *testing.T) {
	s := NewSet[string]()

	s.Add("this")
	s.Add("is")
	s.Add("a")
	s.Add("test")
	s.Add("of")
	s.Add("a")
	s.Add("test")

	if s.Len() != 5 {
		t.Fatal("wrong number of elements")
	}

	if !s.In("of") {
		t.Fatal("'of' should be in set")
	}

	if s.In("not there") {
		t.Fatal("'not there' should not be in the set")
	}
}
