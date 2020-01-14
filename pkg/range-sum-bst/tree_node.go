package range_sum_bst

// treeNode is a definition for a binary tree node
type treeNode struct {
	val   int
	left  TreeNodeRanger
	right TreeNodeRanger
}

// GetVal returns value of node
func (t *treeNode) GetVal() int {
	return t.val
}

// ShowLeft returns pointer to left node
func (t *treeNode) ShowLeft() TreeNodeRanger {
	return t.left
}

// ShowRight returns pointer to right node
func (t *treeNode) ShowRight() TreeNodeRanger {
	return t.right
}

// RangeSum counts sum of tree for each element that >= L && <= R
func (t *treeNode) RangeSum(root TreeNodeRanger, L int, R int) (sumBST int) {
	val := root.GetVal()

	if root == nil {
		return sumBST
	}
	if val >= L && val <= R {
		sumBST += val
	}
	if val > L && root.ShowLeft() != nil {
		sumBST += t.RangeSum(root.ShowLeft(), L, R)
	}
	if val < R && root.ShowRight() != nil {
		sumBST += t.RangeSum(root.ShowRight(), L, R)
	}
	return sumBST
}

//  NewTreeNodeRanger returns new tree node with input parameters
func NewTreeNodeRanger(tleft TreeNodeRanger, tright TreeNodeRanger, val int) TreeNodeRanger {
	return &treeNode{
		val:   val,
		left:  tleft,
		right: tright,
	}
}
