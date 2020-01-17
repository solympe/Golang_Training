package main

import (
	"fmt"

	list "github.com/solympe/Golang_Training/pkg/reverse-list"
)

func main() {
	node5 := list.NewNode(5, nil)
	node4 := list.NewNode(4, node5)
	node3 := list.NewNode(3, node4)
	node2 := list.NewNode(2, node3)
	node1 := list.NewNode(1, node2)

	result := node1.Reverse()
	for result != nil {
		fmt.Println(result)
		result = result.Next()
	}
}
