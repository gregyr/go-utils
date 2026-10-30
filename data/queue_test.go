package data

import (
	"testing"
)

func TestEnAndDequeueing(t *testing.T) {
	q := NewRingQueue[int](5)

	queuedValue := 1

	err := q.Enqueue(queuedValue)
	if err != nil {
		t.Fatalf("did not expect error: %v", err)
	}
	if q.head != 0 {
		t.Errorf("expected queue head to remain 0, got %v", q.head)
	}
	if q.tail != 1 {
		t.Errorf("expected queue tail to be 1, got %v", q.tail)
	}
	if q.elements != 1 {
		t.Errorf("expected number of elements in queue to be 1, got %v", q.elements)
	}

	val, err2 := q.Dequeue()
	if err2 != nil {
		t.Fatalf("did not expect error: %v", err)
	}
	if val != queuedValue {
		t.Fatalf("expected dequeued value to be %v, got %v", queuedValue, val)
	}
	if q.head != 1 {
		t.Errorf("expected queue head to remain 0, got %v", q.head)
	}
	if q.tail != 1 {
		t.Errorf("expected queue tail to be 1, got %v", q.tail)
	}
	if q.elements != 0 {
		t.Errorf("expected number of elements in queue to be 1, got %v", q.elements)
	}
}

func TestRejectsEnAndDequeingOnFullAndEmptyQueue(t *testing.T) {
	q := NewRingQueue[int](2)

	err := q.Enqueue(1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	err = q.Enqueue(2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expectsEnqueueErr := q.Enqueue(3)
	if expectsEnqueueErr == nil {
		t.Fatalf("expected error but got nil")
	}
	if expectsEnqueueErr != QueueFullErr {
		t.Fatalf("expected QueueFullErr but got: %v", expectsEnqueueErr)
	}

	_, err = q.Dequeue()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	_, err = q.Dequeue()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	_, expectsDequeueErr := q.Dequeue()
	if expectsDequeueErr == nil {
		t.Fatalf("expected error but got nil")
	}
	if expectsDequeueErr != QueueEmptyErr {
		t.Fatalf("expected QueueEmptyErr but got: %v", expectsDequeueErr)
	}
}
