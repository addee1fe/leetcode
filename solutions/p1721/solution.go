package solution

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapNodes(head *ListNode, k int) *ListNode {
	var a, b *ListNode
	for p := head; p != nil; p = p.Next {
		k--
		if b != nil {
			b = b.Next
		}
		if k == 0 {
			a = p
			b = head
		}
	}
	a.Val, b.Val = b.Val, a.Val
	return head
}
