package rangeSum

import bt "github.com/solympe/Golang_Training/leetcode/rangeSumBST/bTree"

//RangeSum counts sum of tree for each element that >= L && <= R
func RangeSum(root *bt.BTree, L int, R int) (sumBST int) {
	val := bt.BTree.GetVal(*root)

	if root == nil {
		return sumBST
	}
	if val >= L && val <= R {
		sumBST += val
	}
	if val > L && bt.BTree.ShowLeft(*root) != nil {
		sumBST += RangeSum(bt.BTree.ShowLeft(*root), L, R)
	}
	if val < R && bt.BTree.ShowRight(*root) != nil {
		sumBST += RangeSum(bt.BTree.ShowRight(*root), L, R)
	}
	return sumBST
}
