package range_sum_bst

// TreeNodeRanger ...
type TreeNodeRanger interface {
	RangeSum(root TreeNodeRanger, L int, R int) (sumBST int)
	getVal() int
	showLeft() TreeNodeRanger
	showRight() TreeNodeRanger
}
