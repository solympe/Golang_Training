package main

import (
	"fmt"

	lf "github.com/solympe/Golang_Training/pkg/reverse-list"
)

func main() {
	n5 := lf.NewListNode(5, nil)
	n4 := lf.NewListNode(4, n5)
	n3 := lf.NewListNode(3, n4)
	n2 := lf.NewListNode(2, n3)
	n1 := lf.NewListNode(1, n2)

	result := n1.ReverseList(n1)
	for {
		if result == nil {
			break
		}
		fmt.Println(result)
		result = result.GetNext()
	}
}
