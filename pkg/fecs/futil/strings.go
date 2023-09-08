package futil

import "strconv"

func MustParseHexUint32(value string) uint32 {
	if out, err := strconv.ParseUint(value, 16, 32); err != nil {
		panic(err)
	} else {
		return uint32(out)
	}
}
