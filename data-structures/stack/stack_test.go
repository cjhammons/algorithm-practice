package stack

import (
	"testing"
)

func TestNewStack(t *testing.T) {
	s := NewStack()
	if s == nil {
		t.Error("New stack should not be nil")
	}

	if s.Size() != 0 {
		t.Errorf("New stack should have size 0, got %d", s.Size())
	}
}

func TestPush(t *testing.T) {
	s := NewStack()
	s.Push(1)

	if s.Size() != 1 {
		t.Errorf("Expected stack size to be 1, got %d", s.Size())
	}

	if s.Peak() != 1 {
		t.Errorf("Expected peak to return 1, got %v", s.Peak())
	}
}

func TestPop(t *testing.T) {
	s := NewStack()
	s.Push(1)
	s.Push(2)
	popped := s.Pop()

	if popped != 2 {
		t.Errorf("Expected Pop to return 2, got %v", popped)
	}

	if s.Size() != 1 {
		t.Errorf("Expected stack size to be 1, got %d", s.Size())
	}

	if s.Peak() != 1 {
		t.Errorf("Expected peak to return 1, got %v", s.Peak())
	}
}

func TestPeak(t *testing.T) {
	s := NewStack()
	s.Push(1)

	if s.Peak() != 1 {
		t.Errorf("Expected peak to return 1, got %v", s.Peak())
	}

	if s.Size() != 1 {
		t.Errorf("Expected stack size to be 1, got %d", s.Size())
	}
}
