package futil

type Numeric interface {
	uint8 | uint16 | uint32 | uint64 |
	int8 | int16 | int32 | int64 |
	float32 | float64
}

type Vec2[T Numeric] [2]T
type Vec3[T Numeric] [3]T
type Vec4[T Numeric] [4]T
