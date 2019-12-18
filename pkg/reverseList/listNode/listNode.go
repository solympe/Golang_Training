package listNode

// List represents interface of list
type List interface {
	ReverseList(head *listNode) *listNode
}

// listNode represents list struct
type listNode struct {
	Val  int
	Next *listNode
}

// ReverseList return
func ReverseList(head *listNode) *listNode {
	if head == nil || head.Next == nil {
		return head
	}
	revPoint := ReverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return revPoint
}

// NewNode returns new element of list
func NewNode(val int, next *listNode) *listNode {
	return &listNode{val, next}
}
