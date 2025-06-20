// package String provides a custom String type with chainable methods
// with API similar to that of the standard library's strings package.
package String

import (
	"iter"
	"strings"
)

// type String is alias for native string
type String string

// New return a new instance of the String type
func New(s string) String {
	return String(s)
}

// Value return the original underlying string value
func (self String) Value() string {
	return string(self)
}

// Length return the length of underlying string
func (self String) Length() int {
	return len(self.Value())
}

// Contains reports whether substr is within self
func (self String) Contains(substr string) bool {
	return strings.Contains(self.Value(), substr)
}

// ContainsAny reports whether any character in chars is within self
func (self String) ContainsAny(chars string) bool {
	return strings.ContainsAny(self.Value(), chars)
}

// ContainsFunc reports if any character c in self satisfy f(c)
func (self String) ContainsFunc(f func(rune) bool) bool {
	return strings.ContainsFunc(self.Value(), f)
}

// ContainsRune reports whether character r is within self
func (self String) ContainsRune(r rune) bool {
	return strings.ContainsRune(self.Value(), r)
}

// TODO: Count

// Cut slices self around the first instance of sep, returning the text before and after sep.
// The found result reports whether sep appears in self. If sep does not appear in self,
// Cut returns self, "", false.
func (self String) Cut(sep string) (before, after String, found bool) {
	b, a, f := strings.Cut(self.Value(), sep)
	return New(b), New(a), f
}

// TODO: CutPrefix
// TODO: CutSuffix
// TODO: EqualFold

// Fields splits the String self around each instance of one or more consecutive
// white space characters, as defined by unicode.IsSpace, returning a slice of 
// substrings of self or an empty slice if self contains only white space.
func (self String) Fields() []string {
	return strings.Fields(self.Value())
}

// TODO: FieldsFunc
// TODO: FieldsFuncSeq
// TODO: FieldsSeq

// HasPrefix reports if self starts with prefix
func (self String) HasPrefix(prefix string) bool {
	return strings.HasPrefix(self.Value(), prefix)
}

// HasSuffix reports if self ends with suffix
func (self String) HasSuffix(suffix string) bool {
	return strings.HasSuffix(self.Value(), suffix)
}

// Index returns the index of the first instance of substr in self,
// or -1 if substr is not present in self.
func (self String) Index(substr string) int {
	return strings.Index(self.Value(), substr)
}

// IndexAny returns the index of the first instance of any character 
// from chars in self or -1 if no character from chars is present in self
func (self String) IndexAny(chars string) int {
	return strings.IndexAny(self.Value(), chars)
}

// TODO: IndexByte
// TODO: IndexFunc
// TODO: IndexRune

// LastIndex returns the index of the last instance of substr in self, 
// or -1 if substr is not present in self.
func (self String) LastIndex(substr string) int {
	return strings.LastIndex(self.Value(), substr)
}

// LastIndexAny returns the index of the last instance of any character from 
// chars in self, or -1 if no character from chars is present in self.
func (self String) LastIndexAny(chars string) int {
	return strings.LastIndexAny(self.Value(), chars)
}

// TODO: LastIndexByte
// TODO: LastIndexFunc

// Lines returns an iterator over the newline-terminated lines in the 
// string self. The lines yielded by the iterator include their terminating 
// newlines. If self is empty, the iterator yields no lines at all. 
// If self does not end in a newline, the final yielded line will not end in 
// a newline. It returns a single-use iterator.
func (self String) Lines() iter.Seq[String] {
	strLines := strings.Lines(self.Value())

	return func(yield func(String) bool) {
		for line := range strLines {
			if !yield(New(line)) {
				return
			}
		}
	}
}

// Map returns a copy of the string self with all its characters modified 
// according to the mapping function. If mapping returns a negative value, 
// the character is dropped from the string with no replacement.
func (self String) Map(mapping func(rune) rune) String {
	return New(strings.Map(mapping, self.Value()))
}

