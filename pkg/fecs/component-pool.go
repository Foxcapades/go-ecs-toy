package fecs

import "github.com/Foxcapades/go-ecs-toy/pkg/fecs/futil"

const initialComponentPoolSize = 32

type componentRef struct {
	component Component
	version   ComponentVersion
}

// NewComponentPool creates a new ComponentPool instance.
func NewComponentPool() ComponentPool {
	return &componentPool{
		unused: futil.NewStack[ComponentID](),
		data:   make([]componentRef, initialComponentPoolSize),
	}
}

// ComponentPool defines a block of contiguous memory containing zero or more
// registered components.
type ComponentPool interface {
	// Size returns the current number of components in the component pool.
	Size() (poolSize int)

	// Get returns the component with the given component ID.
	//
	// If the requested component does not exist in the target ComponentPool, the
	// return values will be nil and false.  If the requested component DOES exist
	// in the target ComponentPool, the return values will be the target component
	// and true.
	Get(id ComponentID) (component Component, ok bool)

	// Add adds the given Component to the target ComponentPool.
	//
	// WARNING! No verification is performed to ensure that the given Component is
	// unique in the target pool.  Adding the same Component to the target
	// ComponentPool more than once may result in undefined behavior.
	Add(component Component) (componentID ComponentID)

	// Remove removes the target component from the ComponentPool.
	//
	// Returns true if the component was removed, returns false if the target
	// component was not found in the ComponentPool.
	Remove(id ComponentID) bool

	// ForEach calls the given function on each Component in the ComponentPool
	// in registration order.
	ForEach(func(component Component))
}

type componentPool struct {
	unused futil.Stack[ComponentID]
	data   []componentRef
}

func (p *componentPool) Size() int {
	return len(p.data) - p.unused.Size()
}

func (p *componentPool) Get(id ComponentID) (Component, bool) {
	if !p.componentIdIsValid(id) {
		return nil, false
	}

	return p.data[id.index].component, true
}

func (p *componentPool) Add(component Component) ComponentID {
	if !p.unused.IsEmpty() {
		oldId := p.unused.Pop()
		newId := NewComponentID(oldId.index, oldId.version+1)

		p.data[newId.index].component = component
		p.data[newId.index].version = newId.version

		return newId
	}

	l := len(p.data)
	i := NewComponentID(componentIndex(l), 1)

	p.data = append(p.data, componentRef{component, i.version})

	return i
}

func (p *componentPool) Remove(id ComponentID) bool {
	if p.componentIdIsValid(id) {
		p.data[id.index].component = nil
		p.unused.Push(id)
		return true
	}

	return false
}

func (p *componentPool) ForEach(fn func(component Component)) {
	for i := range p.data {
		if p.data[i].component != nil {
			fn(p.data[i].component)
		}
	}
}

func (p *componentPool) componentIdIsValid(id ComponentID) bool {
	return id.index < componentIndex(len(p.data)) &&
		p.data[id.index].component != nil &&
		p.data[id.index].version == id.version
}
