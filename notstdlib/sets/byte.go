/*
Copyright 2015 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by set-gen. DO NOT EDIT.

package sets

import (
	"reflect"
	"sort"
)

// sets.Byte is a set of bytes, implemented via map[byte]struct{} for minimal memory consumption.
type Byte map[byte]Empty

// NewByte creates a Byte from a list of values.
func NewByte(items ...byte) Byte {
	ss := Byte{}
	ss.Insert(items...)
	return ss
}

// ByteKeySet creates a Byte from a keys of a map[byte](? extends interface{}).
// If the value passed in is not actually a map, this will panic.
func ByteKeySet(theMap interface{}) Byte {
	v := reflect.ValueOf(theMap)
	ret := Byte{}

	for _, keyValue := range v.MapKeys() {
		ret.Insert(keyValue.Interface().(byte))
	}
	return ret
}

// Insert adds items to the set.
func (s Byte) Insert(items ...byte) Byte {
	for _, item := range items {
		s[item] = Empty{}
	}
	return s
}

// Delete removes all items from the set.
func (s Byte) Delete(items ...byte) Byte {
	for _, item := range items {
		delete(s, item)
	}
	return s
}

// Has returns true if and only if item is contained in the set.
func (s Byte) Has(item byte) bool {
	_, contained := s[item]
	return contained
}

// HasAll returns true if and only if all items are contained in the set.
func (s Byte) HasAll(items ...byte) bool {
	for _, item := range items {
		if !s.Has(item) {
			return false
		}
	}
	return true
}

// HasAny returns true if any items are contained in the set.
func (s Byte) HasAny(items ...byte) bool {
	for _, item := range items {
		if s.Has(item) {
			return true
		}
	}
	return false
}

// Difference returns a set of objects that are not in s2
// For example:
// s1 = {a1, a2, a3}
// s2 = {a1, a2, a4, a5}
// s1.Difference(s2) = {a3}
// s2.Difference(s1) = {a4, a5}
func (s Byte) Difference(s2 Byte) Byte {
	result := NewByte()
	for key := range s {
		if !s2.Has(key) {
			result.Insert(key)
		}
	}
	return result
}

// Union returns a new set which includes items in either s1 or s2.
// For example:
// s1 = {a1, a2}
// s2 = {a3, a4}
// s1.Union(s2) = {a1, a2, a3, a4}
// s2.Union(s1) = {a1, a2, a3, a4}
func (s1 Byte) Union(s2 Byte) Byte {
	result := NewByte()
	for key := range s1 {
		result.Insert(key)
	}
	for key := range s2 {
		result.Insert(key)
	}
	return result
}

// Intersection returns a new set which includes the item in BOTH s1 and s2
// For example:
// s1 = {a1, a2}
// s2 = {a2, a3}
// s1.Intersection(s2) = {a2}
func (s1 Byte) Intersection(s2 Byte) Byte {
	var walk, other Byte
	result := NewByte()
	if s1.Len() < s2.Len() {
		walk = s1
		other = s2
	} else {
		walk = s2
		other = s1
	}
	for key := range walk {
		if other.Has(key) {
			result.Insert(key)
		}
	}
	return result
}

// IsSuperset returns true if and only if s1 is a superset of s2.
func (s1 Byte) IsSuperset(s2 Byte) bool {
	for item := range s2 {
		if !s1.Has(item) {
			return false
		}
	}
	return true
}

// Equal returns true if and only if s1 is equal (as a set) to s2.
// Two sets are equal if their membership is identical.
// (In practice, this means same elements, order doesn't matter)
func (s1 Byte) Equal(s2 Byte) bool {
	return len(s1) == len(s2) && s1.IsSuperset(s2)
}

type sortableSliceOfByte []byte

func (s sortableSliceOfByte) Len() int           { return len(s) }
func (s sortableSliceOfByte) Less(i, j int) bool { return lessByte(s[i], s[j]) }
func (s sortableSliceOfByte) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// List returns the contents as a sorted byte slice.
func (s Byte) List() []byte {
	res := make(sortableSliceOfByte, 0, len(s))
	for key := range s {
		res = append(res, key)
	}
	sort.Sort(res)
	return []byte(res)
}

// UnsortedList returns the slice with contents in random order.
func (s Byte) UnsortedList() []byte {
	res := make([]byte, 0, len(s))
	for key := range s {
		res = append(res, key)
	}
	return res
}

// Returns a single element from the set.
func (s Byte) PopAny() (byte, bool) {
	for key := range s {
		s.Delete(key)
		return key, true
	}
	var zeroValue byte
	return zeroValue, false
}

// Len returns the size of the set.
func (s Byte) Len() int {
	return len(s)
}

func lessByte(lhs, rhs byte) bool {
	return lhs < rhs
}
