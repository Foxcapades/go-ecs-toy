package fecs

import (
	"fmt"
	"strconv"

	"github.com/Foxcapades/go-ecs-toy/pkg/fecs/futil"
)

var sceneID SceneID = 0

func NewScene() Scene {
	sceneID++
	return &scene{
		sceneID:    sceneID,
		entities:   newEntityPool(),
		components: make(map[ComponentType]*componentPool, 16),
	}
}

type scene struct {
	sceneID    SceneID
	entities   entityPool
	components map[ComponentType]*componentPool
}

func (s *scene) ID() SceneID {
	return s.sceneID
}

func (s *scene) ContainsEntity(id *EntityID) bool {
	return s.entities.containsEntity(id)
}

func (s *scene) DestroyEntity(id *EntityID) bool {
	// If the target entity is not in this scene, return false as we aren't
	// removing it.
	if !s.entities.containsEntity(id) {
		return false
	}

	// Grab all the component ids attached to the target entity.
	comps := s.entities.getEntityComponents(id)

	// Remove the entity's components from the component pools.
	for _, ref := range comps {
		if pool, ok := s.components[ref.ctype]; ok {
			pool.removeComponent(ref)
		}
	}

	// kill the entity.
	s.entities.removeEntity(id)

	// return true because we did remove the entity from the scene.
	return true
}

func (s *scene) Entities(ct ...ComponentType) futil.Iterator[EntityID] {
	return s.entities.entities(ct)
}

func (s *scene) NewEntity() EntityID {
	return s.entities.newEntity(s.sceneID)
}

func (s *scene) AttachComponent(id *EntityID, constructor ComponentConstructor) ComponentID {
	if !s.entities.containsEntity(id) {
		panic(fmt.Errorf("attempted to attach a new component to an entity (%s) which is not currently registered to the target scene (%s)", id.String(), s.String()))
	}

	comp := constructor()

	if pool, ok := s.components[comp.Type()]; ok {
		cid := pool.newComponent(comp)
		s.entities.addComponent(id, &cid)
		return cid
	}

	pool := newComponentPool()
	s.components[comp.Type()] = pool

	cid := pool.newComponent(comp)
	s.entities.addComponent(id, &cid)

	return cid
}

func (s *scene) GetComponent(cid *ComponentID) (Component, bool) {
	if pool, ok := s.components[cid.ctype]; ok {
		return pool.getComponent(cid)
	}

	return nil, false
}

func (s *scene) GetComponentByType(eid *EntityID, ct ComponentType) (Component, bool) {
	if id, ok := s.entities.getEntityComponent(eid, ct); ok {
		if pool, ok := s.components[ct]; ok {
			if comp, ok := pool.getComponent(id); ok {
				return comp, true
			}

			panic("illegal state")
		}

		panic("illegal state")
	}

	return nil, false
}

func (s *scene) HasComponent(eid *EntityID, cid *ComponentID) bool {
	// TODO implement me
	panic("implement me")
}

func (s *scene) RemoveComponent(eid *EntityID, cid *ComponentID) bool {
	// TODO implement me
	panic("implement me")
}

func (s *scene) String() string {
	return "scene-" + strconv.FormatUint(uint64(s.sceneID), 16)
}
