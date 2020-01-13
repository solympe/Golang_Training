package main

import (
	"fmt"

	mt "github.com/solympe/Golang_Training/pkg/merge-trees"
)

func main() {
	trr := mt.NewTreeMerger(100, nil, nil)

	bl := mt.NewTreeMerger(1, nil, nil)

	br := mt.NewTreeMerger(4, nil, nil)
	tr := mt.NewTreeMerger(3, nil, trr)

	b0 := mt.NewTreeMerger(2, bl, br)
	t0 := mt.NewTreeMerger(4, nil, tr)

	fmt.Println(mt.MergeTrees(t0, b0).GetVal(), "is a value of head")
}
