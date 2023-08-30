package fecs

import "github.com/Foxcapades/go-ecs-toy/pkg/fecs/futil"

// NewScene creates a new Scene instance.
func NewScene() Scene {
	return &scene{
		free:  futil.NewStack[uint32](),
		pools: make(map[ComponentType]ComponentPool),
	}
}

type scene struct {
	free     futil.Stack[uint32]
	pools    map[ComponentType]ComponentPool
	entities []entity
	index    uint32
}

func (s *scene) NewEntity() EntityID {
	if s.free.IsEmpty() {
		s.ensureCapacity(s.index)

		newID := newEntityID(s.index, 1)

		s.entities[s.index].id = newID
		s.index++

		return newID
	}

	index := s.free.Pop()
	newID := newEntityID(index, s.entities[index].id.version)

	s.entities[index].id = newID

	return newID
}

func (s *scene) AssignComponent(entityID EntityID, cons ComponentConstructor) ComponentID {
	if !s.isValidEntityID(entityID) {
		panic("attempted to assign a new component to the dead or invalid entity id " + entityID.String())
	}

	component := cons()
	entity := &s.entities[entityID.index]

	if entity.mask.Has(component.Type()) {
		panic("attempted to attach multiple components of type " + component.Type().String() + " to entity " + entityID.String())
	}

	var pool ComponentPool

	if p, ok := s.pools[component.Type()]; ok {
		pool = p
	} else {
		pool = newComponentPool()
		s.pools[component.Type()] = pool
	}

	id := pool.Add(component)
	entity.addComponent(id)
	return id
}

func (s *scene) GetComponentByType(eid EntityID, ctp ComponentType) (Component, bool) {
	if !s.isValidEntityID(eid) {
		return nil, false
	}

	if cid, ok := s.entities[eid.index].findComponentID(ctp); ok {
		return s.pools[ctp].Get(cid)
	} else {
		return nil, false
	}
}

func (s *scene) GetComponentByID(cid ComponentID) (Component, bool) {
	return s.pools[cid.cType].Get(cid)
}

func (s *scene) RemoveEntity(eid EntityID) bool {
	if !s.isValidEntityID(eid) {
		return false
	}

	entity := &s.entities[eid.index]
	entity.mask.Clear()
	entity.id = newEntityID(eid.index, entity.id.version+1)

	for i := range entity.components {
		s.pools[entity.components[i].cType].Remove(entity.components[i])
	}

	entity.components = nil

	return true
}

func (s *scene) RemoveComponentByType(eid EntityID, ctp ComponentType) bool {
	if !s.isValidEntityID(eid) {
		return false
	}

	entity := &s.entities[eid.index]
	if cid, ok := entity.findComponentID(ctp); ok {
		entity.removeComponentByType(ctp)
		return s.pools[ctp].Remove(cid)
	}

	return false
}

func (s *scene) RemoveComponentByID(eid EntityID, cid ComponentID) bool {
	if !s.isValidEntityID(eid) {
		return false
	}

	s.entities[eid.index].removeComponentByID(cid)
	return s.pools[cid.cType].Remove(cid)
}

func (s *scene) EntityHasComponent(entityID EntityID, componentType ComponentType) bool {
	return s.isValidEntityID(entityID) &&
		s.entities[entityID.index].hasComponentType(componentType)
}

func (s *scene) Components(ctp ComponentType) ComponentView {
	if pool, ok := s.pools[ctp]; ok {
		return pool.ComponentView()
	} else {
		return newEmptyComponentView()
	}
}

func (s *scene) EntityCount() int {
	return len(s.entities) - s.free.Size()
}

func (s *scene) Entities(ctp ComponentType) EntityView {
	return newEntityView(s.entities, func(e *entity) bool { return e.hasComponentType(ctp) })
}

func (s *scene) ensureCapacity(minimum uint32) {
	if uint32(len(s.entities)) >= minimum {
		return
	}

	newSize := uint32(float32(len(s.entities)) * 1.5)

	if newSize < minimum {
		newSize = minimum
	}

	tmp := make([]entity, newSize)
	copy(tmp, s.entities)
	s.entities = tmp
}

func (s *scene) isValidEntityID(i EntityID) bool {
	return i.index < uint32(len(s.entities)) &&
		s.entities[i.index].id.Equals(&i)
}
