package stl

type IKeyofValue[Key, Val any] interface {
	KeyofValue(Val) Key
}

type Pair[Tp, Ty any] struct {
	first  Tp
	second Ty
}

type Select1st[Key, Val any] struct{}
type Select2st[Key, Val any] struct{}

type Identity[Val any] struct{}

func (Identity[Val]) KeyofValue(val Val) Val {
	return val
}

func (Select1st[Key, Val]) KeyofValue(val Pair[Key, Val]) Key {
	return val.first
}

func (Select2st[Key, Val]) KeyofValue(val Pair[Key, Val]) Val {
	return val.second
}

func MakePair[Tp any, Ty any](first Tp, second Ty) Pair[Tp, Ty] {
	return Pair[Tp, Ty]{first, second}
}
