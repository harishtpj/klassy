// package String provides a custom String type with chainable methods
// with API similar to that of the standard library's strings package.
package string

import "strings"

// type String is alias for native string
type String string

// New return a new instance of the String type
func New(s string) String {
	return String(s)
}

// Value returns the original underlying string value
func (self String) Value() string {
	return string(self)
}

// ToLower returns the Lowercased version of self
func (self String) ToLower() String {
	return New(strings.ToLower(string(self)))
}

// ToLowerSpecial
// ToTitle
// ToTitleSpecial

// ToUpper returns the Uppercased version of self
func (self String) ToUpper() String {
	return New(strings.ToUpper(string(self)))
}

// ToUpperSpecial
// ToValidUTF8

// Trim return the sliced version of self with all leading and trailing
// characters in cutset removed
func (self String) Trim(cutset string) String {
	return New(strings.Trim(self.Value(), cutset))
}


