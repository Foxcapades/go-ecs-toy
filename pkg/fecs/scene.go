package fecs

import "github.com/Foxcapades/go-ecs-toy/pkg/fecs/futil"

type Scene interface {
	NewEntity() EntityID
	Assign(EntityID, Component) ComponentID
	Remove(EntityID)
	View(componentType ComponentType) EntityView
	Components(componentType ComponentType) ComponentView
}

type scene struct {
	freeEntities futil.Stack[entityIndex]
	entities     []entityDesc
	cPools       map[ComponentType]*componentPool
}

// NewEntity creates a new entity attached to the current scene.
//
// Returns the ID of the newly generated entity.
func (s *scene) NewEntity() (id EntityID) {

	// If we have an unused entity already in the slice, reuse it, else append a
	// new one.
	if !s.freeEntities.IsEmpty() {
		index := s.freeEntities.Pop()
		version := s.entities[index].id.version + 1

		id = createEntityID(index, version)
		s.entities[index].id = id

	} else {
		index := entityIndex(len(s.entities))
		version := entityVersion(1)

		id = createEntityID(index, version)
		s.entities = append(s.entities, entityDesc{id: id})
	}

	return
}

func (s *scene) Remove(id EntityID) {
	idx := id.index

	if s.entities[idx].id != id {
		return
	}

	s.entities[idx].mask.reset()
	s.freeEntities.Push(idx)
}

func (s *scene) Assign(id EntityID, component Component) ComponentID {
	s.entities[id.index].mask.set(component.GetType())

	if pool, ok := s.cPools[component.GetType()]; ok {
		return pool.add(component)
	}

	pool := newComponentPool()
	s.cPools[component.GetType()] = pool
	return pool.add(component)
}

func (s *scene) View(componentType ComponentType) EntityView {
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
