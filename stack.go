package datastructures

type Stack[T any] struct {
	items []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{items: make([]T, 0)}
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Top() *T {
	if s.IsEmpty() {
		var zero *T
		return zero
	}
	return &s.items[len(s.items)-1]
}

func (s *Stack[T]) Pop() *T {
	if s.IsEmpty() {
		var zero *T
		return zero
	}

	val := s.Top()
	s.items = s.items[:len(s.items)-1]
	return val
}

func (s *Stack[T]) Len() int {
	return len(s.items)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}
