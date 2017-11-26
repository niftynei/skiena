package five

type lnode struct {
	val int
	next *lnode
}

type Stack struct {
	Next *lnode
}

func (s *Stack) Push(val int) {
	node := &lnode{val:val}

	if s.Next == nil {
		s.Next = node
	} else {
		node.next = s.Next
		s.Next = node
	}
}

func (s *Stack) Pop() int {
	if s.Next == nil {
		return -1 // or panic?
	}
	node := s.Next
	s.Next = node.next
	return node.val
}

func (s *Stack) Empty() bool {
	return s.Next == nil
}
