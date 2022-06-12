package stl

func NewTreeMapWithCompare[Key, Val any](cmp func(x, y Key) bool) *TreeMap[Key, Val] {
	return &TreeMap[Key, Val]{
		RbTree[Key, Pair[Key, Val], Select1st[Key, Val]]{keyCompare: cmp},
	}
}

func NewTreeMap[Key Ordered, Val any]() *TreeMap[Key, Val] {
	return &TreeMap[Key, Val]{
		RbTree[Key, Pair[Key, Val], Select1st[Key, Val]]{keyCompare: Less[Key]},
	}
}

type TreeMap[Key, Val any] struct {
	RbTree[Key, Pair[Key, Val], Select1st[Key, Val]]
}
