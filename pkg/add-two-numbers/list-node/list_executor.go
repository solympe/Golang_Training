package list_node

// ListNodeExecutor ...
type ListNodeExecutor interface {
	SumVal(val int)
	GetVal() int
	SetVal(val int)
	GetNext() ListNodeExecutor
	SetNext(next ListNodeExecutor)
}
