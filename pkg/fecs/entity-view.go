package fecs

// EntityView defines an iterator over Scene entities that have a target
// ComponentType attached.
type EntityView interface {
	// HasNext returns a boolean flag indicating whether there exists at least one
	// more entity in the scene with the target ComponentType attached.
	HasNext() bool

	// Next returns the EntityID for the next available entity which has the
	// target ComponentType attached.
	Next() EntityID
}

type entityView struct {
	scene *scene
	cType ComponentType
	index int
	next  int
}

func (s *entityView) HasNext() bool {
	for ; s.index < len(s.scene.entities); s.index++ {
		if s.scene.entities[s.index].mask.has(s.cType) {
			s.next = s.index
			return true
		}
	}

	s.next = -1
	return false
}

func (s *entityView) Next() (out EntityID) {
	if s.next == -1 && !s.HasNext() {
		panic("no such element")
	}

	out = s.scene.entities[s.next].id
	s.next = -1
	s.index++
	return
}
