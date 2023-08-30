package fecs

// ComponentView is an iterator over a set of Component instances.
type ComponentView interface {

	// HasNext returns whether there exists at least one more Component instance
	// in the underlying collection.
	HasNext() bool

	// Next returns the next Component instance in the underlying collection and
	// moves the iterator index forward by one.
	//
	// If Next is called when there are no more Component instances remaining in
	// the view, this method will panic.
	//
	// Basic usage:
	//   for view.HasNext() {
	//     doSomething(view.Next())
	//   }
	Next() Component
}

// TypedComponentView is an iterator over a set of Component instances that
// casts each returned Component to type T.
//
// This is useful when the type of the components being iterated over is known
// ahead of time.
//
// Basic usage:
//   untypedView := scene.Components(ComponentTypeTransform)
//   transforms := NewTypedComponentView[Transform](untypedView)
type TypedComponentView[T Component] interface {

	// HasNext returns whether there exists at least one more Component instance
	// in the underlying ComponentView.
	HasNext() bool

	// Next returns the next Component instance in the underlying collection, cast
	// to type T, and moves the iterator index forward by one.
	//
	// If Next is called when there are no more Component instances remaining in
	// the view, this method will panic.
	//
	// If the Component instances in the underlying collection are not of type T,
	// this method will panic.
	Next() T
}
