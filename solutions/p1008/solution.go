package solution

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstFromPreorder(preorder []int) *TreeNode {
	root := &TreeNode{preorder[0], nil, nil}
	for _, v := range preorder[1:] {
		root = insert(root, v)
	}
	return root
}

func insert(node *TreeNode, i int) *TreeNode {
	switch {
	case node == nil:
		return &TreeNode{i, nil, nil}
	case node.Val > i:
		node.Left = insert(node.Left, i)
	case node.Val < i:
		node.Right = insert(node.Right, i)
	}
	return node
}
