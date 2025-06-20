// package String provides a custom String type with chainable methods
// with API similar to that of the standard library's strings package.
package String

import (
	"fmt"
	"iter"
	"strings"
	"unicode"

	"github.com/harishtpj/klassy/Slice"
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

// Count counts the number of non-overlapping instances of substr in self. 
// If substr is an empty string, Count returns 1 + the number of characters in self.
func (self String) Count(substr string) int {
	return strings.Count(self.Value(), substr)
}

// Cut slices self around the first instance of sep, returning the text before and after sep.
// The found result reports whether sep appears in self. If sep does not appear in self,
// Cut returns self, "", false.
func (self String) Cut(sep string) (before, after String, found bool) {
	b, a, f := strings.Cut(self.Value(), sep)
	return New(b), New(a), f
}

// CutPrefix returns self without the provided leading prefix string and reports 
// whether it found the prefix. If self doesn't start with prefix, CutPrefix 
// returns self, false. If prefix is the empty string, CutPrefix returns self, true.
func (self String) CutPrefix(prefix string) (after String, found bool) {
	a, f := strings.CutPrefix(self.Value(), prefix)
	return New(a), f
}

// CutSuffix returns self without the provided ending suffix string and reports 
// whether it found the suffix. If self doesn't end with suffix, CutSuffix returns 
// self, false. If suffix is the empty string, CutSuffix returns self, true.
func (self String) CutSuffix(suffix string) (before String, found bool) {
	b, f := strings.CutSuffix(self.Value(), suffix)
	return New(b), f
}

// EqualFold reports whether self and t, interpreted as UTF-8 strings, are equal 
// under simple Unicode case-folding, which is a more general form of case-insensitivity.
func (self String) EqualFold(t string) bool {
	return strings.EqualFold(self.Value(), t)
}

// Fields splits the String self around each instance of one or more consecutive
// white space characters, as defined by unicode.IsSpace, returning a slice of 
// substrings of self or an empty slice if self contains only white space.
func (self String) Fields() Slice.Slice[String] {
	return Slice.MapTo(Slice.New(strings.Fields(self.Value())), New)
}

// FieldsFunc splits self at each run of character c satisfying f(c) and returns 
// an array of Slices of String. If all characters in self satisfy f(c) or the 
// string is empty, an empty slice is returned.
//
// FieldsFunc makes no guarantees about the order in which it calls f(c) and 
// assumes that f always returns the same value for a given c.
func (self String) FieldsFunc(f func(rune) bool) Slice.Slice[String] {
	data := strings.FieldsFunc(self.Value(), f)
	return Slice.MapTo(Slice.New(data), New)
}

// FieldsFuncSeq returns an iterator over substrings of self split around runs 
// of characters satisfying f(c). The iterator yields the same strings that 
// would be returned by self.[FieldsFunc](), but without constructing the slice.
func (self String) FieldsFuncSeq(f func(rune) bool) iter.Seq[String] {
	strFields := strings.FieldsFuncSeq(self.Value(), f)

	return func(yield func(String) bool) {
		for field := range strFields {
			if !yield(New(field)) {
				return
			}
		}
	}
}

// FieldsSeq returns an iterator over substrings of self split around runs of 
// whitespace characters, as defined by unicode.IsSpace. The iterator yields 
// the same strings that would be returned by self.[Fields](), but without 
// constructing the slice.
func (self String) FieldsSeq() iter.Seq[String] {
	strFields := strings.FieldsSeq(self.Value())

	return func(yield func(String) bool) {
		for field := range strFields {
			if !yield(New(field)) {
				return
			}
		}
	}
}

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

// IndexByte returns the index of the first instance of c in self, 
// or -1 if c is not present in self.
func (self String) IndexByte(c byte) int {
	return strings.IndexByte(self.Value(), c)
}

// IndexFunc returns the index into self of the first character 
// satisfying f(c), or -1 if none do.
func (self String) IndexFunc(f func(rune) bool) int {
	return strings.IndexFunc(self.Value(), f)
}

// IndexRune returns the index of the first instance of the character r, 
// or -1 if rune is not present in s. If r is utf8.RuneError, it returns 
// the first instance of any invalid UTF-8 byte sequence.
func (self String) IndexRune(r rune) int {
	return strings.IndexRune(self.Value(), r)
}

// Join stringifies each element in elems and joins it using self
// Works similar to Python's str.join method
func (self String) Join(elems Slice.Slice[any]) String {
	strSlice := Slice.MapTo(elems, func(v any) string { return fmt.Sprint(v) })
	return New(strings.Join(strSlice.Items, self.Value()))
}

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

// LastIndexByte returns the index of the last instance of c in self, 
// or -1 if c is not present in self.
func (self String) LastIndexByte(c byte) int {
	return strings.LastIndexByte(self.Value(), c)
}

// LastIndexFunc returns the index into self of the last 
// character satisfying f(c), or -1 if none do.
func (self String) LastIndexFunc(f func(rune) bool) int {
	return strings.LastIndexFunc(self.Value(), f)
}

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
func (self String) Split(sep string) Slice.Slice[String] {
	return Slice.MapTo(Slice.New(strings.Split(self.Value(), sep)), New)
}

// SplitAfter slices self into all substrings after each instance of 
// sep and returns a Slice of those substrings.
//
// If self does not contain sep and sep is not empty, SplitAfter returns 
// a Slice of length 1 whose only element is self.
//
// If sep is empty, SplitAfter splits after each UTF-8 sequence. If both 
// self and sep are empty, SplitAfter returns an empty Slice.
//
// It is equivalent to [SplitAfterN] with a count of -1.
func (self String) SplitAfter(sep string) Slice.Slice[String] {
	return Slice.MapTo(Slice.New(strings.SplitAfter(self.Value(), sep)), New)
}

// SplitAfterN slices self into substrings after each instance of sep and 
// returns a Slice of those substrings.
//
// The count determines the number of substrings to return:
//
// - n > 0: at most n substrings; the last substring will be the unsplit remainder;
// - n == 0: the result is nil (zero substrings);
// - n < 0: all substrings.
// Edge cases for self and sep (for example, empty strings) are handled as 
// described in the documentation for [SplitAfter].
func (self String) SplitAfterN(sep string, n int) Slice.Slice[String] {
	return Slice.MapTo(Slice.New(strings.SplitAfterN(self.Value(), sep, n)), New)
}

// SplitAfterSeq returns an iterator over substrings of self split after each 
// instance of sep. The iterator yields the same strings that would be returned 
// by self.[SplitAfter](sep), but without constructing the slice. It returns a 
// single-use iterator.
func (self String) SplitAfterSeq(sep string) iter.Seq[String] {
	strSplits := strings.SplitAfterSeq(self.Value(), sep)

	return func(yield func(String) bool) {
		for split := range strSplits {
			if !yield(New(split)) {
				return
			}
		}
	}
}

// SplitN slices self into substrings separated by sep and returns a slice of 
// the substrings between those separators.
//
// The count determines the number of substrings to return:
//
// - n > 0: at most n substrings; the last substring will be the unsplit remainder;
// - n == 0: the result is nil (zero substrings);
// - n < 0: all substrings.
// Edge cases for self and sep (for example, empty strings) are handled as 
// described in the documentation for [Split].
//
// To split around the first instance of a separator, see [Cut].
func (self String) SplitN(sep string, n int) Slice.Slice[String] {
	return Slice.MapTo(Slice.New(strings.SplitN(self.Value(), sep, n)), New)
}

// SplitSeq returns an iterator over all substrings of self separated by sep. 
// The iterator yields the same strings that would be returned by self.[Split](sep), 
// but without constructing the slice. It returns a single-use iterator.
func (self String) SplitSeq(sep string) iter.Seq[String] {
	strSplits := strings.SplitSeq(self.Value(), sep)

	return func(yield func(String) bool) {
		for split := range strSplits {
			if !yield(New(split)) {
				return
			}
		}
	}
}

// Depreciated: Title

// ToLower returns the Lowercased version of self
func (self String) ToLower() String {
	return New(strings.ToLower(self.Value()))
}

// ToLowerSpecial returns a copy of self with all Unicode letters mapped 
// to their lower case using the case mapping specified by c.
func (self String) ToLowerSpecial(c unicode.SpecialCase) String {
	return New(strings.ToLowerSpecial(c, self.Value()))
}

// ToTitle returns a copy of self with all Unicode letters 
// mapped to their Unicode title case.
func (self String) ToTitle() String {
	return New(strings.ToTitle(self.Value()))
}

// ToTitleSpecial returns a copy of self with all Unicode letters mapped 
// to their Unicode title case, giving priority to the special casing rules.
func (self String) ToTitleSpecial(c unicode.SpecialCase) String {
	return New(strings.ToTitleSpecial(c, self.Value()))
}

// ToUpper returns the Uppercased version of self
func (self String) ToUpper() String {
	return New(strings.ToUpper(self.Value()))
}

// ToUpperSpecial returns a copy of self with all Unicode letters mapped 
// to their upper case using the case mapping specified by c.
func (self String) ToUpperSpecial(c unicode.SpecialCase) String {
	return New(strings.ToUpperSpecial(c, self.Value()))
}

// ToValidUTF8 returns a copy of self with each run of invalid UTF-8 
// byte sequences replaced by the replacement string, which may be empty.
func (self String) ToValidUTF8(replacement string) String {
	return New(strings.ToValidUTF8(self.Value(), replacement))
}

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
