package pipeline

import (
	"errors"
)

type Collector interface {
	in() chan<- any
	get() any
}

type Aggregator[T any, V any] struct {
	fn        func(V, T) V
	acc       V
	channelIn chan any
	done      chan struct{}
}

// Aggregate accumulates the Pipelines Data based on the provided function and base value
//
// This Aggregator would sum up every value of a Pipeline for example:
//
//	Aggregate(0, func(base, a int) int {return base + a})
func Aggregate[T, V any](base V, fn func(V, T) V) *Aggregator[T, V] {

	c := &Aggregator[T, V]{
		fn:        fn,
		acc:       base,
		channelIn: make(chan any),
		done:      make(chan struct{}),
	}

	c.process()
	return c
}

func (c *Aggregator[T, V]) process() {
	go func() {
		defer close(c.done)
		for t := range c.channelIn {
			c.acc = c.fn(c.acc, t.(T))
		}
	}()
}

func (s *Aggregator[T, V]) in() chan<- any {
	return s.channelIn
}

func (c *Aggregator[T, V]) get() any {
	<-c.done
	return c.acc
}

type Number interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 | uintptr |
		float32 | float64
}

type SumAggregator[T Number] struct {
	acc       T
	channelIn chan any
	done      chan struct{}
}

func Sum[T Number]() *SumAggregator[T] {
	c := &SumAggregator[T]{
		acc:       *new(T),
		channelIn: make(chan any),
		done:      make(chan struct{}),
	}

	c.process()
	return c
}

func (c *SumAggregator[T]) process() {
	go func() {
		defer close(c.done)
		for t := range c.channelIn {
			c.acc += t.(T)
		}
	}()
}

func (s *SumAggregator[T]) in() chan<- any {
	return s.channelIn
}

func (c *SumAggregator[T]) get() any {
	<-c.done
	return c.acc
}

type SliceAggregator[T any] struct {
	acc       []T
	channelIn chan any
	done      chan struct{}
}

func ToSlice[T any]() *SliceAggregator[T] {

	c := &SliceAggregator[T]{
		acc:       make([]T, 0),
		channelIn: make(chan any),
		done:      make(chan struct{}),
	}

	c.process()
	return c
}

func (c *SliceAggregator[T]) process() {
	go func() {
		defer close(c.done)
		for t := range c.channelIn {
			c.acc = append(c.acc, t.(T))
		}
	}()
}

func (s *SliceAggregator[T]) in() chan<- any {
	return s.channelIn
}

func (c *SliceAggregator[T]) get() any {
	<-c.done
	return c.acc
}

type ForEachCollector[T any] struct {
	fn        func(T)
	channelIn chan any
	done      chan struct{}
}

// ForEach ends a pipeline by running a function on every element. The Collector always returns nil
func ForEach[T any](fn func(T)) *ForEachCollector[T] {
	fe := &ForEachCollector[T]{
		fn:        fn,
		channelIn: make(chan any),
		done:      make(chan struct{}),
	}

	fe.process()
	return fe
}

func (s *ForEachCollector[T]) in() chan<- any {
	return s.channelIn
}

func (fe *ForEachCollector[T]) process() {
	go func() {
		defer close(fe.done)
		for t := range fe.channelIn {
			fe.fn(t.(T))
		}
	}()
}

func (fe *ForEachCollector[T]) get() any {
	<-fe.done
	return nil
}

type ForEachErrorCollector[T any] struct {
	fn        func(T) error
	acc       error
	channelIn chan any
	done      chan struct{}
}

// ForEachError ends a pipeline by running a function on every element. The returned value is the joined error of each run, if no error occured it returns nil.
func ForEachError[T any](fn func(T) error) *ForEachErrorCollector[T] {
	fe := &ForEachErrorCollector[T]{
		fn:        fn,
		channelIn: make(chan any),
		done:      make(chan struct{}),
	}

	fe.process()
	return fe
}

func (s *ForEachErrorCollector[T]) in() chan<- any {
	return s.channelIn
}

func (fe *ForEachErrorCollector[T]) process() {
	go func() {
		defer close(fe.done)
		for t := range fe.channelIn {
			fe.acc = errors.Join(fe.acc, fe.fn(t.(T)))
		}
	}()
}

func (fe *ForEachErrorCollector[T]) get() any {
	<-fe.done
	return fe.acc
}
