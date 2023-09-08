package fecs

import (
	"fmt"

	"github.com/Foxcapades/go-ecs-toy/pkg/fecs/futil"
)

type SceneID = uint32

type Scene interface {
	fmt.Stringer

	// ID returns the identifier for this Scene instance.
	ID() SceneID

	// ContainsEntity tests whether this Scene currently contains the entity
	// identified by the given EntityID.
	ContainsEntity(id *EntityID) bool

	// DestroyEntity removes the target entity from the Scene and unlinks all the
	// Component instances attached to it.
	DestroyEntity(id *EntityID) bool

	// Entities returns an Iterator over all the entities in this Scene that have
	// all Components of all the given ComponentTypes attached to them.
	//
	// If no ComponentType values are passed to this method, the returned Iterator
	// will iterate over all the entities in this Scene.
	//
	// If ComponentType values are provided, each entity returned by the Iterator
	// this method returns will have Components of all the target types attached.
	//
	// Adding or removing entities from this Scene while an Iterator is in use may
	// cause undefined behavior.
	Entities(componentTypes ...ComponentType) futil.Iterator[EntityID]

	// NewEntity creates a new entity in this Scene and returns its EntityID.
	//
	// Entities themselves consist of the returned EntityID and a mask of attached
	// ComponentTypes.
	NewEntity() EntityID

	// AttachComponent attaches a new Component created by the given constructor
	// to the entity identified by the given EntityID.
	//
	// If the target entity is not found in this Scene, this method will panic.
	// ContainsEntity should be called before this method to ensure that this
	// method will succeed.
	//
	// This method returns the ComponentID generated for the new Component
	// created by the given ComponentConstructor.
	AttachComponent(id *EntityID, constructor ComponentConstructor) ComponentID

	// GetComponent attempts to look up a Component identified by the given
	// ComponentID from this Scene.
	//
	// If the target Component could not be found in this Scene, this method will
	// return nil and false.
	//
	// If the target Component was found in this Scene, this method will return
	// the target Component and true.
	GetComponent(cid *ComponentID) (Component, bool)

	// GetComponentByType attempts to look up a Component with the given
	// ComponentType attached to a target entity identified by the given EntityID.
	//
	// If the target entity does not have a Component of the given ComponentType,
	// this method will return nil and false.
	//
	// If the target entity does have a Component of the given ComponentType, this
	// method will return the located Component and true.
	GetComponentByType(eid *EntityID, ct ComponentType) (Component, bool)

	// HasComponent tests whether the entity identified by the given EntityID has
	// the target Component attached to it.
	HasComponent(eid *EntityID, cid *ComponentID) bool

	// RemoveComponent attempts to remove the target component from the entity
	// identified by the given EntityID.
	//
	// Returns a boolean value indicating whether the target Component was
	// attached to the target entity before this method was called.
	RemoveComponent(eid *EntityID, cid *ComponentID) bool
}
