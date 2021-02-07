package solution

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	var sum int
	if root != nil {
		helper(root, 0, &sum)
	}
	return sum
}

func helper(root *TreeNode, value int, sum *int) {
	if root.Left == nil && root.Right == nil {
		*sum += value*10 + root.Val
		return
	}
	if root.Left != nil {
		helper(root.Left, value*10+root.Val, sum)
	}
	if root.Right != nil {
		helper(root.Right, value*10+root.Val, sum)
	}
}

// func getSum(root *TreeNode,finalSum int) int {
//     if root == nil {
//         return 0
//     }
//     finalSum = finalSum*10 + root.Val
//     if root.Left == nil && root.Right == nil {
//         return finalSum
//     }
//     return getSum(root.Left, finalSum)+getSum(root.Right,finalSum)
// }
