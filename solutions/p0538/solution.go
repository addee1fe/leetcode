package solution

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Morris traversal
func convertBST(root *TreeNode) *TreeNode {
	var sum int
	node := root
	for node != nil {
		if node.Right == nil {
			sum += node.Val
			node.Val = sum
			node = node.Left
		} else {
			succ := getSuccessor(node)
			if succ.Left == nil {
				succ.Left = node
				node = node.Right
			} else {
				succ.Left = nil
				sum += node.Val
				node.Val = sum
				node = node.Left
			}
		}
	}
	return root
}

func getSuccessor(node *TreeNode) *TreeNode {
	succ := node.Right
	for succ.Left != nil && succ.Left != node {
		succ = succ.Left
	}
	return succ
}

// Recursion solution

// func bstToGst(root *TreeNode) *TreeNode {
//     var sum int

//     helper(root, &sum)

//     return root
// }

// func helper(root *TreeNode, sum *int){
//     if root == nil{
//         return
//     }
//     helper(root.Right, sum)
//     *sum += root.Val
//     root.Val = *sum
//     helper(root.Left, sum)
// }
