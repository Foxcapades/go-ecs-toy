package fecs

type EntityView interface {
	HasNext() bool
	Next() EntityID
}

type entityFilter = func(*entity) bool

func newEntityView(entities []entity, filter entityFilter) EntityView {
	return &entityView{
		entities: entities,
		filter:   filter,
		next:     -1,
		index:    0,
	}
}

type entityView struct {
	entities []entity
	filter   entityFilter
	next     int
	index    int
}

func (e *entityView) HasNext() bool {
	if e.next != -1 {
		return true
	}

	for ; e.index < len(e.entities); e.index++ {
		if e.filter(&e.entities[e.index]) {
			e.next = e.index
			e.index++
			return true
		}
	}

	return false
}

func (e *entityView) Next() (out EntityID) {
	if !e.HasNext() {
		panic("no such element")
	}

	out = e.entities[e.next].id
	e.next = -1

	return
}

func newEmptyEntityView() EntityView {
	return emptyEntityView{}
}

type emptyEntityView struct{}

func (e emptyEntityView) HasNext() bool {
	return false
}

func (e emptyEntityView) Next() EntityID {
	panic("no such element")
}
