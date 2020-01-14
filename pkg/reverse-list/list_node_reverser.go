package reverse_list

// ListNodeReverser ...
type ListNodeReverser interface {
	ReverseList(head ListNodeReverser) ListNodeReverser
	getNext() ListNodeReverser
	setNext(next ListNodeReverser)
}
