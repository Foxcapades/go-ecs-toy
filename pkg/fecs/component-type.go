package fecs

import "strconv"

const MaxComponentTypes uint8 = 255

var currentComponentTypeId uint8 = 0

func NewComponentType() ComponentType {
	if currentComponentTypeId == 255 {
		panic("attempted to register more than 255 component types")
	}

	currentComponentTypeId++
	return ComponentType(currentComponentTypeId)
}

type ComponentType uint8

func (c ComponentType) String() string {
	return "ct-" + strconv.FormatUint(uint64(c), 16)
}

func (c ComponentType) toBitMask() uint64 {
	if c > 192 {
		return 1 << (c - 193)
	} else if c > 128 {
		return 1 << (c - 129)
	} else if c > 64 {
		return 1 << (c - 65)
	} else {
		return 1 << (c - 1)
	}
}
