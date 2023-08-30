package fecs

import "github.com/Foxcapades/go-ecs-toy/pkg/fecs/futil"

// componentPoolItem contains a reference to a target Component and ComponentID.
//
// this struct is used by the componentPool implementation to track contained
// Component instances with any necessary additional metadata.
type componentPoolItem struct {
	id  ComponentID
	ref Component
}

// newComponentPool creates a new ComponentPool instance with a default
// Component capacity of 32.
func newComponentPool() ComponentPool {
	return &componentPool{
		free:  futil.NewStack[uint32](),
		pool:  make([]componentPoolItem, 32),
		index: 0,
	}
}

// componentPool implements the ComponentPool interface, hiding the
// implementation details from consuming code.
type componentPool struct {
	// free tracks unused slots in the pool slice of componentPoolItem instances.
	//
	// Unused slots happen when a Component is removed from the pool.  Rather than
	// shifting all the items around in the pool slice, we track the indices of
	// removed componentPoolItems for reuse when a new Component is added to the
	// ComponentPool later.
	free futil.Stack[uint32]

	// pool is the data container for the registered Component instances.
	pool []componentPoolItem

	// index is the position in the pool slice where new Components will be added
	// if there are no already unused indices tracked by free.
	index uint32
}

func (c *componentPool) Size() int {
	return len(c.pool) - c.free.Size()
}

func (c *componentPool) Get(componentID ComponentID) (Component, bool) {
	if c.isValidComponentID(componentID) {
		return c.pool[componentID.index].ref, true
	}

	return nil, false
}

func (c *componentPool) Has(componentID ComponentID) bool {
	return c.isValidComponentID(componentID)
}

func (c *componentPool) Add(component Component) ComponentID {
	if c.free.IsEmpty() {
		return c.append(component)
	}

	return c.overwrite(component)
}

func (c *componentPool) Remove(componentID ComponentID) bool {
	if !c.isValidComponentID(componentID) {
		return false
	}

	c.pool[componentID.index].id = newComponentID(0, componentID.version+1, componentID.cType)
	c.pool[componentID.index].ref = nil

	c.free.Push(componentID.index)

	return true
}

func (c *componentPool) ComponentView() ComponentView {
	if c.Size() == 0 {
		return emptyComponentView{}
	} else {
		return newComponentView(c.pool)
	}
}

// overwrite overwrites an existing, unused component slot with the given
// component, generating a new ComponentID handle to that Component.
func (c *componentPool) overwrite(component Component) ComponentID {
	in := c.free.Pop()
	id := newComponentID(in, c.pool[in].id.version, component.Type())

	c.pool[in].ref = component

	return id
}

// append appends the given Component to the end of this ComponentPool and
// generates a new ComponentID handle to that Component.
func (c *componentPool) append(component Component) ComponentID {
	c.ensureCapacity(c.index)

	id := newComponentID(c.index, 1, component.Type())

	c.pool[c.index].ref = component
	c.pool[c.index].id = id

	c.index++

	return id
}

// ensureCapacity ensures that the ComponentPool has enough space in memory to
// hold the given minimum number of elements.  If the ComponentPool does not
// have enough space when this method is called, the underlying data container
// will be replaced with a new container of a greater size.
//
// The argument minimum defines the minimum size that the ComponentPool must be,
// however if a resize occurs, the ComponentPool may grow to a capacity that is
// larger than the given minimum.
func (c *componentPool) ensureCapacity(minimum uint32) {

	// If the ComponentPool already has a capacity greater than or equal to the
	// given minimum, then we don't need to resize.
	if uint32(len(c.pool)) >= minimum {
		return
	}

	// Calculate a possible new size for the ComponentPool.
	newSize := uint32(float32(len(c.pool))*1.5) + 1

	// If the calculated size is not big enough for the given minimum, use that
	// value instead.
	if newSize < minimum {
		newSize = minimum
	}

	// Create a new container of the target new container size.
	tmp := make([]componentPoolItem, newSize)

	// Copy over the elements from the original container.
	copy(tmp, c.pool)

	// Replace the current data container with the new container.  GC will handle
	// removing the old data container.
	c.pool = tmp
}

// isValidComponentID verifies that the given ComponentID aligns with a
// Component currently in this ComponentPool.
//
// If no such Component could be found in this ComponentPool, this method
// returns false, otherwise, if the target ComponentID was found, returns true.
func (c *componentPool) isValidComponentID(i ComponentID) bool {
	return i.index < uint32(len(c.pool)) &&
		c.pool[i.index].ref != nil &&
		c.pool[i.index].id.Equals(&i)
}
