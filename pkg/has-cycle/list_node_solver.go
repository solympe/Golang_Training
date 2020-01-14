package has_cycle

// ListNodeSolver ...
type ListNodeSolver interface {
	HasCycle(head ListNodeSolver) bool
	GetNext() ListNodeSolver
	GetVal() int
}
