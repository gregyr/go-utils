package pipeline

type SourceInterface interface {
	// Then attaches a new Pipeline step to the Pipeline
	Then(p Pipeline) Pipeline
	// Finally runs a collector on the Pipelines data, it finishes the Pipeline and returns the calculated value. Some Collectors may return nothing or errors, this depends on the Collector that is Used
	Finally(c Collector) any
	// the outgoing elements from the source
	out() <-chan any
}

type Source[T any] struct {
	channelOut chan any
}

// Stream creates a new source from a slice, from which Pipeline Operations can be chained
func StreamSlice[T any](s []T) *Source[T] {

	source := &Source[T]{
		channelOut: make(chan any),
	}

	go func() {
		defer close(source.channelOut)
		for _, t := range s {
			source.channelOut <- t
		}
	}()
	return source
}

func (s *Source[T]) out() <-chan any {
	return s.channelOut
}

// Then attaches a new Pipeline step to the source
func (s *Source[T]) Then(p Pipeline) Pipeline {
	go thenSource(s, p)
	return p
}

// Finally runs a collector on the Pipelines data, it finishes the Pipeline and returns the calculated value. Some Collectors may return nothing or errors, this depends on the Collector that is Used
func (s *Source[T]) Finally(c Collector) any {
	go collectSource(s, c)
	return c.get()
}

func thenSource(f SourceInterface, p Pipeline) {
	defer close(p.in())
	for e := range f.out() {
		p.in() <- e
	}
}

func collectSource(f SourceInterface, c Collector) {
	defer close(c.in())
	for e := range f.out() {
		c.in() <- e
	}
}
