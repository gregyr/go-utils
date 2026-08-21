package set

import (
	"fmt"
	"iter"
	"strings"
)

type set[T comparable] map[T]struct{}

func Set[T comparable]() set[T] {
	return make(set[T])
}

func SetFromSlice[T comparable](sl []T) set[T] {
	s := Set[T]()
	for _, e := range sl {
		s[e] = struct{}{}
	}
	return s
}

func (s set[T]) String() string {
	var sStr strings.Builder
	fmt.Fprint(&sStr, "(")
	i := 0
	for k := range s {
		fmt.Fprint(&sStr, k)
		if i != len(s)-1 {
			fmt.Fprintf(&sStr, ", ")
		}
		i++
	}
	fmt.Fprint(&sStr, ")")
	return sStr.String()
}

func (s set[T]) Add(t T) T {
	s[t] = struct{}{}
	return t
}

func (s set[T]) Contains(t T) bool {
	_, exists := s[t]
	return exists
}

func (s set[T]) Remove(t T) T {
	delete(s, t)
	return t
}

// Merge behaves like Union, however it merges u into s therefore mutation s in place
func (s set[T]) Merge(u set[T]) {
	for k := range u {
		s[k] = struct{}{}
	}
}

func Union[T comparable](s, u set[T]) set[T] {
	union := Set[T]()
	for k := range u {
		union[k] = struct{}{}
	}
	for k := range s {
		union[k] = struct{}{}
	}
	return union
}

func Intersection[T comparable](u, s set[T]) set[T] {
	intersection := Set[T]()
	if len(s) > len(u) {
		for k := range u {
			if _, ok := s[k]; ok {
				intersection[k] = struct{}{}
			}
		}
	} else {
		for k := range s {
			if _, ok := u[k]; ok {
				intersection[k] = struct{}{}
			}
		}
	}
	return intersection
}

func (s set[T]) Equals(u set[T]) bool {
	if len(s) != len(u) {
		return false
	}
	for k := range s {
		if _, exists := u[k]; !exists {
			return false
		}
	}
	return true
}

func (s set[T]) Iter() iter.Seq[T] {
	return func(yield func(T) bool) {
		for k := range s {
			if !yield(k) {
				return
			}
		}
	}
}
