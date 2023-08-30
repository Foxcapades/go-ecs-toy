package fecs

import "fmt"

// ComponentID is a handle on a Component instance that has been registered to
// a Scene entity.
type ComponentID struct {
	index   uint32
	version uint32
	cType   ComponentType
}

func (c *ComponentID) Type() ComponentType {
	return c.cType
}

// Equals test whether
func (c *ComponentID) Equals(other *ComponentID) bool {
	return c.index == other.index && c.version == other.version && c.cType == other.cType
}

// String returns a string representation of the current ComponentID.
func (c *ComponentID) String() string {
	return fmt.Sprintf("%d%016x%016x", c.cType, c.version, c.index)
}

// newComponentID creates a new ComponentID from the given values.
func newComponentID(index, version uint32, cType ComponentType) ComponentID {
	return ComponentID{index, version, cType}
}
