package tracing

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// Span records one traced operation and the operations it contains.
type Span struct {
	Name     string
	Duration time.Duration
	Children []*Span

	collector *Collector
}

// Collector owns the root spans of a trace tree.
type Collector struct {
	mu    sync.RWMutex
	roots []*Span
}

// NewCollector creates a collector for one or more trace trees.
func NewCollector() *Collector {
	return &Collector{}
}

// Trace runs function as a root span and returns its completed span.
func (c *Collector) Trace(name string, function func(*Span)) *Span {
	span := &Span{Name: name, collector: c}

	c.mu.Lock()
	c.roots = append(c.roots, span)
	c.mu.Unlock()

	return span.run(function)
}

// Trace runs function as a child span and returns its completed span.
func (s *Span) Trace(name string, function func(*Span)) *Span {
	child := &Span{Name: name, collector: s.collector}

	s.collector.mu.Lock()
	s.Children = append(s.Children, child)
	s.collector.mu.Unlock()

	return child.run(function)
}

func (s *Span) run(function func(*Span)) *Span {
	start := time.Now()
	defer func() {
		s.collector.mu.Lock()
		s.Duration = time.Since(start)
		s.collector.mu.Unlock()
	}()

	function(s)
	return s
}

// Spans returns a snapshot of the collected trace trees.
func (c *Collector) Spans() []*Span {
	c.mu.RLock()
	defer c.mu.RUnlock()

	spans := make([]*Span, len(c.roots))
	for i, span := range c.roots {
		spans[i] = cloneSpan(span)
	}
	return spans
}

func cloneSpan(span *Span) *Span {
	copy := &Span{
		Name:     span.Name,
		Duration: span.Duration,
		Children: make([]*Span, len(span.Children)),
	}
	for i, child := range span.Children {
		copy.Children[i] = cloneSpan(child)
	}
	return copy
}

// Collect prints all trace trees using indentation to show parent-child relationships.
func Collect(collector *Collector) {
	for _, span := range collector.Spans() {
		collect(span, 0)
	}
}

func collect(span *Span, depth int) {
	fmt.Printf("%s%s in %s\n", strings.Repeat("  ", depth), span.Name, span.Duration)
	for _, child := range span.Children {
		collect(child, depth+1)
	}
}
