package set

import (
	"fmt"
	"iter"
	"strings"
)

type Set[T comparable] map[T]struct{}

// formats the Set as a String with curly braces
func (s Set[T]) String() string {
	var sStr strings.Builder
	fmt.Fprint(&sStr, "{")
	i := 0
	for k := range s {
		fmt.Fprint(&sStr, k)
		if i != len(s)-1 {
			fmt.Fprintf(&sStr, ", ")
		}
		i++
	}
	fmt.Fprint(&sStr, "}")
	return sStr.String()
}

// creates a new Empty Set
func New[T comparable]() Set[T] {
	return make(Set[T])
}

// creates a new set from a given iterator like slice elements or map values
func FromIter[T comparable](sl iter.Seq2[any, T]) Set[T] {
	s := New[T]()
	for _, e := range sl {
		s[e] = struct{}{}
	}
	return s
}

// returns a set of the given elements
func Of[T comparable](elem ...T) Set[T] {
	s := New[T]()
	for _, e := range elem {
		s[e] = struct{}{}
	}
	return s
}

// adds an element to the set
func (s Set[T]) Add(elem T) {
	s[elem] = struct{}{}
}

// removes the element from the set
func (s Set[T]) Remove(elem T) {
	delete(s, elem)
}

// reports wheter the set contains an element
func (s Set[T]) Contains(elem T) bool {
	_, exists := s[elem]
	return exists
}

// returns the amount of elements in the Set
func (s Set[T]) Len() int {
	return len(s)
}

// creates a new Set of the Union
func (s Set[T]) Union(s2 Set[T]) Set[T] {
	union := New[T]()
	for k := range s {
		union[k] = struct{}{}
	}
	for k := range s2 {
		union[k] = struct{}{}
	}
	return union
}

// creates a new Set of the Intersection
func (s Set[T]) Intersection(s2 Set[T]) Set[T] {
	intersection := New[T]()
	for k := range s {
		if _, in := s2[k]; in {
			intersection[k] = struct{}{}
		}
	}
	return intersection
}

// Creates a new Set of the Difference
func (s Set[T]) Difference(s2 Set[T]) Set[T] {
	difference := New[T]()
	for k := range s {
		if _, in := s2[k]; !in {
			difference[k] = struct{}{}
		}
	}
	return difference
}

// works like Union but in-place
func (s Set[T]) UnionWith(s2 Set[T]) {
	for k := range s2 {
		s[k] = struct{}{}
	}
}

// works like Intersection but in-place
func (s Set[T]) IntersectionWith(s2 Set[T]) {
	for k := range s {
		if _, in := s2[k]; !in {
			delete(s, k)
		}
	}
}

// works like Difference but in-place
func (s Set[T]) DifferenceWith(s2 Set[T]) {
	for k := range s {
		if _, in := s2[k]; in {
			delete(s, k)
		}
	}
}
