package queue

type QNode struct {
	Value interface{}
	Next  *QNode
}

type Queue struct {
	Head   *QNode
	Tail   *QNode
	Length int
}

func NewQueue() *Queue {
	return &Queue{
		Head:   nil,
		Tail:   nil,
		Length: 0,
	}
}

func (q *Queue) Enqueue(item interface{}) {
	q.Length++
	newNode := &QNode{
		Value: item,
		Next:  nil,
	}
	if q.Tail == nil {
		q.Head = newNode
		q.Tail = newNode
	} else {
		q.Tail.Next = newNode
		q.Tail = newNode
	}
}

func (q *Queue) Dequeue() interface{} {
	if q.Head == nil {
		return nil
	}
	q.Length--
	head := q.Head
	q.Head = q.Head.Next
	head.Next = nil
	return head.Value
}

func (q *Queue) Peek() interface{} {
	if q.Head != nil {
		return q.Head.Value
	}
	return nil
}

func (q *Queue) Size() int {
	return q.Length
}
