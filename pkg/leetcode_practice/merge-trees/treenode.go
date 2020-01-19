package treenode

// TreeNode ...
type TreeNode interface {
	MergeTrees(t2 TreeNode) TreeNode
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
func (t *treeNode) MergeTrees(t2 TreeNode) TreeNode {
	t1 := t
	if t1 == nil && t2 == nil {
		return nil
	}
	if t1 == nil {
		return t2
	}

	if t2 == nil {
		if t1.val == 0 {
			return nil
		}
		return t1
	}
	if t1.left == nil {
		t1.left = NewTreeNode(0, nil, nil)
	} else if t1.right == nil {
		t1.right = NewTreeNode(0, nil, nil)
	}
	if t1.val == 0 {
		return t2
	}

	t1.val += t2.GetVal()
	t1.setLeft(t1.left.MergeTrees(t2.GetLeft()))
	t1.setRight(t1.right.MergeTrees(t2.GetRight()))
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
