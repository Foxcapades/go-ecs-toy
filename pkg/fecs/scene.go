package fecs

import "github.com/Foxcapades/go-ecs-toy/pkg/fecs/futil"

type Scene interface {
	// NewEntity generates a new entity and returns its ID.
	NewEntity() EntityID

	// AssignComponent assigns the given component to the target entity.
	AssignComponent(entityID EntityID, component Component) ComponentID

	// RemoveEntity unlinks the target entity from the scene.
	RemoveEntity(entityID EntityID)

	// RemoveComponent removes the Component with the given ComponentID from the
	// scene.
	//
	// WARNING!  This does not remove the component from any entities to which it
	// was attached.  Use RemoveComponentFromEntity to remove a component from
	// entities to which it was attached.
	RemoveComponent(componentID ComponentID)

	// RemoveComponentFromEntity unlinks the Component with the given ComponentID
	// from the target entity.
	//
	// WARNING!  This does not remove the component from the scene, it only
	// removes it from the target entity.  This may leave dangling Components if
	// they are not removed via RemoveComponent.
	RemoveComponentFromEntity(entityID EntityID, componentID ComponentID)

	// EntitiesWith returns an EntityView over all registered entities with a
	// Component of the given ComponentType attached.
	//
	// WARNING!  Under the hood this method iterates over every entity registered
	// with the scene.  Use this method sparingly.
	EntitiesWith(componentType ComponentType) EntityView

	// Components returns a ComponentView over all the registered components of
	// the given type.
	Components(componentType ComponentType) ComponentView
}

//

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
