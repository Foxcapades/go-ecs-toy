package futil

const (
	stackGrowthFactor float32 = 1.5
	trimDifference    int     = 128
	initialCapacity   int     = 32
)

// Stack is a simple stack implementation over generic type T.
type Stack[T interface{}] interface {

	// Size returns the number of elements currently in the stack.
	Size() int

	// IsEmpty returns a boolean indicator for whether the stack is empty or not.
	IsEmpty() bool

	// Clear removes all elements from the stack.
	Clear()

	// Push pushes a new value onto the top of the stack.
	Push(value T)

	// Pop removes and returns the top element from the stack.
	Pop() T
}

// NewStack returns a new Stack instance.
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
	s.values = make([]T, initialCapacity)
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
