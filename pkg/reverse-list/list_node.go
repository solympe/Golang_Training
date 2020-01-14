package reverse_list

type listNode struct {
	val  int
	next ListNodeReverser
}

// ReverseList return reversed list
func (l *listNode) ReverseList(head ListNodeReverser) ListNodeReverser {
	if head == nil || head.getNext() == nil {
		return head
	}
	revPoint := head.ReverseList(head.getNext())
	head.getNext().setNext(head)
	head.setNext(nil)
	return revPoint
}

func (l *listNode) setNext(next ListNodeReverser) {
	l.next = next
}

func (l *listNode) getNext() ListNodeReverser {
	return l.next
}

// NewListNodeReverser returns new element of list
func NewListNodeReverser(val int, next ListNodeReverser) ListNodeReverser {
	return &listNode{val, next}
}
