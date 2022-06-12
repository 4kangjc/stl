package stl

type RbTreeColor bool

const (
	red, black RbTreeColor = false, true
)

type RbTreeNodeBase struct {
	color  RbTreeColor
	left   *RbTreeNodeBase
	right  *RbTreeNodeBase
	parent *RbTreeNodeBase
}

func minimum(x *RbTreeNodeBase) *RbTreeNodeBase {
	for x.left != nil {
		x = x.left
	}
	return x
}

func maximum(x *RbTreeNodeBase) *RbTreeNodeBase {
	for x.right != nil {
		x = x.right
	}
	return x
}

type RbTreeHeader struct {
	header   RbTreeNodeBase
	nodeSize uint64
}

type RbTreeNode[Val any] struct {
	RbTreeNodeBase
	valueField Val
}

func (rtn *RbTreeNode[Val]) valPtr() *Val {
	return &rtn.valueField
}

type RbTree[Key, Val any, KOV IKeyofValue[Key, Val]] struct {
	RbTreeHeader
	keyCompare func(x, y Key) bool
}

func (rbt *RbTree[Key, Val, KOV]) Find(key Key) RbTreeIterator {
	return RbTreeIterator{}
}

func (rbt *RbTree[Key, Val, KOV]) Erase(key Key) {}

func (rbt *RbTree[Key, Val, KOV]) Set(key Key, val Val) {}

func (rbt *RbTree[Key, Val, KOV]) Get(key Key) (val Val) {
	return
}

func (rbt *RbTree[Key, Val, KOV]) Begin() (iterator RbTreeIterator) {
	return
}

func (rbt *RbTree[Key, Val, KOV]) End() (iterator RbTreeIterator) {
	return
}

type RbTreeIterator struct {
}

func (it *RbTreeIterator) Next() {

}

func (it *RbTreeIterator) Prev() {

}
