package queue

import (
	"testing"
)

func TestNewQueue(t *testing.T) {
	q := NewQueue()
	if q == nil {
		t.Error("New queue should not be nil")
	}
}

func TestEnqueue(t *testing.T) {
	q := NewQueue()
	q.Enqueue(1)

	// Assuming a method Size() exists to check the size of the queue
	if q.Size() != 1 {
		t.Errorf("Expected queue size to be 1, got %d", q.Size())
	}

	// Assuming a method Peek() exists to check the item
	if q.Peek() != 1 {
		t.Errorf("Expected Peek to return 1, got %v", q.Peek())
	}
}

func TestDequeue(t *testing.T) {
	q := NewQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	item := q.Dequeue()

	if item != 1 {
		t.Errorf("Expected Dequeue to return 1, got %v", item)
	}

	if q.Size() != 1 {
		t.Errorf("Expected queue size to be 1, got %d", q.Size())
	}

	if q.Peek() != 2 {
		t.Errorf("Expected Peek to return 2, got %v", q.Peek())
	}
}

func TestPeek(t *testing.T) {
	q := NewQueue()
	q.Enqueue(1)
	item := q.Peek()

	if item != 1 {
		t.Errorf("Expected Peek to return 1, got %v", item)
	}

	if q.Size() != 1 {
		t.Errorf("Expected queue size to be 1, got %d", q.Size())
	}
}
