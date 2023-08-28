package fecs

// SceneView defines an iterator over Scene entities that have a target
// ComponentType attached.
type SceneView interface {
	// HasNext returns a boolean flag indicating whether there exists at least one
	// more entity in the scene with the target ComponentType attached.
	HasNext() bool

	// Next returns the EntityID for the next available entity which has the
	// target ComponentType attached.
	Next() EntityID
}

type sceneView struct {
	scene *scene
	cType ComponentType
	index int
	next  int
}

func (s *sceneView) HasNext() bool {
	for ; s.index < len(s.scene.entities); s.index++ {
		if s.scene.entities[s.index].mask.Has(s.cType) {
			s.next = s.index
			s.index++
			return true
		}
	}

	s.next = -1
	return false
}

func (s *sceneView) Next() (out EntityID) {
	out = s.scene.entities[s.next].id
	s.next = -1
	return
}
