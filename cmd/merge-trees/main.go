package main

import (
	"fmt"

	tree "github.com/solympe/Golang_Training/pkg/merge-trees"
)

func main() {
	trr := tree.NewTreeNode(100, nil, nil)

	bl := tree.NewTreeNode(1, nil, nil)

	br := tree.NewTreeNode(4, nil, nil)
	tr := tree.NewTreeNode(3, nil, trr)

	b0 := tree.NewTreeNode(2, bl, br)
	t0 := tree.NewTreeNode(4, nil, tr)

	result := t0.MergeTrees(t0, b0)
	fmt.Println(result.GetVal(), "is a value of head")
}
