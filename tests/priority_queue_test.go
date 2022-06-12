package stl_test

import (
	"github.com/4kangjc/stl"
	"testing"
)

type Timer struct {
	now int64
	cb  func()
}

func TestPriorityQueue(t *testing.T) {
	pq := stl.NewPriorityQueue([]int{1, 2, 3, 4}, stl.Less[int])
	pq.Push(1)
	pq.Push(55)
	pq.Push(9)
	pq.Push(100)

	arr := []int{}
	for !pq.Empty() {
		//Info(pq.Pop())
		arr = append(arr, pq.Pop())
	}
	t.Log(arr)
	if !stl.IsSorted(arr, stl.Greater[int]) {
		t.Error("priority queue is error")
	}

	arr2 := []*Timer{}
	pq2 := stl.NewPriorityQueue(nil, func(x, y *Timer) bool {
		return x.now > y.now
	})
	pq2.Push(&Timer{100, nil})
	pq2.Push(&Timer{1, nil})
	pq2.Push(&Timer{9, nil})
	pq2.Push(&Timer{55, nil})

	t.Log(pq2)
	for !pq2.Empty() {
		//Info(pq2.Pop())
		arr2 = append(arr2, pq2.Top())
		pq2.Pop()
	}

	if !stl.IsSorted(arr2, func(x, y *Timer) bool {
		return x.now < y.now
	}) {
		t.Error("priority queue error")
	}

	pq3 := stl.NewPriorityQueue([]string{}, stl.Greater[string])
	pq3.Push("Hello")
	pq3.Push("World")

	if pq3.Size() != 2 {
		t.Log(pq3)
		t.Error("pq3 size is not 2")
	}
}
