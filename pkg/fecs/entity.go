package fecs

import "fmt"

type entityIndex = uint32
type entityVersion = uint32

// EntityID is an opaque identifier representing an entity.
type EntityID struct {
	index   entityIndex
	version entityVersion
}

func (e EntityID) String() string {
	return fmt.Sprintf("EntityID{ %d, %d }", e.index, e.version)
}

type entityDesc struct {
	id   EntityID
	mask componentMask
}

func createEntityID(index entityIndex, version entityVersion) EntityID {
	return EntityID{index, version}
}
