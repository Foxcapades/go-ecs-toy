package fecs

import (
	"fmt"
	"strings"
)

// A ComponentMask is a bitmask for tracking what ComponentTypes are attached to
// a Component parent, such as an entity.
type ComponentMask struct {
	mask [4]uint64
}

// Add adds the given ComponentType to the bitmask.
func (c *ComponentMask) Add(cType ComponentType) {
	switch true {
	case cType > 192:
		c.mask[3] |= 1 << (cType - 193)
	case cType > 128:
		c.mask[2] |= 1 << (cType - 129)
	case cType > 64:
		c.mask[1] |= 1 << (cType - 65)
	default:
		c.mask[0] |= 1 << (cType - 1)
	}
}

// Has tests whether the bitmask has the given ComponentType.
func (c *ComponentMask) Has(cType ComponentType) bool {
	switch true {
	case cType > 192:
		return c.mask[3]&(1<<(cType-193)) > 0
	case cType > 128:
		return c.mask[2]&(1<<(cType-129)) > 0
	case cType > 64:
		return c.mask[1]&(1<<(cType-65)) > 0
	default:
		return c.mask[0]&(1<<(cType-1)) > 0
	}
}

// Remove removes the given ComponentType from the bitmask.
func (c *ComponentMask) Remove(cType ComponentType) {
	switch true {
	case cType > 192:
		c.mask[3] &= ^(1 << (cType - 193))
	case cType > 128:
		c.mask[2] &= ^(1 << (cType - 129))
	case cType > 64:
		c.mask[1] &= ^(1 << (cType - 65))
	default:
		c.mask[0] &= ^(1 << (cType - 1))
	}
}

// Clear clears the bitmask, removing all ComponentType values.
func (c *ComponentMask) Clear() {
	c.mask = [4]uint64{}
}

// String returns a string representation of the bitmask value.
func (c ComponentMask) String() string {
	b := new(strings.Builder)
	b.Grow(200)
	b.Reset()

	b.WriteByte('[')
	b.WriteString(fmt.Sprintf("%x ", c.mask[0]))
	b.WriteString(fmt.Sprintf("%x ", c.mask[1]))
	b.WriteString(fmt.Sprintf("%x ", c.mask[2]))
	b.WriteString(fmt.Sprintf("%x", c.mask[3]))
	b.WriteByte(']')

	return b.String()
}
