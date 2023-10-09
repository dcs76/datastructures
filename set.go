package datastructures

type Set[T comparable] struct {
	items map[T]empty
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{items: make(map[T]empty, 0)}
}

func (s *Set[T]) Add(item T) {
	s.items[item] = empty{}
}

func (s *Set[T]) In(item T) bool {
	if _, ok := s.items[item]; ok {
		return true
	}
	return false
}

func (s *Set[T]) Len() int {
	return len(s.items)
}
