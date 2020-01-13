package list_node

type listNode struct {
	val  int
	next ListNodeExecutor
}

// SetVal ..
func (l *listNode) SetVal(val int) {
	l.val = val
}

// SumVal ...
func (l *listNode) SumVal(val int) {
	l.val += val
}

// GetVal ...
func (l *listNode) GetVal() int {
	return l.val
}

// GetNext ...
func (l *listNode) GetNext() ListNodeExecutor {
	return l.next
}

// SetNext ...
func (l *listNode) SetNext(next ListNodeExecutor) {
	l.next = next
}

// NewListNodeExecutor returns new copy of node
func NewListNodeExecutor(value int, next ListNodeExecutor) ListNodeExecutor {
	return &listNode{value, next}
}
