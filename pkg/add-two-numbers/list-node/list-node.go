package list_node

type listNode struct {
	Val  int
	Next *listNode
}

func newNode() *listNode {
	return &listNode{}
}

func makeSum(l1 *listNode, l2 *listNode, node *listNode, overflow int) *listNode {
	var next1 *listNode = nil
	var next2 *listNode = nil
	node.Val += overflow

	if l1 != nil {
		node.Val += l1.Val
		next1 = l1.Next
	}
	if l2 != nil {
		node.Val += l2.Val
		next2 = l2.Next
	}
	if node.Val >= 10 {

		overflow = node.Val / 10

		node.Val %= 10
	} else {
		overflow = 0
	}
	if next1 != nil || next2 != nil || overflow > 0 {
		node.Next = newNode()
		makeSum(next1, next2, node.Next, overflow)
	}
	return node
}

// AddTwoNumbers sums 2 lists and return a new list
func AddTwoNumbers(l1 *listNode, l2 *listNode) *listNode {
	overflow := 0
	result := newNode()
	return makeSum(l1, l2, result, overflow)
}

// NewlistNode returns new copy of node
func NewlistNode(value int, next *listNode) *listNode {
	return &listNode{value, next}
}
