package rangeSum

import (
	"github.com/solympe/Golang_Training/pkg/rangeSumBST/bTree"
)

// RangeSum counts sum of tree for each element that >= L && <= R
func RangeSum(root *bTree.BTree, L int, R int) (sumBST int) {
	val := bTree.BTree.GetVal(*root)

	if root == nil {
		return sumBST
	}
	if val >= L && val <= R {
		sumBST += val
	}
	if val > L && bTree.BTree.ShowLeft(*root) != nil {
		sumBST += RangeSum(bTree.BTree.ShowLeft(*root), L, R)
	}
	if val < R && bTree.BTree.ShowRight(*root) != nil {
		sumBST += RangeSum(bTree.BTree.ShowRight(*root), L, R)
	}
	return sumBST
}
