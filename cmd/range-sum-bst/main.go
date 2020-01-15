package main

import (
	"fmt"

	tree "github.com/solympe/Golang_Training/pkg/range-sum-bst"
)

func main() {
	trr := tree.NewTreeNode(nil, nil, 40)
	tr := tree.NewTreeNode(nil, trr, 15)

	tll := tree.NewTreeNode(nil, nil, 5)
	tl := tree.NewTreeNode(tll, nil, 6)

	t0 := tree.NewTreeNode(tl, tr, 10)

	result := t0.RangeSum(t0, 5, 10)
	fmt.Println("Answer:", result)
}
