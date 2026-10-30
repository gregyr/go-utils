package data

import (
	"fmt"
)

var (
	StackOverflowErr  = fmt.Errorf("Can not push to full Stack")
	StackUnderflowErr = fmt.Errorf("Can not pop from empty Stack")
)

type Stack[T any] struct {
	p   int
	arr []T
}

func NewStack[T any](bufferSize int) *Stack[T] {
	return &Stack[T]{p: -1, arr: make([]T, bufferSize)}
}

func (s *Stack[T]) Push(elem T) error {
	if s.p == len(s.arr)-1 {
		return StackOverflowErr
	}
	s.p++
	s.arr[s.p] = elem
	return nil
}

func (s *Stack[T]) Pop() (T, error) {
	if s.p < 0 {
		return *new(T), StackUnderflowErr
	}
	elem := s.arr[s.p]
	s.p--
	return elem, nil
}

func (s *Stack[T]) IsEmpty() bool {
	return s.p == -1
}

func (s *Stack[T]) Elems() []T {
	return s.arr[:s.p+1]
}
