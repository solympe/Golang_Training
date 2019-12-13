package main

import (
	"fmt"
	rs "github.com/solympe/Golang_Training/leetcode/rangeSumBST"
)

func main() {


	tlr := rs.TreeNode{
		Val:   6,
		Left:  nil,
		Right: nil,
	}

	tll2 := rs.TreeNode{
		Val:   0,
		Left:  &tlr,
		Right: nil,
	}

	tll := rs.TreeNode{
		Val:   18,
		Left:  &tll2,
		Right: nil,
	}
	trr3 := rs.TreeNode{
		Val:   13,
		Left:  &tll,
		Right: nil,
	}
	trr2 := rs.TreeNode{
		Val:   7,
		Left:  nil,
		Right: &trr3,
	}
	trr1 := rs.TreeNode{
		Val:   3,
		Left:  nil,
		Right: &trr2,
	}

	trr := rs.TreeNode{
		Val:   15,
		Left:  nil,
		Right: &trr1,
	}

	tr := rs.TreeNode{
		Val:   5,
		Left:  nil,
		Right: &trr,
	}

	t0 := rs.TreeNode{
		Val:   10,
		Left:  nil,
		Right: &tr,
	}

	result := rs.RangeSumBST(&t0, 6, 10)

	fmt.Println("Answer:", result)

}
