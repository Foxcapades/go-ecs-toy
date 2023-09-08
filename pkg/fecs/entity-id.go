package fecs

import (
	"fmt"
	"regexp"

	"github.com/Foxcapades/go-ecs-toy/pkg/fecs/futil"
)

var entityIdRegex = regexp.MustCompile(`^eid-([0-9a-fA-F]{1,8})-([0-9a-fA-F]{1,8})-([0-9a-fA-F]{1,8})$`)

// ParseEntityID parses the given stringified EntityID value into a real
// EntityID instance.
//
// EntityID string values resemble "eid-ffff-ffff-ffff", where each group of 'f'
// characters in the example could be any hex string up to 8 characters in
// length.
func ParseEntityID(idString string) (EntityID, error) {
	matches := entityIdRegex.FindStringSubmatch(idString)

	if len(matches) != 4 {
		return EntityID{}, fmt.Errorf("invalid entity id string: %s", idString)
	}

	return EntityID{
		scene:   futil.MustParseHexUint32(matches[1]),
		index:   futil.MustParseHexUint32(matches[2]),
		version: futil.MustParseHexUint32(matches[3]),
	}, nil
}

type EntityID struct {
	// scene holds a reference to the id of the represented entity's parent Scene.
	scene uint32

	// index holds a reference to the position of the represented entity in the
	// parent/source entityPool.
	index uint32

	// version differentiates between entity instance reuses.  Each time an entity
	// instance is (re)used, this version value will be incremented.
	version uint32
}

func (e *EntityID) birth(sid SceneID, idx uint32) {
	if e.scene != 0 {
		panic("attempted to 'create' an already existing entity")
	}

	e.scene = sid
	e.index = idx
	e.version++
}

func (e *EntityID) resurrect(idx uint32) {
	e.index = idx
	e.version++
}

func (e *EntityID) isLiving(idx uint32) bool {
	return e.index == idx
}

func (e *entity) is(id *EntityID) bool {
	return e.id.Equals(id)
}

func (e *EntityID) kill() {
	// Flip the index bits.  We do this to differentiate between "live" and "dead"
	// entity instances.  A live instance will have an internal index that matches
	// its external index (in the parent entityPool), a dead instance will have an
	// index that does not match its external index.
	e.index = ^e.index
}

func (e *EntityID) Equals(other *EntityID) bool {
	return e.scene == other.scene &&
		e.index == other.index &&
		e.version == other.version
}

func (e *EntityID) String() string {
	return fmt.Sprintf("eid-%x-%x-%x", e.scene, e.index, e.version)
}
