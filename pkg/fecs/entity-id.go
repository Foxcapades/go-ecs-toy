package fecs

import "fmt"

// An EntityID is a handle on an entity instance in belonging to a Scene.
type EntityID struct {
	version uint32
	index   uint32
}

// Equals tests whether the given other EntityID is equal to the current
// EntityID.
func (e *EntityID) Equals(o *EntityID) bool {
	return e.index == o.index && e.version == o.version
}

// String returns a string representation of the EntityID.
func (e EntityID) String() string {
	return fmt.Sprintf("%x%016x", e.version, e.index)
}

// newEntityID generates a new EntityID instance from the given values.
func newEntityID(index uint32, version uint32) EntityID {
	return EntityID{
		version: version,
		index:   index,
	}
}
