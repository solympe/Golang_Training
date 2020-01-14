package main

import (
	"fmt"

	hc "github.com/solympe/Golang_Training/pkg/has-cycle"
)

func main() {
	var node1 hc.ListNodeSolver
	node2 := hc.NewListNodeSolver(10, nil)
	node1 = hc.NewListNodeSolver(11, node2)

	fmt.Println(node1.HasCycle(node1))
}
