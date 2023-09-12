package linkedlist

type Node struct {
	Value interface{}
	Next  *Node
	Prev  *Node
}

type DoublyLinkedList struct {
	Head *Node
	Tail *Node
}

func (l *DoublyLinkedList) Append(value int) {
	newNode := &Node{Value: value}

	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
		return
	}

	l.Tail.Next = newNode
	newNode.Prev = l.Tail
	l.Tail = newNode

}

func (l *DoublyLinkedList) Prepend(value int) {
	newNode := &Node{Value: value}

	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
		return
	}

	l.Head.Prev = newNode
	newNode.Next = l.Head
	l.Head = newNode
}

func (l *DoublyLinkedList) Delete(value int) {
	current := l.Head
	for current != nil {
		if current.Value == value {
			if current.Prev != nil {
				current.Prev.Next = current.Next
			} else {
				l.Head = current.Next
			}

			if current.Next != nil {
				current.Next.Prev = current.Prev
			} else {
				l.Tail = current.Prev
			}

			return
		}
		current = current.Next
	}
}

func (l *DoublyLinkedList) length() int {
	count := 0
	current := l.Head
	for current != nil {
		count++
		current = current.Next
	}
	return count
}

func (l *DoublyLinkedList) InsertAt(val int, index int) {
	i := 0
	n := l.Head
	for i != index {
		if n.Next == nil {
			return
		}
		n = n.Next
		i++
	}
	newNode := &Node{Value: val}

	n.Prev.Next = newNode
	newNode.Next = n
	newNode.Prev = n.Prev
}

func (l *DoublyLinkedList) DeleteAt(index int) {
	i := 0
	current := l.Head
	for current != nil {
		if index == i {
			if current.Prev != nil {
				current.Prev.Next = current.Next
			} else {
				l.Head = current.Next
			}

			if current.Next != nil {
				current.Next.Prev = current.Prev
			} else {
				l.Tail = current.Prev
			}

			return
		}
		current = current.Next
		i++
	}
}
