package stl

func NewTreeSetWithCompare[Key, Val any](cmp func(x, y Key) bool) *TreeSet[Key] {
	return &TreeSet[Key]{
		RbTree[Key, Key, Identity[Key]]{keyCompare: cmp},
	}
}

func NewTreeSet[Key Ordered, Val any]() *TreeSet[Key] {
	return &TreeSet[Key]{
		RbTree[Key, Key, Identity[Key]]{keyCompare: Less[Key]},
	}
}

type TreeSet[Key any] struct {
	RbTree[Key, Key, Identity[Key]]
}
