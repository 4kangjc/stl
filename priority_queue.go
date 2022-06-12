package stl

type PriorityQueue[T any] struct {
	arr []T
	cmp func(x, y T) bool
}

func NewPriorityQueue[T any](arr []T, cmp func(x, y T) bool) *PriorityQueue[T] {
	pq := &PriorityQueue[T]{
		arr,
		cmp,
	}
	pq.init()
	return pq
}

func (pq *PriorityQueue[T]) init() {
	n := len(pq.arr)
	for i := (n - 1) / 2; i >= 0; i-- {
		pq.sink(i)
	}
}

func (pq *PriorityQueue[T]) swim(k int) {
	for k > 0 && pq.cmp(pq.arr[(k-1)/2], pq.arr[k]) {
		pq.arr[(k-1)/2], pq.arr[k] = pq.arr[k], pq.arr[(k-1)/2]
		k = (k - 1) / 2
	}
}

func (pq *PriorityQueue[T]) sink(k int) {
	n := len(pq.arr)
	for 2*k+1 < n {
		j := 2*k + 1
		if j+1 < n && pq.cmp(pq.arr[j], pq.arr[j+1]) {
			j++
		}
		if !pq.cmp(pq.arr[k], pq.arr[j]) {
			break
		}
		pq.arr[k], pq.arr[j] = pq.arr[j], pq.arr[k]
		k = j
	}
}

func (pq *PriorityQueue[T]) Push(val T) {
	n := len(pq.arr)
	pq.arr = append(pq.arr, val)
	pq.swim(n)
}

func (pq *PriorityQueue[T]) Pop() (val T) {
	val = pq.arr[0]
	pq.arr[len(pq.arr)-1], pq.arr[0] = pq.arr[0], pq.arr[len(pq.arr)-1]
	pq.arr = pq.arr[:(len(pq.arr) - 1)]
	pq.sink(0)
	return
}

func (pq *PriorityQueue[T]) Top() T {
	return pq.arr[0]
}

func (pq *PriorityQueue[T]) Size() int {
	return len(pq.arr)
}

func (pq *PriorityQueue[T]) Empty() bool {
	return pq.Size() == 0
}

func Swap[T any](x, y *T) {
	*x, *y = *y, *x
}

type Comparable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64 | ~string
}

func Less[T Comparable](x, y T) bool {
	return x < y
}

func Greater[T Comparable](x, y T) bool {
	return x > y
}
