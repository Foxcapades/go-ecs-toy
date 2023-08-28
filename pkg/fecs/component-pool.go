package fecs

import "github.com/Foxcapades/go-ecs-toy/pkg/fecs/futil"

const initialComponentPoolSize = 32

type componentRef struct {
	component Component
	version   componentVersion
}

// newComponentPool creates a new ComponentPool instance.
func newComponentPool() *componentPool {
	return &componentPool{
		unused: futil.NewStack[ComponentID](),
		data:   make([]componentRef, initialComponentPoolSize),
	}
}

type componentPool struct {
	unused futil.Stack[ComponentID]
	data   []componentRef
}

func (p *componentPool) toSlice() []Component {
	out := make([]Component, 0, p.size())
	for i := range p.data {
		if p.data[i].component != nil {
			out = append(out, p.data[i].component)
		}
	}
	return out
}

func (p *componentPool) size() int {
	return len(p.data) - p.unused.Size()
}

func (p *componentPool) get(id ComponentID) (Component, bool) {
	if !p.componentIdIsValid(id) {
		return nil, false
	}

	return p.data[id.index].component, true
}

func (p *componentPool) add(component Component) ComponentID {
	if !p.unused.IsEmpty() {
		oldId := p.unused.Pop()
		newId := newComponentID(oldId.index, oldId.version+1)

		p.data[newId.index].component = component
		p.data[newId.index].version = newId.version

		return newId
	}

	l := len(p.data)
	i := newComponentID(componentIndex(l), 1)

	p.data = append(p.data, componentRef{component, i.version})

	return i
}

func (p *componentPool) remove(id ComponentID) bool {
	if p.componentIdIsValid(id) {
		p.data[id.index].component = nil
		p.unused.Push(id)
		return true
	}

	return false
}

func (p *componentPool) componentIdIsValid(id ComponentID) bool {
	return id.index < componentIndex(len(p.data)) &&
		p.data[id.index].component != nil &&
		p.data[id.index].version == id.version
}
