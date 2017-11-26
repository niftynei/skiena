package five


// implement a queue using a doubly linked list
type dnode struct {
	next *dnode
	prev *dnode
	val int // hi, being lazy
}

type Queue struct {
	Head *dnode
	Tail *dnode
}

// add to the end
func (q *Queue) Enqueue(val int) {
	node := &dnode{val:val}
	if q.Tail != nil {
		node.prev = q.Tail
	}

	q.Tail = node
	if q.Head == nil {
		q.Head = node
	}
}

// remove from the front
func (q *Queue) Dequeue() int {
	if q.Head == nil {
		return 0
	}

	head := q.Head
	if q.Head.next != nil {
		q.Head = q.Head.next
	} else {
		q.Head = nil
		q.Tail = nil
	}

	return head.val
}

func (q *Queue) Empty() bool {
	return q.Head == nil
}
