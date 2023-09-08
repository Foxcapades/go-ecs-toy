package futil

// MapIfPresent calls the given mapping function on the given value and returns
// the result if and only if the given value is not nil.
//
// It doesn't make sense to call this function on values that are not nillable.
//
// Returns either the mapped value or the default value for the type R depending
// on whether the given input value is nil.
func MapIfPresent[T, R interface{}](value T, mapper func(T) R) (out R) {
	if value != nil {
		out = mapper(value)
	}

	return
}

// CallIfPresent calls the given function on the given value if and only if that
// value is not nil.
//
// It doesn't make sense to call this function on values that are not nillable.
func CallIfPresent[T interface{}](value T, fun func(T)) {
	if value != nil {
		fun(value)
	}
}

func IfElse[T interface{}](condition bool, ifTrue, ifFalse T) T {
	if condition {
		return ifTrue
	} else {
		return ifFalse
	}
}

func Apply[T interface{}](value T, fun func(T)) T {
	fun(value)
	return value
}

func With[T, R interface{}](value T, mapper func(T) R) R {
	return mapper(value)
}

func SliceIsEmpty[T interface{}](slice []T) bool {
	return len(slice) == 0
}

func IsMapEmpty[K comparable, V interface{}](mp map[K]V) bool {
	return len(mp) == 0
}
