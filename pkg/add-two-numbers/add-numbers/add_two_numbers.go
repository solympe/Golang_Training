package add_numbers

import ln "github.com/solympe/Golang_Training/pkg/add-two-numbers/list-node"

// AddTwoNumbers sums 2 lists and return a new list
func AddTwoNumbers(l1 ln.ListNodeExecutor, l2 ln.ListNodeExecutor) ln.ListNodeExecutor {
	overflow := 0
	result := ln.NewListNodeExecutor(0, nil)
	return makeSum(l1, l2, result, overflow)
}

func makeSum(l1 ln.ListNodeExecutor, l2 ln.ListNodeExecutor, node ln.ListNodeExecutor, overflow int) ln.ListNodeExecutor {
	var next1 ln.ListNodeExecutor = nil
	var next2 ln.ListNodeExecutor = nil

	node.SetVal(node.GetVal() + overflow)

	if l1 != nil {
		node.SumVal(l1.GetVal())
		next1 = l1.GetNext()
	}
	if l2 != nil {
		node.SumVal(l2.GetVal())
		next2 = l2.GetNext()
	}
	if node.GetVal() >= 10 {

		overflow = node.GetVal() / 10

		node.SetVal(node.GetVal() % 10)
	} else {
		overflow = 0
	}
	if next1 != nil || next2 != nil || overflow > 0 {
		node.SetNext(ln.NewListNodeExecutor(0, nil))
		makeSum(next1, next2, node.GetNext(), overflow)
	}
	return node
}
