package slice

import (
	"iter"
	"slices"
)

// provides a slice over a range
func Range(n int) []int {
	r := make([]int, n)
	for i := range n {
		r[i] = i
	}
	return r
}

func RangeOver(start, length, step int) []int {
	r := make([]int, length)
	for i := range length {
		r[i] = start + (i * step)
	}
	return r
}

func RangeFunc[T any](n int, fillFn func(i int) T) []T {
	r := make([]T, n)
	for i := range n {
		r[i] = fillFn(i)
	}
	return r
}

func IterToSlice[T any](i iter.Seq[T]) []T {
	s := make([]T, 0)
	for e := range i {
		s = append(s, e)
	}
	return s
}

// splits a slice in two before the first occurence of elem, meaning that elem is the first element of the second slice.
// if elem is not in s, Cut returns s and an empty slice
// if elem is the first element of s, cut returns an empty slice and s
func Cut[T comparable](s []T, elem T) ([]T, []T) {
	i := slices.Index(s, elem)
	if i == -1 {
		return s, []T{}
	}
	return s[:i], s[i:]
}

// splits a slice in two after the first occurence of elem, meaning that elem is the last element of the first slice.
// if elem is not in s or elem is the last element of s, Cut returns s and an empty slice
func CutAfter[T comparable](s []T, elem T) ([]T, []T) {
	i := slices.Index(s, elem)
	if i == -1 || i == len(s)-1 {
		return s, []T{}
	}
	return s[:i+1], s[i+1:]
}

// splits the slice before the given index meaning that the i'th element of s is the first element of the second slice starting at 0
// if i < 0 or i > len(s) - 1, CutAt returns s and an empty slice
func CutAt[T any](s []T, i int) ([]T, []T) {
	if i < 0 || i > len(s)-1 {
		return s, []T{}
	}
	return s[:i], s[i:]
}

// splits the slice at the first occurence of the element returning all elements before, the element itself and all elements after
// if elem is not present in the slice, it returns s, the default value for T and an empty slice
func CutOut[T comparable](s []T, elem T) ([]T, T, []T) {
	i := slices.Index(s, elem)
	if i == -1 {
		return s, *new(T), []T{}
	}
	return s[:i], s[i], s[i+1:]
}

// Filters the given slice based on the given function, this creates a new slice
func Filter[T comparable](s []T, fn func(T) bool) []T {
	new := make([]T, 0, len(s))
	for _, e := range s {
		if fn(e) {
			new = append(new, e)
		}
	}
	return new
}

// Maps every value of the given slice onto a new value, this creates a new slice
func Map[T, V comparable](s []T, fn func(T) V) []V {
	new := make([]V, len(s))
	for i, e := range s {
		new[i] = fn(e)
	}
	return new
}

// reduces the slice into a given type starting at a base value
func Reduce[T comparable, A any](s []T, base A, fn func(base A, v T) A) A {
	for _, e := range s {
		base = fn(base, e)
	}
	return base
}

// Casting a slice to a Pipe allows chaining of Operations like Filter, Map, Reduce
type Pipe[T comparable] []T

// Filters the given Slice based on the function, this returns a new Pipe which can be assigned to a slice of type T
func (p Pipe[T]) Filter(fn func(T) bool) Pipe[T] {
	new := make(Pipe[T], 0, len(p))
	for _, e := range p {
		if fn(e) {
			new = append(new, e)
		}
	}
	return new
}

// Maps every value of the given slice onto a new value, this returns a new Pipe which can be assigned to a slice of type T
func (p Pipe[T]) Map[V comparable](fn func(T) V) Pipe[V] {
	new := make(Pipe[V], len(p))
	for i, e := range p {
		new[i] = fn(e)
	}
	return new
}

// reduces the slice into a given type starting at a base value
func (p Pipe[T]) Reduce[A any](base A, fn func(base A, v T) A) A {
	for _, e := range p {
		base = fn(base, e)
	}
	return base
}
