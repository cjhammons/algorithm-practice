package linkedlist

import "testing"

func TestAppend(t *testing.T) {
	list := &DoublyLinkedList{}
	list.Append(1)
	list.Append(2)
	list.Append(3)

	if list.Head.Value != 1 || list.Tail.Value != 3 {
		t.Errorf("Append did not work as expected.")
	}
}

func TestPrepend(t *testing.T) {
	list := &DoublyLinkedList{}
	list.Prepend(1)
	list.Prepend(2)

	if list.Head.Value != 2 || list.Tail.Value != 1 {
		t.Errorf("Prepend did not work as expected.")
	}
}

func TestDelete(t *testing.T) {
	list := &DoublyLinkedList{}
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Delete(2)

	if list.Head.Next.Value != 3 {
		t.Errorf("Delete did not work as expected.")
	}
}

func TestLength(t *testing.T) {
	list := &DoublyLinkedList{}
	list.Append(1)
	list.Append(2)

	if list.length() != 2 {
		t.Errorf("Length did not return the expected value.")
	}
}

func TestInsertAt(t *testing.T) {
	list := &DoublyLinkedList{}
	list.Append(1)
	list.Append(3)
	list.InsertAt(2, 1)

	if list.Head.Next.Value != 2 {
		t.Errorf("InsertAt did not work as expected.")
	}
}

func TestDeleteAt(t *testing.T) {
	list := &DoublyLinkedList{}
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.DeleteAt(1)

	if list.Head.Next.Value != 3 {
		t.Errorf("DeleteAt did not work as expected.")
	}
}
