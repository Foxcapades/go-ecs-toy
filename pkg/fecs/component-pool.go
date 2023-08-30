package fecs

// ComponentPool is a contiguous block of memory containing zero or more
// Component instances.
type ComponentPool interface {

	// Size returns the number of Component instances currently in this
	// ComponentPool.
	Size() int

	// Get attempts to retrieve a component from the pool by the given
	// ComponentID.  If the target Component was not found, returns nil and false.
	// If the target Component was found, returns the target component and true.
	Get(componentID ComponentID) (Component, bool)

	// Has tests whether this ComponentPool contains the component with the given
	// ComponentID.
	Has(componentID ComponentID) bool

	// Add adds the target Component to this ComponentPool, generating and
	// returning a ComponentID for the Component.
	Add(component Component) ComponentID

	// Remove removes the target Component from this ComponentPool.  If the target
	// Component was not found, returns false.  If the target Component was found
	// and removed, returns true.
	Remove(componentID ComponentID) bool

	// ComponentView returns a new ComponentView instance over the contents of
	// this ComponentPool.
	ComponentView() ComponentView
}
