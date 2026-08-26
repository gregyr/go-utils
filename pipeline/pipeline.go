// Package pipeline provides simple, channel-based data streams. Create a stream from a slice, chain operations such as mapping and filtering, and finish with collectors
// The Package streams provides a simpler usage for go versions > 1.27.0
package pipeline

// A Pipeline describes a Pipeline operation that is run on every element of the Data Stream. It Offers chaining of Operations and Collection of the Elements of a Stream
type Pipeline interface {
	out() <-chan any
	in() chan<- any
	// Then runs another pipeline step and provides the interface to chain further
	Then(Pipeline) Pipeline
	// Finally runs a collector on the Pipelines data, it finishes the Pipeline and returns the calculated value. Some Collectors may return nothing or errors, this depends on the Collector that is Used
	Finally(Collector) any
}

type MapPipeline[T any, V any] struct {
	fn         func(T) V
	channelIn  chan any
	channelOut chan any
}

// Map Creates a new Pipeline step that maps the given value through a given function on another value. The incoming and outgoing types may differ
func Map[T any, V any](fn func(T) V) *MapPipeline[T, V] {

	f := &MapPipeline[T, V]{
		fn:         fn,
		channelIn:  make(chan any),
		channelOut: make(chan any),
	}

	f.process()
	return f
}

// Then runs another pipeline step and provides the interface to chain further
func (f *MapPipeline[T, V]) Then(p Pipeline) Pipeline {
	go ThenImpl(f, p)
	return p
}

// Finally returns the value calculated by a given collector
func (f *MapPipeline[T, V]) Finally(c Collector) any {
	go CollectImpl(f, c)
	return c.get()
}

func (f *MapPipeline[T, V]) process() {
	go func() {
		defer close(f.channelOut)
		for t := range f.channelIn {
			f.channelOut <- f.fn(t.(T))
		}
	}()
}

func (f *MapPipeline[T, V]) out() <-chan any {
	return f.channelOut
}

func (f *MapPipeline[T, V]) in() chan<- any {
	return f.channelIn
}

// FILTER

type FilterPipeline[T any] struct {
	fn         func(T) bool
	channelIn  chan any
	channelOut chan any
}

// Filter Creates a new Pipeline step that filters the stream based on a given function. The Pipeline only keeps the values for which the function returns true.
func Filter[T any](fn func(T) bool) *FilterPipeline[T] {

	f := &FilterPipeline[T]{
		fn:         fn,
		channelIn:  make(chan any),
		channelOut: make(chan any),
	}

	f.process()
	return f
}

func (f *FilterPipeline[T]) Then(p Pipeline) Pipeline {
	go ThenImpl(f, p)
	return p
}

func (f *FilterPipeline[T]) process() {
	go func() {
		defer close(f.channelOut)
		for t := range f.channelIn {
			if f.fn(t.(T)) {
				f.channelOut <- t
			}
		}
	}()
}

func (f *FilterPipeline[T]) out() <-chan any {
	return f.channelOut
}

func (f *FilterPipeline[T]) in() chan<- any {
	return f.channelIn
}

func (f *FilterPipeline[T]) Finally(c Collector) any {
	go CollectImpl(f, c)
	return c.get()
}

func ThenImpl(f Pipeline, p Pipeline) {
	defer close(p.in())
	for e := range f.out() {
		p.in() <- e
	}
}

func CollectImpl(p Pipeline, c Collector) {
	defer close(c.in())
	for e := range p.out() {
		c.in() <- e
	}
}
