package solution

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

// dummy solution
func deleteNode(node *ListNode) {
	if node == nil {
		return
	}
	*node = *node.Next
}
