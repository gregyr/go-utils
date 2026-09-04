package streams

import (
	"bufio"
	"os"
	"sync"
)

type ParallelStream[T any] struct {
	buffer     chan T
	routines   int
	buffersize int
}

func ParallelStreamSlice[T any](s []T, routines int, buffersize int) *ParallelStream[T] {
	stream := &ParallelStream[T]{
		buffer:     make(chan T, buffersize),
		routines:   routines,
		buffersize: buffersize,
	}

	go func() {
		defer close(stream.buffer)
		for _, value := range s {
			stream.buffer <- value
		}
	}()

	return stream
}

// streams a file one line at a time
func ParallelStreamFile(file *os.File, routines int, buffersize int) *ParallelStream[[]byte] {

	stream := &ParallelStream[[]byte]{
		buffer:     make(chan []byte, buffersize),
		routines:   routines,
		buffersize: buffersize,
	}

	go func() {
		defer close(stream.buffer)
		buf := bufio.NewScanner(file)
		for buf.Scan() {
			stream.buffer <- buf.Bytes()
		}
		if buf.Err() != nil {
			panic(buf.Err())
		}
	}()

	return stream
}

func (s *ParallelStream[T]) Map[V any](fn func(T) V) *ParallelStream[V] {
	stream := &ParallelStream[V]{
		buffer:     make(chan V, s.buffersize),
		routines:   s.routines,
		buffersize: s.buffersize,
	}

	var wg sync.WaitGroup

	for range stream.routines {
		wg.Go(func() {
			for value := range s.buffer {
				stream.buffer <- fn(value)
			}
		})
	}

	go func() {
		wg.Wait()
		close(stream.buffer)
	}()

	return stream
}

func (s *ParallelStream[T]) Filter(fn func(T) bool) *ParallelStream[T] {
	stream := &ParallelStream[T]{
		buffer:     make(chan T, s.buffersize),
		routines:   s.routines,
		buffersize: s.buffersize,
	}

	var wg sync.WaitGroup

	for range stream.routines {
		wg.Go(func() {
			for value := range s.buffer {
				if fn(value) {
					stream.buffer <- value
				}
			}
		})
	}

	go func() {
		wg.Wait()
		close(stream.buffer)
	}()

	return stream
}

func (s *ParallelStream[T]) Collect[V any](base V, fn func(V, T) V) V {

	for value := range s.buffer {
		base = fn(base, value)
	}

	return base
}

// collects the values of the stream to a slice, this does not guarantee order
func (s *ParallelStream[T]) ToSlice() []T {
	c := make([]T, 0)

	for value := range s.buffer {
		c = append(c, value)
	}

	return c
}
