package has_cycle

// ListNodeSolver ...
type ListNodeSolver interface {
	GetNext() ListNodeSolver
	GetVal() int
}
