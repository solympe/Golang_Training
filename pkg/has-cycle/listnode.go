package listnode

// ListNode ...
type ListNode interface {
	HasCycle(head ListNode) bool
	GetNext() ListNode
	GetVal() int
}

type listNode struct {
	val  int
	next ListNode
}

// GetNext ...
func (l *listNode) GetNext() ListNode {
	return l.next
}

// GetVal ...
func (l *listNode) GetVal() int {
	return l.val
}

// HasCycle ...
func (l *listNode) HasCycle(head ListNode) bool {
	m := map[ListNode]int{}
	for head != nil && head.GetNext() != nil {
		_, ok := m[head.GetNext()]
		if ok == true {
			return true
		}
		m[head.GetNext()] = head.GetVal()
		head = head.GetNext()
	}
	return false
}

// NewListNode ...
func NewListNode(value int, next ListNode) ListNode {
	return &listNode{
		val:  value,
		next: next,
	}
}
