package fecs

// A ComponentConstructor is a function that takes no arguments and returns a
// new Component instance.
type ComponentConstructor = func() Component

type Component interface {
	Type() ComponentType
}
