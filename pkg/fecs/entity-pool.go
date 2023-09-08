package fecs

import (
	"github.com/Foxcapades/go-ecs-toy/pkg/fecs/futil"
)

func newEntityPool() (out entityPool) {
	out.free = futil.NewStack[uint32]()
	out.pool = make([]entity, 32)
	return
}

type entityPool struct {
	free futil.Stack[uint32]
	pool []entity
	size uint32
}

func (e *entityPool) addComponent(eid *EntityID, cid *ComponentID) {
	if !e.containsEntity(eid) {
		panic("illegal state")
	}

	e.pool[eid.index].addComponent(cid)
}

// containsEntity tests whether this entityPool contains an entity matching the
// given EntityID.
func (e *entityPool) containsEntity(id *EntityID) bool {
	return id.index < e.size && e.pool[id.index].is(id)
}

// entities returns an Iterator over the EntityIDs of the entities contained in
// this entityPool, optionally filtered to only those entities that have
// components of all the given ComponentTypes attached.
//
// Altering this entityPool while an entity Iterator is in use may cause
// undefined behavior.
func (e *entityPool) entities(ct []ComponentType) futil.Iterator[EntityID] {
	it := &entityIterator{pool: e.pool}

	// If no component type filters were specified, then return the raw iterator.
	if futil.SliceIsEmpty(ct) {
		// Map the raw
		return futil.NewMappingIterator[*entity, EntityID](it, _entityIteratorMapper)
	}

	// Build a mask for the filtered iterator predicate.
	mask := componentMask{}
	for _, t := range ct {
		mask.add(t)
	}

	// Apply the mask filter to the source iterator.
	filtered := futil.NewFilteredIterator[*entity](it, func(e *entity) bool { return e.mask.hasAll(&mask) })

	// Map the raw *entity values to EntityID values.
	return futil.NewMappingIterator[*entity, EntityID](filtered, _entityIteratorMapper)
}

// entityCount returns the number of entities currently in this entityPool.
func (e *entityPool) entityCount() uint32 {
	return e.size - uint32(e.free.Size())
}

func (e *entityPool) entityHasComponentType(id *EntityID, ct ComponentType) bool {
	if !e.containsEntity(id) {
		return false
	}

	return e.pool[id.index].mask.has(ct)
}

func (e *entityPool) getEntityComponents(id *EntityID) []*ComponentID {
	if e.containsEntity(id) {
		return e.pool[id.index].comps
	}

	return nil
}

func (e *entityPool) getEntityComponent(id *EntityID, ct ComponentType) (*ComponentID, bool) {
	if !e.containsEntity(id) {
		return nil, false
	}

	ent := &e.pool[id.index]

	if !ent.mask.has(ct) {
		return nil, false
	}

	for i := range ent.comps {
		if ent.comps[i].ctype == ct {
			return ent.comps[i], true
		}
	}

	panic("illegal state")
}

// newEntity creates a new entity instance and returns its EntityID.
func (e *entityPool) newEntity(id SceneID) EntityID {
	if e.free.IsEmpty() {
		return e._append(id)
	} else {
		return e._overwrite()
	}
}

// removeEntity removes the entity identified by the given EntityID from this
// entityPool, returning a boolean value that indicates whether the target
// entity was in this entityPool to begin with.-
func (e *entityPool) removeEntity(id *EntityID) bool {
	if !e.containsEntity(id) {
		return false
	}

	e.pool[id.index].kill()
	e.free.Push(id.index)

	return true
}

// _append adds a new entity to the end of the entityPool.
func (e *entityPool) _append(id SceneID) EntityID {
	e.pool[e.size].birth(id, e.size)
	e.size++

	return e.pool[e.size-1].id
}

// _overwrite overwrites an unused entity in the pool with a new entity value.
func (e *entityPool) _overwrite() EntityID {
	idx := e.free.Pop()
	return e.pool[idx].resurrect(idx)
}

// _ensureCapacity ensures that the entityPool's pool capacity is greater than
// or equal to the given minimum capacity.
//
// If the current capacity is already greater than or equal to the given minimum
// value, this method does nothing.
//
// If the current capacity is less than the given minimum value, this method
// will replace the pool slice with a new, larger pool slice.
func (e *entityPool) _ensureCapacity(minimum uint32) {
	if minimum < uint32(len(e.pool)) {
		return
	}

	newSize := max(uint32(float32(len(e.pool))*1.5), minimum)

	tmp := make([]entity, newSize)
	copy(tmp, e.pool)
	e.pool = tmp
}

// _entityIteratorMapper maps a given entity value to it's EntityID.
//
// This function is meant to be used with a mapping Iterator to convert the
// internal entityPool items to values that can be passed back to the external
// caller.
func _entityIteratorMapper(e *entity) EntityID {
	return e.id
}
