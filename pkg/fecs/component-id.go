package fecs

import "fmt"

type ComponentID struct {
	index   uint32
	version uint32
	ctype   ComponentType
}

func (c *ComponentID) init(idx uint32, ct ComponentType) {
	if c.index != 0 {
		panic("attempted to initialize a component id more than once")
	}

	c.index = idx
	c.version++
	c.ctype = ct
}

func (c *ComponentID) isActive(idx uint32) bool {
	return c.index == idx
}

func (c *ComponentID) clear() {
	c.index = ^c.index
}

func (c *ComponentID) Equals(other *ComponentID) bool {
	return c.ctype == other.ctype &&
		c.index == other.index &&
		c.version == other.version
}

func (c *ComponentID) String() string {
	return fmt.Sprintf("cid-%x-%x-%x", c.ctype, c.index, c.version)
}
