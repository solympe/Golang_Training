package main

import (
	"fmt"

	addn "github.com/solympe/Golang_Training/pkg/add-two-numbers/add-numbers"
	ln "github.com/solympe/Golang_Training/pkg/add-two-numbers/list-node"
)

func main() {
	list6 := ln.NewListNodeExecutor(5, nil)
	list5 := ln.NewListNodeExecutor(5, list6)
	list4 := ln.NewListNodeExecutor(5, list5)

	list3 := ln.NewListNodeExecutor(5, nil)
	list2 := ln.NewListNodeExecutor(5, list3)
	list1 := ln.NewListNodeExecutor(5, list2)

	res := addn.AddTwoNumbers(list1, list4)
	for {
		fmt.Println(res.GetVal())
		if res.GetNext() != nil {
			res = res.GetNext()
		} else {
			break
		}
	}
}
