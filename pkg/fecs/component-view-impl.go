package fecs

// // // // // // // // // // // // // // // // // // // // // // // // // // //
//
// Component View
//
// // // // // // // // // // // // // // // // // // // // // // // // // // //

func newComponentView(components []componentPoolItem) ComponentView {
	return &componentView{
		pool: components,
		next: -1,
	}
}

type componentView struct {
	pool  []componentPoolItem
	next  int
	index int
}

func (c *componentView) HasNext() bool {
	// if the next index has been set then we already know the position of the
	// next value.
	if c.next != -1 {
		return true
	}

	// Iterate through the pool of component refs to find the next one that is not
	// nil.
	for ; c.index < len(c.pool); c.index++ {

		// If we find a ref that is not nil, record its index and return true.
		if c.pool[c.index].ref != nil {
			c.next = c.index
			c.index++
			return true
		}
	}

	// If we exhausted the pool of references and didn't find one that wasn't nil
	// return false.
	return false
}

func (c *componentView) Next() Component {
	if !c.HasNext() {
		panic("no such element")
	}

	out := c.pool[c.next].ref
	c.next = -1
	return out
}

// // // // // // // // // // // // // // // // // // // // // // // // // // //
//
// Empty Component View
//
// // // // // // // // // // // // // // // // // // // // // // // // // // //

func newEmptyComponentView() ComponentView {
	return emptyComponentView{}
}

type emptyComponentView struct{}

func (e emptyComponentView) HasNext() bool {
	return false
}

func (e emptyComponentView) Next() Component {
	panic("no such element")
}

// // // // // // // // // // // // // // // // // // // // // // // // // // //
//
// Typed Component View
//
// // // // // // // // // // // // // // // // // // // // // // // // // // //

// NewTypedComponentView wraps the given ComponentView in a thin wrapper that
// casts the untyped Component instances to type T.
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
