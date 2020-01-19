package main

import (
	"fmt"

	list "github.com/solympe/Golang_Training/pkg/leetcode_practice/has-cycle"
)

func main() {

	node2 := list.NewListNode(10, nil)
	node1 := list.NewListNode(11, node2)

	fmt.Println(node1.HasCycle(node1))
}
