package add_two_numbers

type listNode struct {
	val  int
	next ListNodeExecutor
}

// GetVal ...
func (l *listNode) GetVal() int {
	return l.val
}

// GetNext ...
func (l *listNode) GetNext() ListNodeExecutor {
	return l.next
}

// AddTwoNumbers sums 2 lists and returns new list
func (l *listNode) AddTwoNumbers(l1 ListNodeExecutor, l2 ListNodeExecutor) ListNodeExecutor {
	overflow := 0
	result := NewListNodeExecutor(0, nil)
	return l.makeSum(l1, l2, result, overflow)
}

func (l *listNode) makeSum(l1 ListNodeExecutor, l2 ListNodeExecutor, node ListNodeExecutor, overflow int) ListNodeExecutor {
	var next1 ListNodeExecutor = nil
	var next2 ListNodeExecutor = nil

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
		node.setNext(NewListNodeExecutor(0, nil))
		l.makeSum(next1, next2, node.GetNext(), overflow)
	}
	return node
}

func (l *listNode) setNext(next ListNodeExecutor) {
	l.next = next
}

func (l *listNode) setVal(val int) {
	l.val = val
}

func (l *listNode) sumVal(val int) {
	l.val += val
}

// NewListNodeExecutor returns new copy of node
func NewListNodeExecutor(value int, next ListNodeExecutor) ListNodeExecutor {
	return &listNode{value, next}
}
