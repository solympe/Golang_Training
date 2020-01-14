package add_two_numbers

// ListNodeExecutor ...
type ListNodeExecutor interface {
	AddTwoNumbers(l1 ListNodeExecutor, l2 ListNodeExecutor) ListNodeExecutor
	GetVal() int
	GetNext() ListNodeExecutor
	sumVal(val int)
	setVal(val int)
	setNext(next ListNodeExecutor)
}
