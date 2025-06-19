// package Slice provides a custom Slice type with chainable methods
// with API similar to that of the standard library's slices package.
package Slice

import (
	"fmt"
	"strings"

	"github.com/harishtpj/klassy/String"
)

// type Slice is alias for custom Generic slice type
type Slice[T any] struct {
	data []T
}

// New return a new instance of Slice type
func New[T any](items []T) Slice[T] {
	data := make([]T, len(items))
	copy(data, items)
	return Slice[T]{data: data}
}

// Length return the length of underlying slice
func (self Slice[T]) Length() int {
	return len(self.data)
}

// Items return the copy of underlying slice
func (self Slice[T]) Items() []T {
	data := make([]T, self.Length())
	copy(data, self.data)
	return data
}

// Push inserts a single element to end of self
func (self *Slice[T]) Push(elem T) {
	self.data = append(self.data, elem)
}

// Append appends a new single element to end of self
func (self *Slice[T]) Append(elems ...T) {
	self.data = append(self.data, elems...)
}

// Concat appends every element in elems to end of self
func (self *Slice[T]) Concat(elems []T) {
	self.data = append(self.data, elems...)
}

// At returns the element at nth index of self
func (self Slice[T]) At(n int) T {
	return self.data[n]
}

// Set modifies the element at nth index of self
func (self *Slice[T]) Set(index int, value T) {
	self.data[index] = value
}

// Join stringifies each element and joins it using the given sep
func (self Slice[T]) Join(sep string) String.String {
	strList := make([]string, self.Length())
	for i, v := range self.data {
		strList[i] = fmt.Sprint(v)
	}
	return String.New(strings.Join(strList, sep))
}
