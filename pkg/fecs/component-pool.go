package fecs

import "github.com/Foxcapades/go-ecs-toy/pkg/fecs/futil"

const (
	componentPoolScaleFactor     float32 = 1.5
	componentPoolInitialCapacity         = 32
)

func newComponentPool() *componentPool {
	return &componentPool{
		free: futil.NewStack[uint32](),
		ids:  make([]ComponentID, componentPoolInitialCapacity),
		pool: make([]Component, componentPoolInitialCapacity),
		size: 0,
	}
}

type componentPool struct {
	free futil.Stack[uint32]
	ids  []ComponentID
	pool []Component
	size uint32
}

func (c *componentPool) containsComponent(id *ComponentID) bool {
	return id.index < c.size && c.ids[id.index].Equals(id)
}

func (c *componentPool) getComponent(id *ComponentID) (Component, bool) {
	if !c.containsComponent(id) {
		return nil, false
	}

	return c.pool[id.index], true
}

func (c *componentPool) newComponent(comp Component) ComponentID {
	if c.free.IsEmpty() {
		return c._append(comp)
	} else {
		return c._overwrite(comp)
	}
}

func (c *componentPool) removeComponent(id *ComponentID) bool {
	if !c.containsComponent(id) {
		return false
	}

	idRef := &c.ids[id.index]
	idRef.clear()

	c.pool[id.index] = nil
	c.free.Push(id.index)

	return true
}

func (c *componentPool) _append(comp Component) ComponentID {
	c.ids[c.size].init(c.size, comp.Type())
	c.pool[c.size] = comp
	c.size++
	return c.ids[c.size-1]
}

func (c *componentPool) _overwrite(comp Component) ComponentID {
	c._ensureCapacity(c.size + 1)
	idx := c.free.Pop()

	c.ids[idx].init(idx, comp.Type())
	c.pool[idx] = comp

	return c.ids[idx]
}

func (c *componentPool) _ensureCapacity(minimum uint32) {
	if minimum <= uint32(len(c.pool)) {
		return
	}

	newSize := max(uint32(float32(len(c.pool))*componentPoolScaleFactor), minimum)

	newIDs := make([]ComponentID, newSize)
	copy(newIDs, c.ids)
	c.ids = newIDs

	newComponents := make([]Component, newSize)
	copy(newComponents, c.pool)
	c.pool = newComponents
}
