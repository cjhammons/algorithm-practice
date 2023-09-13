package stack

type SNode struct {
	Value interface{}
	Next  *SNode
}

type Stack struct {
	Head   *SNode
	Length int
}

func NewStack() *Stack {
	return &Stack{
		Head:   nil,
		Length: 0,
	}
}

func (s *Stack) Push(item interface{}) {
	newNode := &SNode{
		Value: item,
		Next:  nil,
	}

	if s.Head == nil {
		s.Head = newNode
	} else {
		newNode.Next = s.Head
		s.Head = newNode
	}
	s.Length++
}

func (s *Stack) Pop() interface{} {
	popped := s.Head.Value
	s.Head = s.Head.Next
	s.Length--
	return popped
}

func (s *Stack) Peak() interface{} {
	return s.Head.Value
}

func (s *Stack) Size() int {
	return s.Length
}
