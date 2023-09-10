package linkedlist

import "testing"

func TestAppend(t *testing.T) {
	list := &DoublyLinkedList{}

	list.Append(1)
	list.Append(2)
	list.Append(3)

	if list.Head.Value != 1 {
		t.Errorf("Head node value should be 1, got %d", list.Head.Value)
	}

	if list.Tail.Value != 3 {
		t.Errorf("Tail node value should be 3, got %d", list.Tail.Value)
	}

	if list.Head.Next.Value != 2 {
		t.Errorf("Second node value should be 2, got %d", list.Head.Next.Value)
	}

	if list.Head.Next.Prev.Value != 1 {
		t.Errorf("Second node's previous should point back to head, got %d", list.Head.Next.Prev.Value)
	}

	if list.Tail.Prev.Value != 2 {
		t.Errorf("Tail node's previous should be 2, got %d", list.Tail.Prev.Value)
	}
}

func TestDelete(t *testing.T) {
	list := &DoublyLinkedList{}
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)

	if !list.Delete(2) {
		t.Error("Delete should return true for a value that exists in the list")
	}

	if list.Head.Value != 1 {
		t.Errorf("Head node value should be 1, got %d", list.Head.Value)
	}

	if list.Head.Next.Value != 3 {
		t.Errorf("Second node value should be 3, got %d", list.Head.Next.Value)
	}

	if list.Head.Next.Prev.Value != 1 {
		t.Errorf("Second node's previous should point back to head, got %d", list.Head.Next.Prev.Value)
	}

	if list.Tail.Value != 4 {
		t.Errorf("Tail node value should be 4, got %d", list.Tail.Value)
	}

	if !list.Delete(4) {
		t.Error("Delete should return true for a value that exists in the list")
	}

	if list.Tail.Value != 3 {
		t.Errorf("Tail node value should be 3, got %d", list.Tail.Value)
	}

	if list.Delete(10) {
		t.Error("Delete should return false for a value that does not exist in the list")
	}

	if !list.Delete(1) {
		t.Error("Delete should return true for a value that exists in the list")
	}

}
