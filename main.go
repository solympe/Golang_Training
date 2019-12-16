package main

import (
	"fmt"

	rl "github.com/solympe/Golang_Training/leetcode/reverseList"
	lf "github.com/solympe/Golang_Training/leetcode/reverseList/listNode"
)

func main() {

	n5 := lf.NewNode(5, nil)
	n4 := lf.NewNode(4, &n5)
	n3 := lf.NewNode(3, &n4)
	n2 := lf.NewNode(2, &n3)
	n1 := lf.NewNode(1, &n2)

	fmt.Println(rl.ReverseList(&n1))

}
