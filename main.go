package main

import (
	"fmt"
	rs "github.com/solympe/Golang_Training/leetcode/rangeSumBST"
)

func main() {

	tr1 := rs.TreeNode{
		Val:   32,
		Left:  nil,
		Right: nil,
	}

	tl := rs.TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}

	tr := rs.TreeNode{
		Val:   15,
		Left:  nil,
		Right: &tr1,
	}

	t0 := rs.TreeNode{
		Val:   10,
		Left:  &tl,
		Right: &tr,
	}

	result := rs.RangeSumBST(&t0, 6, 10)

	fmt.Println("Answer:", result)

}
