package treenode

// TreeNode ...
type TreeNode interface {
	MergeTrees(t1 TreeNode, t2 TreeNode) TreeNode
	GetLeft() TreeNode
	GetRight() TreeNode
	GetVal() int
	setVal(val int)
	setLeft(node TreeNode)
	setRight(node TreeNode)
}

type treeNode struct {
	val   int
	left  TreeNode
	right TreeNode
}

// MergeTrees ...
func (t *treeNode) MergeTrees(t1 TreeNode, t2 TreeNode) TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}
	if t1 == nil {
		return t2
	}

	if t2 == nil {
		return t1
	}
	t1.setVal(t2.GetVal() + t1.GetVal())
	t1.setLeft(t1.MergeTrees(t1.GetLeft(), t2.GetLeft()))
	t1.setRight(t1.MergeTrees(t1.GetRight(), t2.GetRight()))
	return t1
}

// GetLeft ...
func (t *treeNode) GetLeft() TreeNode {
	return t.left
}

// GetRight ...
func (t *treeNode) GetRight() TreeNode {
	return t.right
}

// GetVal ...
func (t *treeNode) GetVal() int {
	return t.val
}

func (t *treeNode) setLeft(node TreeNode) {
	t.left = node
}

func (t *treeNode) setRight(node TreeNode) {
	t.right = node
}

func (t *treeNode) setVal(val int) {
	t.val = val
}

// NewTreeNode ...
func NewTreeNode(val int, left TreeNode, right TreeNode) TreeNode {
	return &treeNode{
		val:   val,
		left:  left,
		right: right,
	}
}
