package fecs

import "fmt"

// ComponentType is an identifier for a distinct component type.
type ComponentType uint8

func (c ComponentType) String() string {
	return fmt.Sprintf("ComponentType#%d", c)
}

// MaxComponentType defines the maximum number of distinct component types that
// may be created.
const MaxComponentType uint8 = 255

var curComponentType uint8 = 0

// NewComponentType issues a new ComponentType value.
func NewComponentType() ComponentType {
	if curComponentType == MaxComponentType {
		panic(fmt.Sprintf("attempted to create more than %d distinct component types!", MaxComponentType))
	}

	curComponentType++
	return ComponentType(curComponentType)
}
