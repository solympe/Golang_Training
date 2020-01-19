package listnode

// ListNode ...
type ListNode interface {
	AddTwoNumbers(l1 ListNode, l2 ListNode) ListNode
	GetVal() int
	GetNext() ListNode
	sumVal(val int)
	setVal(val int)
	setNext(next ListNode)
}

type listNode struct {
	val  int
	next ListNode
}

// GetVal ...
func (l *listNode) GetVal() int {
	return l.val
}

// GetNext ...
func (l *listNode) GetNext() ListNode {
	return l.next
}

// AddTwoNumbers sums 2 lists and returns new list
func (l *listNode) AddTwoNumbers(l1 ListNode, l2 ListNode) ListNode {
	overflow := 0
	result := NewListNode(0, nil)
	return l.makeSum(l1, l2, result, overflow)
}

func (l *listNode) makeSum(l1 ListNode, l2 ListNode, node ListNode, overflow int) ListNode {
	var next1 ListNode = nil
	var next2 ListNode = nil

	node.setVal(node.GetVal() + overflow)
	if l1 != nil {
		node.sumVal(l1.GetVal())
		next1 = l1.GetNext()
	}
	if l2 != nil {
		node.sumVal(l2.GetVal())
		next2 = l2.GetNext()
	}
	if node.GetVal() >= 10 {

		overflow = node.GetVal() / 10

		node.setVal(node.GetVal() % 10)
	} else {
		overflow = 0
	}
	if next1 != nil || next2 != nil || overflow > 0 {
		node.setNext(NewListNode(0, nil))
		l.makeSum(next1, next2, node.GetNext(), overflow)
	}
	return node
}

func (l *listNode) setNext(next ListNode) {
	l.next = next
}

func (l *listNode) setVal(val int) {
	l.val = val
}

func (l *listNode) sumVal(val int) {
	l.val += val
}

// NewListNode returns new copy of node
func NewListNode(value int, next ListNode) ListNode {
	return &listNode{value, next}
}
