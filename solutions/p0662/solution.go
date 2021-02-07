package solution

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func widthOfBinaryTree(root *TreeNode) int {
	var l []int
	return dfs(root, 1, 0, &l)
}

func dfs(node *TreeNode, id int, depth int, lefts *[]int) int {
	if node == nil {
		return 0
	}
	if depth == len(*lefts) {
		*lefts = append(*lefts, id)
	}
	return max(id-(*lefts)[depth]+1, max(dfs(node.Left, id*2, depth+1, lefts), dfs(node.Right, id*2+1, depth+1, lefts)))
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
