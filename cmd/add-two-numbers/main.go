package main

import (
	"fmt"

	tree "github.com/solympe/Golang_Training/pkg/add-two-numbers"
)

func main() {
	list6 := tree.NewListNode(5, nil)
	list5 := tree.NewListNode(5, list6)
	list4 := tree.NewListNode(5, list5)

	list3 := tree.NewListNode(5, nil)
	list2 := tree.NewListNode(5, list3)
	list1 := tree.NewListNode(5, list2)

	res := list1.AddTwoNumbers(list1, list4)
	for {
		fmt.Println(res.GetVal())
		if res.GetNext() != nil {
			res = res.GetNext()
		} else {
			break
		}
	}
}
