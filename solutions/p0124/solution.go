package solution

import "math"

// TreeNode is definition for a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	res := math.MinInt32
	helper(root, &res)
	return res
}

func helper(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	left := max(helper(root.Left, res), 0)
	right := max(helper(root.Right, res), 0)
	*res = max(*res, root.Val+left+right)
	return root.Val + max(left, right)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
