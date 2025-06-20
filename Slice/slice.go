// package Slice provides a custom Slice type with chainable methods
// with API similar to that of the standard library's slices package.
package Slice

import (
	"iter"
	"slices"
)

// type Slice is alias for custom Generic slice type
type Slice[T comparable] struct {
	Items []T
}

// New return a new instance of Slice type
func New[T comparable](items []T) Slice[T] {
	data := make([]T, len(items))
	copy(data, items)
	return Slice[T]{Items: data}
}

// Length return the length of underlying slice
func (self Slice[T]) Length() int {
	return len(self.Items)
}

// Items return the shallow copy of underlying slice
func (self Slice[T]) Clone() []T {
	return slices.Clone(self.Items)
}

// Push inserts a single element to end of self
func (self *Slice[T]) Push(elem T) {
	self.Items = append(self.Items, elem)
}

// Append appends a new single element to end of self
func (self *Slice[T]) Append(elems ...T) {
	self.Items = append(self.Items, elems...)
}

// AppendSeq appends the values from seq to the Slice
func (self *Slice[T]) AppendSeq(seq iter.Seq[T]) {
	self.Items = slices.AppendSeq(self.Items, seq)
}

// Concat appends every element in elems to end of self
func (self *Slice[T]) Concat(elems []T) {
	self.Items = append(self.Items, elems...)
}

// At returns the element at nth index of self.
// This is provided as a helper method for convenience, one can 
// directly use Slice{}.Items[n]
func (self Slice[T]) At(n int) T {
	return self.Items[n]
}

// All returns an iterator over index-value pairs in the 
// slice in the usual order.
func (self Slice[T]) All() iter.Seq2[int, T] {
	return slices.All(self.Items)
}

// Backward returns an iterator over index-value pairs in the Slice,
// traversing it backward with descending indices.
func (self Slice[T]) Backward() iter.Seq2[int, T] {
	return slices.Backward(self.Items)
}

// TODO: BinarySearch
// TODO: BinarySearchFunc
// TODO: Chunk
// -Clip
// -Collect
// TODO: Compact
// TODO: CompactFunc
// TODO: Compare
// TODO: CompareFunc

// Contains reports whether v is present in self.
func (self Slice[T]) Contains(v T) bool {
	return slices.Contains(self.Items, v)
}

// ContainsFunc reports whether at least one element e of Slice satisfies f(e).
func (self Slice[T]) ContainsFunc(f func(T) bool) bool {
	return slices.ContainsFunc(self.Items, f)
}

// Delete removes the elements self.Items[i:j] from self. Delete panics if j > self.Length()
// or self.Items[i:j] is not a valid slice of self. Delete is O(self.Length()-i), so if 
// many items must be deleted, it is better to make a single call deleting them all 
// together than to delete one at a time. Delete zeroes the elements self.Items[self.Length()-(j-i):self.Length()].
func (self *Slice[T]) Delete(i, j int) {
	self.Items = slices.Delete(self.Items, i, j)
}

// DeleteFunc removes any elements from self for which del returns true. 
// DeleteFunc zeroes the elements between the new length and the original length.
func (self *Slice[T]) DeleteFunc(del func(T) bool) {
	self.Items = slices.DeleteFunc(self.Items, del)
}

// Equal reports whether two slices are equal: the same length and all elements 
// equal. If the lengths are different, Equal returns false. Otherwise, the 
// elements are compared in increasing index order, and the comparison stops at 
// the first unequal pair. Empty and nil slices are considered equal. 
// Floating point NaNs are not considered equal.
func (self Slice[T]) Equal(other Slice[T]) bool {
	return slices.Equal(self.Items, other.Items)
}

// Index returns the index of the first occurrence of v in self, 
// or -1 if not present.
func (self Slice[T]) Index(v T) int {
	return slices.Index(self.Items, v)
}

// IndexFunc returns the first index i satisfying f(self.At(i)), or -1 if none do.
func (self Slice[T]) IndexFunc(f func(T) bool) int {
	return slices.IndexFunc(self.Items, f)
}

// Insert inserts the values v... into self at index i. The elements 
// at self.Items[i:] are shifted up to make room. In the modified Slice, 
// self.At(i) == v[0], and, if i < self.Length(), self.At(i+len(v)) == value 
// originally at self.At(i). Insert panics if i > self.Length(). This function 
// is O(self.Length() + len(v)).
func (self *Slice[T]) Insert(i int, v ...T) {
	self.Items = slices.Insert(self.Items, i, v...)
}

// Reverse reverses the elements of the slice in place.
func (self *Slice[T]) Reverse() {
	slices.Reverse(self.Items)
}

// Values returns an iterator that yields the slice elements in order.
func (self Slice[T]) Values() iter.Seq[T] {
	return slices.Values(self.Items)
}

// Map applies the given function to each element and returns a new Slice
// with the transformed elements of type U
func (self Slice[T]) Map(fn func(T) any) Slice[any] {
	return MapTo(self, fn)
}

// MapTo applies the given function to each element and returns a new Slice
// with the transformed elements of the specified type U.
// This is the type-safe version of [Map] method.
func MapTo[T, U comparable](self Slice[T], fn func(T) U) Slice[U] {
	result := make([]U, self.Length())
	for i, v := range self.Items {
		result[i] = fn(v)
	}
	return New(result)
}

