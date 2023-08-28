package fecs

import "github.com/Foxcapades/go-ecs-toy/pkg/fecs/futil"

type scene struct {
	freeEntities futil.Stack[entityIndex]
	entities     []entityDesc
	cPools       map[ComponentType]*componentPool
}

func (s *scene) NewEntity() (id EntityID) {
	// If we have an unused entity already in the slice, reuse it, else append a
	// new one.
	if !s.freeEntities.IsEmpty() {
		index := s.freeEntities.Pop()

		id = createEntityID(index, s.entities[index].id.version)
		s.entities[index].id = id
	} else {
		index := entityIndex(len(s.entities))

		id = createEntityID(index, 1)
		s.entities = append(s.entities, entityDesc{id: id})
	}

	return
}

func (s *scene) RemoveEntity(id EntityID) {
	idx := id.index

	if s.entityIdIsValid(id) {
		s.entities[idx].mask.reset()
		s.entities[idx].id.version++
		s.freeEntities.Push(idx)
	}
}

func (s *scene) AssignComponent(id EntityID, component Component) ComponentID {
	if !s.entityIdIsValid(id) {
		panic("attempted to attach a component to an invalid or deleted entity")
	}

	s.entities[id.index].mask.set(component.GetType())

	if pool, ok := s.cPools[component.GetType()]; ok {
		return pool.add(component)
	}

	pool := newComponentPool()
	s.cPools[component.GetType()] = pool
	return pool.add(component)
}

func (s *scene) RemoveComponent(id ComponentID) {
	if pool, ok := s.cPools[id.cType]; ok {
		pool.remove(id)
	}
}

func (s *scene) RemoveComponentFromEntity(entityID EntityID, componentID ComponentID) {
	if s.entityIdIsValid(entityID) {
		s.entities[entityID.index].mask.unset(componentID.cType)
	}
}

func (s *scene) EntitiesWith(componentType ComponentType) EntityView {
	return &entityView{
		scene: s,
		cType: componentType,
		index: 0,
		next:  -1,
	}
}

func (s *scene) Components(componentType ComponentType) ComponentView {
	if pool, ok := s.cPools[componentType]; ok {
		return &componentView{
			pool:  pool,
			index: 0,
			next:  -1,
		}
	}

	return emptyComponentView{}
}

func (s *scene) entityIdIsValid(entityID EntityID) bool {
	return s.entities[entityID.index].id.version == entityID.version
}
