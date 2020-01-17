package listnode

// ListNode ...
type ListNode interface {
	Reverse() ListNode
	Next() ListNode
	set(next ListNode)
}

type listNode struct {
	val  int
	next ListNode
}

// Reverse return reversed list
func (node *listNode) Reverse() ListNode {
	if node == nil || node.next == nil {
		return node
	}
	revPoint := node.next.Reverse()
	node.next.set(node)
	node.next = nil
	return revPoint
}

// Next ...
func (node *listNode) Next() ListNode {
	return node.next
}

func (node *listNode) set(next ListNode) {
	node.next = next
}

// NewListNode returns new element of list
func NewListNode(val int, next ListNode) ListNode {
	return &listNode{val, next}
}
