package fecs

import "github.com/Foxcapades/go-ecs-toy/pkg/fecs/futil"

type Scene interface {
	NewEntity() EntityID
	Assign(EntityID, Component) ComponentID
	Remove(EntityID)
	View(componentType ComponentType) SceneView
}

type scene struct {
	freeEntities futil.Stack[entityIndex]
	entities     []entityDesc
	cPools       map[ComponentType]ComponentPool
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

	s.entities[idx].mask.Reset()
	s.freeEntities.Push(idx)
}

func (s *scene) Assign(id EntityID, component Component) ComponentID {
	s.entities[id.index].mask.Set(component.GetType())

	if pool, ok := s.cPools[component.GetType()]; ok {
		return pool.Add(component)
	}

	pool := NewComponentPool()
	s.cPools[component.GetType()] = pool
	return pool.Add(component)
}

func (s *scene) View(componentType ComponentType) SceneView {
	return &sceneView{
		scene: s,
		cType: componentType,
		index: 0,
		next:  -1,
	}
}
