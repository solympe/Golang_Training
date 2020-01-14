package b_tree

// BTree represents interface of binary tree struct
type BTree interface {
	GetVal() int
	ShowLeft() *BTree
	ShowRight() *BTree
}

// treeNode is a definition for a binary tree node
type treeNode struct {
	val   int
	left  *BTree
	right *BTree
}

// GetVal returns value of node
func (t *treeNode) GetVal() int {
	return t.val
}

// ShowLeft returns pointer to left node
func (t *treeNode) ShowLeft() *BTree {
	return t.left
}

// ShowRight returns pointer to right node
func (t *treeNode) ShowRight() *BTree {
	return t.right
}

// NewNode returns new btree node with input parameters
func NewNode(tleft *BTree, tright *BTree, val int) BTree {
	return &treeNode{
		val:   val,
		left:  tleft,
		right: tright,
	}
}
