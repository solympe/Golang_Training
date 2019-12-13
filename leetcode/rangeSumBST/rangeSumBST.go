package rangeSumBST

//Definition for a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//count sum of tree
func RangeSumBST(root *TreeNode, L int, R int) int {
	var sumBST int

	if root != nil {
		if root.Val >= L && root.Val <= R {
			sumBST += root.Val
		}
		sumBST += RangeSumBST(root.Left, L, R)
		sumBST += RangeSumBST(root.Right, L, R)
	}
	return sumBST
}
