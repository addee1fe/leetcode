package solution

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// TODO: Refactor this
func isValidSequence(root *TreeNode, arr []int) bool {
	if root == nil || len(arr) == 0 {
		return false
	}
	return helper(root, arr, 0)
}

func helper(node *TreeNode, arr []int, i int) bool {
	if node.Val != arr[i] {
		return false
	}
	if i == len(arr)-1 {
		return isLeaf(node)
	}
	if node.Left != nil && helper(node.Left, arr, i+1) {
		return true
	}
	if node.Right != nil && helper(node.Right, arr, i+1) {
		return true
	}
	return false
}

func isLeaf(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}
