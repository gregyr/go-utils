// Package streams provides a generic implementation of a Stream. It allows for streaming of Slices and chaining operations on such streams. It utilizes the new Generic Methods added in Go 1.27.0
package streams

type Stream[T any] chan T

func StreamSlice[T any](s []T) Stream[T] {
	c := make(chan T)

	go func() {
		defer close(c)
		for _, value := range s {
			c <- value
		}
	}()

	return c
}

func (s Stream[T]) Map[V any](fn func(T) V) Stream[V] {
	c := make(chan V)

	go func() {
		defer close(c)
		for value := range s {
			c <- fn(value)
		}
	}()

	return c
}

func (s Stream[T]) Filter(fn func(T) bool) Stream[T] {
	c := make(chan T)

	go func() {
		defer close(c)
		for value := range s {
			if fn(value) {
				c <- value
			}
		}
	}()

	return c
}

func (s Stream[T]) Collect[V any](base V, fn func(V, T) V) V {

	for value := range s {
		base = fn(base, value)
	}

	return base
}

func (s Stream[T]) ToSlice() []T {
	c := make([]T, 0)

	for value := range s {
		c = append(c, value)
	}

	return c
}
