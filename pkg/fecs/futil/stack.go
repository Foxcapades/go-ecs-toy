package futil

const (
	stackGrowthFactor float32 = 1.5
	trimDifference    int     = 256
	initialCapacity   int     = 32
)

type Stack[T interface{}] interface {
	Size() int
	IsEmpty() bool
	Clear()
	Push(value T)
	Pop() T
}

func NewStack[T interface{}]() Stack[T] {
	return &stack[T]{
		values: make([]T, initialCapacity),
		index:  0,
	}
}

type stack[T interface{}] struct {
	values []T
	index  int
}

func (s *stack[T]) Size() int {
	return s.index
}

func (s *stack[T]) IsEmpty() bool {
	return s.index == 0
}

func (s *stack[T]) Clear() {
	s.index = 0
}

func (s *stack[T]) Push(value T) {
	s.ensureCapacity(s.index)
	s.values[s.index] = value
	s.index++
}

func (s *stack[T]) Pop() T {
	if s.index == 0 {
		panic("attempted to pop an empty stack")
	}
	s.index--
	s.trim()
	return s.values[s.index]
}

func (s *stack[T]) ensureCapacity(size int) {
	newSize := int(float32(len(s.values)) * stackGrowthFactor)

	if newSize < size {
		newSize = size
	}

	if cap(s.values) < size {
		tmp := make([]T, size)
		copy(tmp, s.values)
		s.values = tmp
	}
}

func (s *stack[T]) trim() {
	l := len(s.values)
	if l-s.index > trimDifference {
		tmp := make([]T, l-trimDifference/2)
		copy(tmp, s.values[0:s.index])
		s.values = tmp
	}
}