// Repeat returns a new string consisting of count copies of the string self. 
//
// It panics if count is negative or if the result of (len(self) * count) overflows.
func (self String) Repeat(count int) String {
	return New(strings.Repeat(self.Value(), count))
}

// Replace returns a copy of the string self with the first n non-overlapping 
// instances of old replaced by new. If old is empty, it matches at the 
// beginning of the string and after each UTF-8 sequence, yielding up to k+1 
// replacements for a k-rune string. If n < 0, there is no limit 
// on the number of replacements.
func (self String) Replace(old, new string, n int) String {
	return New(strings.Replace(self.Value(), old, new, n))
}

// ReplaceAll returns a copy of the string self with all non-overlapping 
// instances of old replaced by new. If old is empty, it matches at the 
// beginning of the string and after each UTF-8 sequence, yielding up to k+1 
// replacements for a k-rune string.
//
// Equivalent to self.[Replace](old, new, -1)
func (self String) ReplaceAll(old, new string) String {
	return New(strings.ReplaceAll(self.Value(), old, new))
}

// Split slices self into all substrings separated by sep and returns a slice of
// the substrings between those separators.
//
// If self does not contain sep and sep is not empty, Split returns a
// slice of length 1 whose only element is self.
//
// If sep is empty, Split splits after each UTF-8 sequence. If both self
// and sep are empty, Split returns an empty slice.
//
// It is equivalent to [SplitN] with a count of -1.
//
// To split around the first instance of a separator, see [Cut].
func (self String) Split(sep string) []string {
	return strings.Split(self.Value(), sep)
}

// TODO: SplitAfter
// TODO: SplitAfterN
// TODO: SplitAfterSeq
// TODO: SplitN
// TODO: SplitSeq

// Depreciated: Title

// ToLower returns the Lowercased version of self
func (self String) ToLower() String {
	return New(strings.ToLower(self.Value()))
}

// TODO: ToLowerSpecial
// TODO: ToTitle
// TODO: ToTitleSpecial

// ToUpper returns the Uppercased version of self
func (self String) ToUpper() String {
	return New(strings.ToUpper(self.Value()))
}

// TODO: ToUpperSpecial
// TODO: ToValidUTF8

// Trim return the sliced version of self with all leading and trailing
// characters in cutset removed
func (self String) Trim(cutset string) String {
	return New(strings.Trim(self.Value(), cutset))
}

// TrimFunc return the sliced version of self with all leading and trailing
// characters satisfying f(c) removed
func (self String) TrimFunc(f func(rune) bool) String {
	return New(strings.TrimFunc(self.Value(), f))
}

// TrimLeft return the sliced version of self with all leading 
// characters in cutset removed
func (self String) TrimLeft(cutset string) String {
	return New(strings.TrimLeft(self.Value(), cutset))
}

// TrimLeftFunc return the sliced version of self with all leading 
// characters satisfying f(c) removed
func (self String) TrimLeftFunc(f func(rune) bool) String {
	return New(strings.TrimLeftFunc(self.Value(), f))
}

// TrimPrefix returns the sliced version of self with given prefix removed.
// If the prefix is not found in self, it is returned as it is
func (self String) TrimPrefix(prefix string) String {
	return New(strings.TrimPrefix(self.Value(), prefix))
}

// TrimRight return the sliced version of self with all leading 
// characters in cutset removed
func (self String) TrimRight(cutset string) String {
	return New(strings.TrimRight(self.Value(), cutset))
}

// TrimRightFunc return the sliced version of self with all leading 
// characters satisfying f(c) removed
func (self String) TrimRightFunc(f func(rune) bool) String {
	return New(strings.TrimRightFunc(self.Value(), f))
}

// TrimSpace return the sliced version of self with all leading
// and trailing whitespaces removed, as defined in Unicode
func (self String) TrimSpace() String {
	return New(strings.TrimSpace(self.Value()))
}
