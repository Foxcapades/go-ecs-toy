package fecs

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
