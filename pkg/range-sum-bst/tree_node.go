package range_sum_bst

// treeNode is a definition for a binary tree node
type treeNode struct {
	val   int
	left  TreeNodeRanger
	right TreeNodeRanger
}

// RangeSum counts sum of tree for each element that >= L && <= R
func (t *treeNode) RangeSum(root TreeNodeRanger, L int, R int) (sumBST int) {
	val := root.getVal()

	if root == nil {
		return sumBST
	}
	if val >= L && val <= R {
		sumBST += val
	}
	if val > L && root.showLeft() != nil {
		sumBST += t.RangeSum(root.showLeft(), L, R)
	}
	if val < R && root.showRight() != nil {
		sumBST += t.RangeSum(root.showRight(), L, R)
	}
	return sumBST
}

func (t *treeNode) getVal() int {
	return t.val
}

func (t *treeNode) showLeft() TreeNodeRanger {
	return t.left
}

func (t *treeNode) showRight() TreeNodeRanger {
	return t.right
}

//  NewTreeNodeRanger returns new tree node with input parameters
func NewTreeNodeRanger(tleft TreeNodeRanger, tright TreeNodeRanger, val int) TreeNodeRanger {
	return &treeNode{
		val:   val,
		left:  tleft,
		right: tright,
	}
}
