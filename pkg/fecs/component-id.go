package fecs

type componentIndex = uint32
type componentVersion = uint32

// ComponentID defines an opaque identifier for a Component.
type ComponentID struct {
	index   componentIndex
	version componentVersion
}

func newComponentID(index componentIndex, version componentVersion) ComponentID {
	return ComponentID{index, version}
}
