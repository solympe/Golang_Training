package range_sum_bst

// TreeNode ...
type TreeNode interface {
	RangeSum(root TreeNode, L int, R int) (sumBST int)
	getVal() int
	showLeft() TreeNode
	showRight() TreeNode
}

// treeNode is a definition for a binary treenode node
type treeNode struct {
	val   int
	left  TreeNode
	right TreeNode
}

// RangeSum counts sum of treenode for each element that >= L && <= R
func (t *treeNode) RangeSum(root TreeNode, L int, R int) (sumBST int) {
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

func (t *treeNode) showLeft() TreeNode {
	return t.left
}

func (t *treeNode) showRight() TreeNode {
	return t.right
}

//  NewTreeNode returns new treenode node with input parameters
func NewTreeNode(tleft TreeNode, tright TreeNode, val int) TreeNode {
	return &treeNode{
		val:   val,
		left:  tleft,
		right: tright,
	}
}
