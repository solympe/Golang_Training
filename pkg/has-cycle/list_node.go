package has_cycle

type listNode struct {
	val  int
	next ListNodeSolver
}

// GetNext ...
func (l *listNode) GetNext() ListNodeSolver {
	return l.next
}

// GetVal ...
func (l *listNode) GetVal() int {
	return l.val
}

// HasCycle ...
func HasCycle(head ListNodeSolver) bool {
	m := map[ListNodeSolver]int{}
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

// NewListNodeSolver ...
func NewListNodeSolver(value int, next ListNodeSolver) ListNodeSolver {
	return &listNode{
		val:  value,
		next: next,
	}
}
