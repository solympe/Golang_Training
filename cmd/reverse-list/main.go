package main

import (
	"fmt"

	lf "github.com/solympe/Golang_Training/pkg/reverse-list"
)

func main() {
	n5 := lf.NewListNodeReverser(5, nil)
	n4 := lf.NewListNodeReverser(4, n5)
	n3 := lf.NewListNodeReverser(3, n4)
	n2 := lf.NewListNodeReverser(2, n3)
	n1 := lf.NewListNodeReverser(1, n2)

	fmt.Println(n1.ReverseList(n1))
}
