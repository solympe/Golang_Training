package range_sum

import (
	"github.com/solympe/Golang_Training/pkg/range-sum-bst/b-tree"
)

// RangeSum counts sum of tree for each element that >= L && <= R
func RangeSum(root *b_tree.BTree, L int, R int) (sumBST int) {
	val := b_tree.BTree.GetVal(*root)

	if root == nil {
		return sumBST
	}
	if val >= L && val <= R {
		sumBST += val
	}
	if val > L && b_tree.BTree.ShowLeft(*root) != nil {
		sumBST += RangeSum(b_tree.BTree.ShowLeft(*root), L, R)
	}
	if val < R && b_tree.BTree.ShowRight(*root) != nil {
		sumBST += RangeSum(b_tree.BTree.ShowRight(*root), L, R)
	}
	return sumBST
}
