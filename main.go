package main

import (
	"fmt"
	"github.com/solympe/Golang_Training/mergeTrees"
)

func main() {

	br := mergeTrees.TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}

	bl := mergeTrees.TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}

	b0 := mergeTrees.TreeNode{
		Val:   2,
		Left:  &bl,
		Right: &br,
	}

	trr := mergeTrees.TreeNode{
		Val:   100,
		Left:  nil,
		Right: nil,
	}

	tr := mergeTrees.TreeNode{
		Val:   3,
		Left:  nil,
		Right: &trr,
	}

	t0 := mergeTrees.TreeNode{
		Val:   4,
		Left:  nil,
		Right: &tr,
	}

	fmt.Println("res: ", mergeTrees.MergeTrees(&t0, &b0), "is a head of tree")

}
