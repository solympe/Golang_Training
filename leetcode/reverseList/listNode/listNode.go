package listNode

type ListFunctions interface {
	GetNext() ListFunctions
	NillNext()
}

// listNode is a structure of node
type listNode struct {
	Val  int
	Next *ListFunctions
}

// GetNext returns next element of list
func (l *listNode) GetNext() ListFunctions {
	return *l.Next
}

func (l *listNode) NillNext() {
	l.Next = nil
}

//
func AfterNext(head *ListFunctions) {
//	head.
}

// NewNode returns new element of list
func NewNode(val int, next *ListFunctions) ListFunctions {
	return &listNode{val, next}
}
