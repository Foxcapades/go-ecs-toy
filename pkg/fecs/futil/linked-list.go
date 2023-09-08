package futil

type LinkedList[T comparable] interface {

	// Size returns the count of elements currently in this LinkedList.
	Size() int

	// IsEmpty returns whether this LinkedList is empty.
	IsEmpty() bool

	// IsNotEmpty returns whether this LinkedList contains one or more items.
	IsNotEmpty() bool

	// PushFront pushes the given value onto the front of this LinkedList.
	PushFront(value T)

	// PushBack pushes the given value onto the back of this LinkedList.
	PushBack(value T)

	// PopFront removes and returns the value at the front of this LinkedList.
	//
	// If this LinkedList is empty, this method will panic.
	PopFront() T

	// PopBack removes and returns the value at the back of this LinkedList.
	//
	// If this LinkedList is empty, this method will panic.
	PopBack() T

	// PeekFront returns the value at the front of this LinkedList without
	// removing it.
	//
	// If this LinkedList is empty, this method will panic.
	PeekFront() T

	// PeekBack returns the value at the back of this LinkedList without removing
	// it.
	//
	// If this LinkedList is empty, this method will panic.
	PeekBack() T

	// Get fetches the value at the given index in this LinkedList.
	//
	// NOTE: This list is not indexed, meaning that this method may have to
	// iterate through up to half the size of the list to retrieve the value.
	//
	// If the given index is less than 0 or is greater than Size, this method will
	// panic.
	//
	// This method has the worst-case time signature of O(n), where `n` is the
	// Size of this LinkedList.
	//
	// Returns the value at the given index.
	Get(index int) T

	// Has tests whether the given value is present in this LinkedList instance.
	//
	// This method has the worst-case time signature of O(n), where `n` is the
	// Size of this LinkedList.
	//
	// Returns a boolean value indicating whether the target value was found in
	// this LinkedList.
	Has(value T) bool

	// HasAny tests whether any of the elements in this LinkedList match the given
	// predicate.
	HasAny(predicate func(T) bool) bool

	// FindFirst attempts to find the first instance of a value that matches the
	// given predicate.
	//
	// This method has the worst-case time signature of O(n), where `n` is the
	// Size of this LinkedList.
	//
	// Returns a value of type T and a boolean value indicating whether a matching
	// value was found.
	FindFirst(predicate func(T) bool) (T, bool)

	// FindLast attempts to find the last instance of a value that matches the
	// given predicate.
	//
	// This method has the worst-case time signature of O(n), where `n` is the
	// Size of this LinkedList.
	//
	// Returns a value of type T and a boolean value indicating whether a matching
	// value was found.
	FindLast(predicate func(T) bool) (T, bool)

	// ForEach calls the given function on each element in this LinkedList.
	ForEach(fun func(T))

	// RemoveIndex removes the value at the given index from this LinkedList,
	// returning that value.
	//
	// If the given index is less than zero or is greater than Size, this method
	// will panic.
	RemoveIndex(index int) T

	// RemoveValue removes the target value from this LinkedList.
	//
	// NOTE: This method may have to iterate up to Size entries to find the target
	// value for removal.
	//
	// Returns a boolean indicator whether the target was found in the LinkedList.
	RemoveValue(value T) bool

	RemoveFirstMatching(predicate func(T) bool) bool

	RemoveAllMatching(predicate func(T) bool) int

	// Clear removes all items from this LinkedList.
	Clear()

	// Iterator returns a new Iterator instance over the values currently in this
	// LinkedList.
	Iterator() Iterator[T]
}
