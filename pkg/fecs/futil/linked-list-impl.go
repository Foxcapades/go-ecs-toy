package futil

// NewLinkedList creates a new LinkedList instance of the target generic type T.
func NewLinkedList[T comparable]() LinkedList[T] {
	return new(linkedList[T])
}

type llNode[T comparable] struct {
	value T
	prev  *llNode[T]
	next  *llNode[T]
}

type linkedList[T comparable] struct {
	size int
	head *llNode[T]
	tail *llNode[T]
}

func (l *linkedList[T]) Size() int {
	return l.size
}

func (l *linkedList[T]) IsEmpty() bool {
	return l.size == 0
}

func (l *linkedList[T]) IsNotEmpty() bool {
	return l.size > 0
}

func (l *linkedList[T]) PushFront(value T) {
	if l.size == 0 {
		tmp := llNode[T]{value: value}
		l.head = &tmp
		l.tail = &tmp
	} else {
		tmp := llNode[T]{value, nil, l.head}
		l.head.prev = &tmp
		l.head = &tmp
	}

	l.size++
}

func (l *linkedList[T]) PushBack(value T) {
	if l.size == 0 {
		tmp := llNode[T]{value: value}
		l.head = &tmp
		l.tail = &tmp
	} else {
		tmp := llNode[T]{value, l.tail, nil}
		l.tail.next = &tmp
		l.tail = &tmp
	}

	l.size++
}

func (l *linkedList[T]) PopFront() T {
	if l.size == 0 {
		panic("no such element")
	}

	o := l.head

	if l.size == 1 {
		l.head = nil
		l.tail = nil
	} else {
		l.head = o.next
		l.head.prev = nil
	}

	l.size--

	return o.value
}

func (l *linkedList[T]) PopBack() T {
	if l.size == 0 {
		panic("no such element")
	}

	o := l.tail

	if l.size == 1 {
		l.head = nil
		l.tail = nil
	} else {
		l.tail = o.prev
		l.tail.next = nil
	}

	l.size--

	return o.value
}

func (l *linkedList[T]) PeekFront() T {
	if l.size == 0 {
		panic("no such element")
	}

	return l.head.value
}

func (l *linkedList[T]) PeekBack() T {
	if l.size == 0 {
		panic("no such element")
	}

	return l.tail.value
}

func (l *linkedList[T]) Get(index int) T {
	return l.getByIndex(index).value
}

func (l *linkedList[T]) RemoveIndex(index int) T {
	node := l.getByIndex(index)

	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}

	if node.next == nil {
		l.tail = node.prev
	}
	if node.prev == nil {
		l.head = node.next
	}

	l.size--

	return node.value
}

func (l *linkedList[T]) RemoveValue(value T) bool {
	node := l.getByValue(value)

	if node == nil {
		return false
	}

	l.unlink(node)

	if node.next == nil {
		l.tail = node.prev
	}
	if node.prev == nil {
		l.head = node.next
	}

	l.size--

	return true
}

func (l *linkedList[T]) RemoveFirstMatching(predicate func(T) bool) bool {
	cur := l.head

	for cur != nil {
		if predicate(cur.value) {
			l.unlink(cur)
			l.size--
			return true
		}

		cur = cur.next
	}

	return false
}

func (l *linkedList[T]) RemoveAllMatching(predicate func(T) bool) int {
	cur := l.head
	out := 0

	for cur != nil {
		if predicate(cur.value) {
			l.unlink(cur)
			l.size--
			out++
		}
	}

	return out
}

func (l *linkedList[T]) Clear() {
	l.head = nil
	l.tail = nil
	l.size = 0
}

func (l *linkedList[T]) Iterator() Iterator[T] {
	return &linkedListIterator[T]{l.head}
}

func (l *linkedList[T]) Has(value T) bool {
	cur := l.head

	for cur != nil {
		if cur.value == value {
			return true
		}

		cur = cur.next
	}

	return false
}

func (l *linkedList[T]) HasAny(fn func(T) bool) bool {
	cur := l.head

	for cur != nil {
		if fn(cur.value) {
			return true
		}

		cur = cur.next
	}

	return false
}

func (l *linkedList[T]) FindFirst(fn func(T) bool) (out T, found bool) {
	cur := l.head

	for cur != nil {
		if fn(cur.value) {
			out = cur.value
			found = true
			return
		}

		cur = cur.next
	}

	return
}

func (l *linkedList[T]) FindLast(fn func(T) bool) (out T, found bool) {
	cur := l.tail

	for cur != nil {
		if fn(cur.value) {
			out = cur.value
			found = true
			return
		}

		cur = cur.prev
	}

	return
}

func (l *linkedList[T]) ForEach(fn func(T)) {
	cur := l.head

	for cur != nil {
		fn(cur.value)
	}
}

func (l *linkedList[T]) getByIndex(index int) *llNode[T] {
	if index >= l.size {
		panic("no such element")
	}

	if l.size == 1 || index == 0 {
		return l.head
	}

	if index == l.size-1 {
		return l.tail
	}

	if index < l.size/2+1 {
		c := l.head

		for i := 0; i < l.size; i++ {
			if i == index {
				return c
			}

			c = c.next
		}
	} else {
		c := l.tail

		for i := l.size - 1; i >= 0; i-- {
			if i == index {
				return c
			}

			c = c.prev
		}
	}

	panic("illegal state")
}

func (l *linkedList[T]) getByValue(value T) *llNode[T] {
	if l.size == 0 {
		return nil
	}

	c := l.head

	for i := 0; i < l.size; i++ {
		if c.value == value {
			return c
		}

		c = c.next
	}

	return nil
}

func (l *linkedList[T]) unlink(node *llNode[T]) {
	if node != nil {
		if node.prev != nil {
			node.prev.next = node.next
		}
		if node.next != nil {
			node.next.prev = node.prev
		}
	}
}

type linkedListIterator[T comparable] struct {
	node *llNode[T]
}

func (l linkedListIterator[T]) HasNext() bool {
	return l.node != nil
}

func (l linkedListIterator[T]) Next() T {
	out := l.node
	l.node = l.node.next
	return out.value
}
