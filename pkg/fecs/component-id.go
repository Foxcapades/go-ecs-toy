package fecs

type componentIndex uint32
type ComponentVersion uint32

type ComponentID struct {
	index   componentIndex
	version ComponentVersion
}

var nextComponentID uint32 = 1

func NewComponentID(index componentIndex, version ComponentVersion) ComponentID {
	return ComponentID{index, version}
}
