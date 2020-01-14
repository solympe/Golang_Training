package main

import (
	"fmt"

	rs "github.com/solympe/Golang_Training/pkg/range-sum-bst"
)

func main() {
	trr := rs.NewTreeNodeRanger(nil, nil, 40)
	tr := rs.NewTreeNodeRanger(nil, trr, 15)
	tll := rs.NewTreeNodeRanger(nil, nil, 5)
	tl := rs.NewTreeNodeRanger(tll, nil, 6)
	t := rs.NewTreeNodeRanger(tl, tr, 10)

	result := t.RangeSum(t, 5, 10)
	fmt.Println("Answer:", result)
}
