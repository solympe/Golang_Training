package main

import (
	"fmt"

	list "github.com/solympe/Golang_Training/pkg/reverse-list"
)

func main() {
	n5 := list.NewListNode(5, nil)
	n4 := list.NewListNode(4, n5)
	n3 := list.NewListNode(3, n4)
	n2 := list.NewListNode(2, n3)
	n1 := list.NewListNode(1, n2)

	result := n1.ReverseList(n1)
	for {
		if result == nil {
			break
		}
		fmt.Println(result)
		result = result.GetNext()
	}
}
