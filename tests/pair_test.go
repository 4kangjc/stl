package stl_test

import (
	"github.com/4kangjc/stl"
	"testing"
)

func Test_Pair(t *testing.T) {
	p1 := stl.Pair[int, int]{1, 2}
	if p1.First != 1 && p1.Second != 2 {
		t.Log(p1)
		t.Error("Pair construct error")
	}

	p2 := stl.MakePair("Hello", "World")
	if p2.First != "Hello" && p2.Second != "World" {
		t.Log(p2)
		t.Error("MakePair error")
	}

	p3 := stl.MakePair(&Timer{1, nil}, "Timer")
	if p3.Second != "Timer" && p3.First.now != 1 {
		t.Log(p3)
		t.Error("MakePair Timer error")
	}
}
