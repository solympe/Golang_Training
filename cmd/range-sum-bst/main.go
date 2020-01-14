package main

import (
	"fmt"

	bt "github.com/solympe/Golang_Training/pkg/range-sum-bst/b-tree"
	rs "github.com/solympe/Golang_Training/pkg/range-sum-bst/range-sum"
)

func main() {
	trr := bt.NewNode(nil, nil, 40)
	tr := bt.NewNode(nil, &trr, 15)
	tll := bt.NewNode(nil, nil, 5)
	tl := bt.NewNode(&tll, nil, 6)
	t := bt.NewNode(&tl, &tr, 10)

	result := rs.RangeSum(&t, 5, 10)
	fmt.Println("Answer:", result)
}
