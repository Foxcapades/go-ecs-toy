package fecs

type ComponentConstructor = func() Component

// Component defines the base functionality for a ECS component.
//
// Component implementations must provide a method that returns the type of the
// component.
type Component interface {
	// Type returns the type of this Component.
	Type() ComponentType
}
