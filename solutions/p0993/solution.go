package solution

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Dummy solution, but beat 100% of submissions with 0ms runtime
func isCousins(root *TreeNode, x int, y int) bool {
	if root == nil || root.Val == x || root.Val == y {
		return false
	}
	m := make(map[int][2]int)
	helper(root, 0, -1, m)
	return m[x][0] == m[y][0] && m[x][1] != m[y][1]
}

func helper(root *TreeNode, depth, parent int, m map[int][2]int) {
	if root == nil {
		return
	}
	m[root.Val] = [2]int{depth, parent}
	helper(root.Left, depth+1, root.Val, m)
	helper(root.Right, depth+1, root.Val, m)
}
