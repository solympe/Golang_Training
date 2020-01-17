package main

import (
	"fmt"

	list "github.com/solympe/Golang_Training/pkg/reverse-list"
)

func main() {
	node5 := list.NewListNode(5, nil)
	node4 := list.NewListNode(4, node5)
	node3 := list.NewListNode(3, node4)
	node2 := list.NewListNode(2, node3)
	node1 := list.NewListNode(1, node2)

	result := node1.Reverse()
	for {
		if result == nil {
			break
		}
		fmt.Println(result)
		result = result.Next()
	}
}
