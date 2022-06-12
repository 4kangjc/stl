package stl

type Complex interface {
	~complex64 | ~complex128
}

type Float interface {
	~float32 | ~float64
}

type Integer interface {
	Signed | Unsigned
}

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Ordered interface {
	Integer | Float | ~string
}

func Less[T Ordered](x, y T) bool {
	return x < y
}

func Greater[T Ordered](x, y T) bool {
	return x > y
}

func IsSorted[T any](arr []T, cmp func(x, y T) bool) bool {
	for i := range arr {
		if i > 0 && cmp(arr[i], arr[i-1]) {
			return false
		}
	}
	return true
}
