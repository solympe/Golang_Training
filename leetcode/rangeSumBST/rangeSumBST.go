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
		if root.Val > L {
			sumBST += RangeSumBST(root.Left, L, R)
		}
		if root.Val < R {
			sumBST += RangeSumBST(root.Right, L, R)
		}
	}
	return sumBST
}
