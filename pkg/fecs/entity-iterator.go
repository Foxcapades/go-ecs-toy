package fecs

type entityIterator struct {
	pool  []entity
	index uint32
	next  *entity
}

func (e *entityIterator) HasNext() bool {
	// If the next pointer is not nil, then we have already found our next value
	// and we can bail here.  This happens in the case where HasNext is called
	// more than once between calls to Next.
	if e.next != nil {
		return true
	}

	// Iterate through the pool until we either hit the end of the pool slice or
	// we hit the end of the living entity block.  (The living entity block may
	// contain dead entities).
	for ; e.index < uint32(len(e.pool)) && e.pool[e.index].id.version > 0; e.index++ {
		// If the entity at the current index is living, then we can use it.
		if e.pool[e.index].isLiving(e.index) {
			// Set our next value to a reference to this entity.
			e.next = &e.pool[e.index]

			// Increment the index so next time HasNext is called it will start on the
			// next entity in the list.
			e.index++

			// Return true indicating we found an entity.
			return true
		}
	}

	// We didn't find another living entity in the pool.
	return false
}

func (e *entityIterator) Next() *entity {
	if !e.HasNext() {
		panic("no such element")
	}

	ref := e.next
	e.next = nil
	return ref
}
