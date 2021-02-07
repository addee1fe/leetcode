package solution

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	res := make([]int, 0)
	helper(root, &res, 0)
	return res
}

func helper(root *TreeNode, res *[]int, lvl int) {
	if root == nil {
		return
	}
	if len(*res) <= lvl {
		*res = append(*res, root.Val)
	}
	// change order to get the left side of tree
	helper(root.Right, res, lvl+1)
	helper(root.Left, res, lvl+1)
}
