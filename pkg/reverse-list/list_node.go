package listnode

// ListNode ...
type ListNode interface {
	ReverseList() ListNode
	GetNext() ListNode
	setNext(next ListNode)
}

type listNode struct {
	val  int
	next ListNode
}

// ReverseList return reversed list
func (l *listNode) ReverseList() ListNode {
	head := l
	if head == nil || head.next == nil {
		return head
	}
	revPoint := head.next.ReverseList()
	head.next.setNext(head)
	head.next = nil
	return revPoint
}

// GetNext ...
func (l *listNode) GetNext() ListNode {
	return l.next
}

func (l *listNode) setNext(next ListNode) {
	l.next = next
}

// NewListNode returns new element of list
func NewListNode(val int, next ListNode) ListNode {
	return &listNode{val, next}
}
