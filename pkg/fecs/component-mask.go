package fecs

import "strconv"

type componentMask struct {
	value [4]uint64
}

func (c *componentMask) add(cType ComponentType) {
	if cType > 192 {
		c.value[3] |= cType.toBitMask()
	} else if cType > 128 {
		c.value[2] |= cType.toBitMask()
	} else if cType > 64 {
		c.value[1] |= cType.toBitMask()
	} else {
		c.value[0] |= cType.toBitMask()
	}
}

func (c *componentMask) has(cType ComponentType) bool {
	tbm := cType.toBitMask()

	if cType > 192 {
		return c.value[3]&tbm == tbm
	} else if cType > 128 {
		return c.value[2]&tbm == tbm
	} else if cType > 64 {
		return c.value[1]&tbm == tbm
	} else {
		return c.value[0]&tbm == tbm
	}
}

func (c *componentMask) hasAll(other *componentMask) bool {
	return c.value[3]&other.value[3] == other.value[3] &&
		c.value[2]&other.value[2] == other.value[2] &&
		c.value[1]&other.value[1] == other.value[1] &&
		c.value[0]&other.value[0] == other.value[0]
}

func (c *componentMask) remove(cType ComponentType) {
	if cType > 192 {
		c.value[3] &= ^cType.toBitMask()
	} else if cType > 128 {
		c.value[2] &= ^cType.toBitMask()
	} else if cType > 64 {
		c.value[1] &= ^cType.toBitMask()
	} else {
		c.value[0] &= ^cType.toBitMask()
	}
}

func (c *componentMask) clear() {
	c.value = [4]uint64{}
}

func (c *componentMask) String() string {
	return "cm-" +
		strconv.FormatUint(c.value[3], 16) +
		"-" +
		strconv.FormatUint(c.value[2], 16) +
		"-" +
		strconv.FormatUint(c.value[1], 16) +
		"-" +
		strconv.FormatUint(c.value[0], 16)
}
