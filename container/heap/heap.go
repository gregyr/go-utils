// package heap provides a generic heap implementation for Go's standard heap interface
package heap

import stdheap "container/heap"

type Heap[T any] struct {
	arr  []T
	less func(a, b T) bool
}

// initialises a new empty heap with a given comparison function, reporting what element is deemed smaller, can be therefore used to create both min and max heaps
func New[T any](less func(a, b T) bool) *Heap[T] {
	if less == nil {
		panic("heap: nil comparator")
	}
	return &Heap[T]{less: less}
}

// creates a new heap containing the elements of a give slice
func FromSlice[T any](values []T, less func(a, b T) bool) *Heap[T] {
	h := New(less)
	h.arr = append(h.arr, values...)
	stdheap.Init(heapAdapter[T]{h})
	return h
}

// returns the amount of elements in a heap
func (h *Heap[T]) Len() int { return len(h.arr) }

// pushes a value to the heap
func (h *Heap[T]) Push(value T) {
	stdheap.Push(heapAdapter[T]{h}, value)
}

// pops the smallest value off the heap and returns true if succesful and false if the heap was empty
func (h *Heap[T]) Pop() (T, bool) {
	if h.Len() == 0 {
		var zero T
		return zero, false
	}
	return stdheap.Pop(heapAdapter[T]{h}).(T), true
}

type heapAdapter[T any] struct {
	h *Heap[T]
}

func (a heapAdapter[T]) Len() int           { return a.h.Len() }
func (a heapAdapter[T]) Less(i, j int) bool { return a.h.less(a.h.arr[i], a.h.arr[j]) }
func (a heapAdapter[T]) Swap(i, j int)      { a.h.arr[i], a.h.arr[j] = a.h.arr[j], a.h.arr[i] }

func (a heapAdapter[T]) Push(value any) {
	v, ok := value.(T)
	if !ok {
		panic("heap: pushed value has the wrong type")
	}
	a.h.arr = append(a.h.arr, v)
}

func (a heapAdapter[T]) Pop() any {
	last := len(a.h.arr) - 1
	value := a.h.arr[last]

	var zero T
	a.h.arr[last] = zero
	a.h.arr = a.h.arr[:last]

	return value
}
