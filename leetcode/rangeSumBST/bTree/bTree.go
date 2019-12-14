package bTree

type BTree interface {
	GetVal() int
	ShowLeft() *BTree
	ShowRight() *BTree
}

//Definition for a binary tree node
type treeNode struct {
	val   int
	left  *BTree
	right *BTree
}

// get value of node
func (t *treeNode) GetVal() int {
	return t.val
}

// get pointer on left node
func (t *treeNode) ShowLeft() *BTree {
	return t.left
}

// get pointer on right node
func (t *treeNode) ShowRight() *BTree {
	return t.right
}

//node constructor
func NewNode(tleft *BTree, tright *BTree, val int) BTree {
	return &treeNode{
		val:   val,
		left:  tleft,
		right: tright,
	}
}
