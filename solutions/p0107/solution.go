package solution

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	res := make([][]int, 0)
	helper(root, &res, 0)
	return res
}
func helper(node *TreeNode, values *[][]int, level int) {
	if node == nil {
		return
	}
	if len(*values) <= level {
		*values = append([][]int{{}}, *values...)
	}
	cur := &(*values)[len(*values)-1-level]
	*cur = append(*cur, node.Val)
	helper(node.Left, values, level+1)
	helper(node.Right, values, level+1)
}
