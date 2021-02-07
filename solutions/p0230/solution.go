package solution

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	res := make([]int, 0)
	helper(root, &res)
	return res[k-1]
}

func helper(r *TreeNode, res *[]int) {
	if r == nil {
		return
	}
	helper(r.Left, res)
	*res = append(*res, r.Val)
	helper(r.Right, res)
}

// func kthSmallest(root *TreeNode, k int) int {
//     node, _ := kthSmallestHelper(root, k, 0)
//     return node.Val
// }

// func kthSmallestHelper(node *TreeNode, k int, seen int) (*TreeNode, int){
//     if node == nil {
//         return nil, seen
//     }

//     newNode, seen := kthSmallestHelper(node.Left, k, seen)
//     if newNode != nil {
//         return newNode, seen
//     }

//     seen += 1
//     if seen == k {
//         return node, seen
//     }

//     newNode, seen = kthSmallestHelper(node.Right, k, seen)
//     if newNode != nil {
//         return newNode, seen
//     }

//     return nil, seen
// }
