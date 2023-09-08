package futil

type Iterable[T interface{}] interface {
	Iterator() Iterator[T]
}

type Iterator[T interface{}] interface {
	HasNext() bool
	Next() T
}

// NewFilteredIterator creates a new, filtered Iterator instance wrapping the
// given Iterator.
//
// A filtered Iterator iterates over only the source values for which the given
// predicate function returns true.
//
// The given predicate function will be called at least once for each call to
// Next.
//
// The source Iterator will be consumed by the returned wrapper Iterator
// instance as the wrapper is consumed.
func NewFilteredIterator[T interface{}](it Iterator[T], predicate func(T) bool) Iterator[T] {
	return &filteredIterator[T]{source: it, filter: predicate}
}

// filteredIterator is an Iterator type that applies a configured predicate to
// the values returned by the source Iterator, returning only those values for
// which the predicate returns true.
type filteredIterator[T interface{}] struct {
	source  Iterator[T]
	hasNext bool
	next    T
	filter  func(T) bool
}

func (f *filteredIterator[T]) HasNext() bool {
	if f.hasNext {
		return true
	}

	for f.source.HasNext() {
		tmp := f.source.Next()

		if f.filter(tmp) {
			f.next = tmp
			f.hasNext = true
			return true
		}
	}

	return false
}

func (f *filteredIterator[T]) Next() T {
	if !f.HasNext() {
		panic("no such element")
	}

	f.hasNext = false
	return f.next
}

func NewMappingIterator[I, O interface{}](source Iterator[I], mapper func(I) O) Iterator[O] {
	return &mappingIterator[I, O]{source, mapper}
}

type mappingIterator[I, O interface{}] struct {
	source Iterator[I]
	mapper func(I) O
}

func (m *mappingIterator[I, O]) HasNext() bool {
	return m.source.HasNext()
}

func (m *mappingIterator[I, O]) Next() O {
	return m.mapper(m.source.Next())
}
