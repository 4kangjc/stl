package stl

type Queue[T any] []T

func (q *Queue[T]) Push(val T) {
	*q = append(*q, val)
}

func (q *Queue[T]) Pop() (val T) {
	val = (*q)[0]
	*q = (*q)[1:]
	return
}

func (q *Queue[T]) Front() T {
	return (*q)[0]
}

func (q *Queue[T]) Size() int {
	return len(*q)
}

func (q *Queue[T]) Empty() bool {
	return q.Size() == 0
}
