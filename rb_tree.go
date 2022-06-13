package stl

import "unsafe"

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

func left[Val any](x *RbTreeNodeBase) *RbTreeNode[Val] {
	return (*RbTreeNode[Val])(unsafe.Pointer(x.left))
}

func right[Val any](x *RbTreeNodeBase) *RbTreeNode[Val] {
	return (*RbTreeNode[Val])(unsafe.Pointer(x.right))
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

func (rtn *RbTreeNode[Val]) basePtr() *RbTreeNodeBase {
	return &rtn.RbTreeNodeBase
}

type RbTree[Key, Val any, KOV IKeyofValue[Key, Val]] struct {
	RbTreeHeader
	keyCompare func(x, y Key) bool
}

func (rbt *RbTree[Key, Val, KOV]) root() *RbTreeNodeBase {
	return rbt.header.parent
}

func (rbt *RbTree[Key, Val, KOV]) leftMost() *RbTreeNodeBase {
	return rbt.header.left
}

func (rbt *RbTree[Key, Val, KOV]) rightMost() *RbTreeNodeBase {
	return rbt.header.right
}

func (rbt *RbTree[Key, Val, KOV]) begin() *RbTreeNode[Val] {
	return (*RbTreeNode[Val])(unsafe.Pointer(rbt.header.parent))
}

func (rbt *RbTree[Key, Val, KOV]) end() *RbTreeNode[Val] {
	return (*RbTreeNode[Val])(unsafe.Pointer(&rbt.header))
}

func (rbt *RbTree[Key, Val, KOV]) Find(key Key) RbTreeIterator[Val] {
	j := rbt.lower_bound(rbt.begin(), rbt.end().basePtr(), key)
	if j != rbt.End() && rbt.keyCompare(key, bKey[Key, Val, KOV](j.node)) {
		j = rbt.End()
	}
	return j
}

//
//func (rbt *RbTree[Key, Val, KOV]) Erase(key Key) {}
//
//func (rbt *RbTree[Key, Val, KOV]) Set(key Key, val Val) {}
//
//func (rbt *RbTree[Key, Val, KOV]) Get(key Key) (val Val) {
//	return
//}

func bKey[Key, Val any, KOV IKeyofValue[Key, Val]](y *RbTreeNodeBase) Key {
	return lKey[Key, Val, KOV]((*RbTreeNode[Val])(unsafe.Pointer(y)))
}

func lKey[Key, Val any, KOV IKeyofValue[Key, Val]](y *RbTreeNode[Val]) Key {
	var kov KOV
	//return kov.KeyofValue(*y.valPtr())
	return kov.KeyofValue(y.valueField)
}

//func sKey[Key, Val any, KOV IKeyofValue[Key, Val], Ptr *RbTreeNodeBase | *RbTreeNode[Val]](y Ptr) Key {
//	var kov KOV
//	return kov.KeyofValue((*RbTreeNode[Val])(unsafe.Pointer(y)).valueField)
//}

func (rbt *RbTree[Key, Val, KOV]) lower_bound(x *RbTreeNode[Val], y *RbTreeNodeBase, k Key) RbTreeIterator[Val] {
	for x != nil {
		if !rbt.keyCompare(lKey[Key, Val, KOV](x), k) {
			//y, x = (*RbTreeNodeBase)(unsafe.Pointer(x)), (*RbTreeNode[Val])(unsafe.Pointer(x.left))
			y, x = x.basePtr(), left[Val](x.basePtr())
		} else {
			//x = (*RbTreeNode[Val])(unsafe.Pointer(x.right))
			x = right[Val](x.basePtr())
		}
	}
	return RbTreeIterator[Val]{x.basePtr()}
	//return RbTreeIterator[Val]{(*RbTreeNodeBase)(unsafe.Pointer(x))}
}

func (rbt *RbTree[Key, Val, KOV]) upper_bound(x *RbTreeNode[Val], y *RbTreeNodeBase, k Key) RbTreeIterator[Val] {
	for x != nil {
		if !rbt.keyCompare(k, lKey[Key, Val, KOV](x)) {
			y, x = x.basePtr(), left[Val](x.basePtr())
		} else {
			x = right[Val](x.basePtr())
		}
	}
	return RbTreeIterator[Val]{y}
}

func (rbt *RbTree[Key, Val, KOV]) Begin() RbTreeIterator[Val] {
	return RbTreeIterator[Val]{rbt.header.parent}
}

func (rbt *RbTree[Key, Val, KOV]) End() RbTreeIterator[Val] {
	return RbTreeIterator[Val]{&rbt.header}
}

type RbTreeIterator[Tp any] struct {
	node *RbTreeNodeBase
}

func (it *RbTreeIterator[Tp]) Next() {

}

func (it *RbTreeIterator[Tp]) Prev() {

}

func (it *RbTreeIterator[Tp]) Value() Tp {
	return (*RbTreeNode[Tp])(unsafe.Pointer(it.node)).valueField
}

func (it *RbTreeIterator[Tp]) Pointer() *Tp {
	return (*RbTreeNode[Tp])(unsafe.Pointer(it.node)).valPtr()
}
