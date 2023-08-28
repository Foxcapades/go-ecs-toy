package fecs

// ComponentView defines an iterator over Component instances in a component
// pool.
type ComponentView interface {
	// HasNext returns a boolean flag indicating whether there exists at least one
	// more Component in the underlying pool.
	HasNext() bool

	// Next returns the next available Component instance.
	Next() Component
}

type componentView struct {
	pool  *componentPool
	index int
	next  int
}

func (c *componentView) HasNext() bool {
	for ; c.index < len(c.pool.data); c.index++ {
		if c.pool.data[c.index].component != nil {
			c.next = c.index
			return true
		}
	}

	c.next = -1
	return false
}

func (c *componentView) Next() Component {
	if c.next == -1 && !c.HasNext() {
		panic("no such element")
	}

	out := c.pool.data[c.next].component
	c.next = -1
	c.index++
	return out
}

type emptyComponentView struct{}

func (e emptyComponentView) HasNext() bool {
	return false
}

func (e emptyComponentView) Next() Component {
	panic("no such element")
}

// TypedComponentView defines a ComponentView wrapper that casts the type of the
// returned Component instances to the defined generic type T.
type TypedComponentView[T Component] interface {
	// HasNext returns a boolean flag indicating whether there exists at least one
	// more Component in the underlying ComponentView.
	HasNext() bool

	// Next returns the next available Component instance cast to type T.
	Next() T
}

// NewTypedComponentView wraps a given ComponentView and casts the values
// returned by the wrapped iterator to values of type T.
func NewTypedComponentView[T Component](view ComponentView) TypedComponentView[T] {
	return &typedComponentView[T]{view}
}

type typedComponentView[T Component] struct {
	view ComponentView
}

func (t *typedComponentView[T]) HasNext() bool {
	return t.view.HasNext()
}

func (t *typedComponentView[T]) Next() T {
	return t.view.Next().(T)
}
