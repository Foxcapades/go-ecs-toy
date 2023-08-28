package fecs

// ComponentType is an identifier for a distinct type of component.
type ComponentType uint8

// MaxComponentTypes defines the maximum number of distinct component types that
// may exist.
const MaxComponentTypes = 255

var nextComponentType uint16 = 1

// NewComponentType issues a new ComponentType identifier value.
func NewComponentType() (componentType ComponentType) {
	if nextComponentType == MaxComponentTypes {
		panic("attempted to create more than 255 component types")
	}

	componentType = ComponentType(nextComponentType)
	nextComponentType++
	return
}
