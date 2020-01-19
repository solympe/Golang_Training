package main

import (
	"fmt"

	list "github.com/solympe/Golang_Training/pkg/add-two-numbers"
)

func main() {
	list6 := list.NewListNode(5, nil)
	list5 := list.NewListNode(5, list6)
	list4 := list.NewListNode(5, list5)

	list3 := list.NewListNode(5, nil)
	list2 := list.NewListNode(5, list3)
	list1 := list.NewListNode(5, list2)

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
