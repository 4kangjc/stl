package stl

type Stack[T any] []T

func (s *Stack[T]) Push(val T) {
	*s = append(*s, val)
}

func (s *Stack[T]) Pop() (val T) {
	val = (*s)[len(*s)-1]
	*s = (*s)[:(len(*s) - 1)]
	return
}

func (s *Stack[T]) Top() T {
	return (*s)[s.Size()-1]
}

func (s *Stack[T]) Size() int {
	return len(*s)
}

func (s *Stack[T]) Empty() bool {
	return s.Size() == 0
}
