package solution

import "container/list"

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func addOneRow(root *TreeNode, v int, d int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: v}
	}
	if d == 1 {
		return &TreeNode{Val: v, Left: root}
	}
	queue := list.New()
	queue.PushBack(root)
	depth := 1
	for queue.Len() > 0 {
		n := queue.Len()
		for n > 0 {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if depth+1 == d {
				node.Left = &TreeNode{Val: v, Left: node.Left}
				node.Right = &TreeNode{Val: v, Right: node.Right}
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			n--
		}
		depth++
		if depth == d {
			break
		}
	}
	return root
}
