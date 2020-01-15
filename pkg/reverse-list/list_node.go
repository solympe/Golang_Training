package listnode

// ListNode ...
type ListNode interface {
	ReverseList(head ListNode) ListNode
	GetNext() ListNode
	setNext(next ListNode)
}

type listNode struct {
	val  int
	next ListNode
}

// ReverseList return reversed list
func (l *listNode) ReverseList(head ListNode) ListNode {
	if head == nil || head.GetNext() == nil {
		return head
	}
	revPoint := head.ReverseList(head.GetNext())
	head.GetNext().setNext(head)
	head.setNext(nil)
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
