package fecs

import (
	"fmt"
)

type entity struct {
	id    EntityID
	mask  componentMask
	comps []*ComponentID
}

// birth initializes an entity value for the first time.
func (e *entity) birth(sid SceneID, idx uint32) EntityID {
	e.id.birth(sid, idx)
	return e.id
}

// resurrect re-initializes an entity for reuse.
func (e *entity) resurrect(idx uint32) EntityID {
	e.id.resurrect(idx)
	return e.id
}

// isLiving tests whether this entity value is currently in use.
func (e *entity) isLiving(idx uint32) bool {
	return e.id.isLiving(idx)
}

// kill "destroys" an entity value, clearing everything out of it.
func (e *entity) kill() {
	// Kill the EntityID.
	e.id.kill()

	// Clear the component mask as this instance will no longer have any
	// components attached.
	e.mask.clear()

	// Clear the component id reference slice.
	e.comps = nil
}

// addComponent adds the given ComponentID reference to this entity value.
func (e *entity) addComponent(id *ComponentID) {
	if e.mask.has(id.ctype) {
		panic(fmt.Errorf("attempted to add multiple components of type %s to entity %s", id.ctype.String(), e.id.String()))
	}

	e.mask.add(id.ctype)

	if e.comps == nil {
		e.comps = make([]*ComponentID, 0, 8)
	}

	e.comps = append(e.comps, id)
}

func (e *entity) removeComponent(id *ComponentID) bool {
	idx := e._indexOf(id)

	if idx < 0 {
		return false
	}

	copy(e.comps[idx:], e.comps[idx+1:])

	return true
}

func (e *entity) hasComponent(id *ComponentID) bool {
	if !e.mask.has(id.ctype) {
		return false
	}

	return e._indexOf(id) > -1
}

func (e *entity) hasComponentType(ct ComponentType) bool {
	return e.mask.has(ct)
}

func (e *entity) _indexOf(id *ComponentID) int {
	for i := range e.comps {
		if e.comps[i].Equals(id) {
			return i
		}
	}

	return -1
}

func (e *entity) String() string {
	return fmt.Sprintf("e-%x-%x-%x", e.id.scene, e.id.index, e.id.version)
}
