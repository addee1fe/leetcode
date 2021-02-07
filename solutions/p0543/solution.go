package solution

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	max := 0
	dfs(root, &max)
	return max
}

func dfs(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}
	l := dfs(root.Left, max)
	r := dfs(root.Right, max)
	if *max < l+r {
		*max = l + r
	}
	if l > r {
		return l + 1
	}
	return r + 1
}
