package data

import "fmt"

type RingQueue[T any] struct {
	arr []T
	// points to the oldest element in the queue
	head int
	// points to the last open slot in the array
	tail int
	// keeps track of the amount of elements in the queue
	elements int
}

var (
	QueueFullErr  = fmt.Errorf("cannot enqueue to full queue")
	QueueEmptyErr = fmt.Errorf("cannot dequeue from empty queue")
)

func NewRingQueue[T any](bufferSize int) *RingQueue[T] {
	return &RingQueue[T]{arr: make([]T, bufferSize), head: 0, tail: 0}
}

func (q *RingQueue[T]) Enqueue(elem T) error {
	if q.head == q.tail && q.elements == len(q.arr) {
		return QueueFullErr
	}
	q.arr[q.tail] = elem
	q.tail = (q.tail + 1) % len(q.arr)
	q.elements++
	return nil
}

func (q *RingQueue[T]) Dequeue() (T, error) {
	if q.head == q.tail && q.elements == 0 {
		return *new(T), QueueEmptyErr
	}
	elem := q.arr[q.head]
	q.head = (q.head + 1) % len(q.arr)
	q.elements--
	return elem, nil
}

func (q *RingQueue[T]) Elems() []T {
	return q.arr[q.head : q.tail+1]
}
