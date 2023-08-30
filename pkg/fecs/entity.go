package fecs

type entity struct {
	id   EntityID
	mask ComponentMask

	components []ComponentID
}

// addComponent adds the given ComponentID to the current entity.
func (e *entity) addComponent(cid ComponentID) {
	if e.mask.Has(cid.cType) {
		panic("attempted to add multiple components of type " + cid.cType.String() + " to entity " + e.id.String())
	}

	e.mask.Add(cid.cType)
	e.components = append(e.components, cid)
}

// hasComponent tests whether the current entity has a Component with the given
// ComponentID.
func (e *entity) hasComponent(cid ComponentID) bool {
	return e.mask.Has(cid.cType)
}

// hasComponentType tests whether the current entity has a Component of the
// given ComponentType.
func (e *entity) hasComponentType(ctp ComponentType) bool {
	return e.mask.Has(ctp)
}

// findComponentID looks up the ComponentID for the component with the given
// ComponentType attached to the current entity.
//
// If this entity does not have a Component of the given ComponentType, the
// returned values will be an invalid ComponentID and the boolean value false.
//
// If this entity does have a Component of the given ComponentType, the returned
// values will be an invalid ComponentID
func (e *entity) findComponentID(ctp ComponentType) (ComponentID, bool) {
	if !e.mask.Has(ctp) {
		return ComponentID{}, false
	}

	for i := range e.components {
		if e.components[i].cType == ctp {
			return e.components[i], true
		}
	}

	panic("illegal state, mask suggests entity " + e.id.String() +
		" has component of type " + ctp.String() +
		" however no such component was found in the component slice")
}

// removeComponentByID removes a Component instance from the current entity if
// it exists.
//
// If the target Component was found on this entity and removed, this method
// returns true.  If the target Component was not found on this entity, this
// method returns false.
func (e *entity) removeComponentByID(id ComponentID) bool {
	return e.removeComponentByType(id.cType)
}

// removeComponentByType removes a Component instance from the current entity if
// it exists.
//
// If the target Component was found on this entity and removed, this method
// returns true.  If the target Component was not found on this entity, this
// method returns false.
func (e *entity) removeComponentByType(t ComponentType) bool {
	if !e.mask.Has(t) {
		return false
	}

	for i := range e.components {
		if e.components[i].cType == t {
			copy(e.components[i:], e.components[i+1:])
			return true
		}
	}

	return false
}
