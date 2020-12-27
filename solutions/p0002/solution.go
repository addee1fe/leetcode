package solution

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	p, q, curr := l1, l2, dummy
	carry := 0
	for p != nil || q != nil {
		x, y := 0, 0
		if p != nil {
			x = p.Val
		}
		if q != nil {
			y = q.Val
		}
		sum := carry + x + y
		carry = sum / 10
		curr.Val = sum % 10
		if p != nil {
			p = p.Next
		}
		if q != nil {
			q = q.Next
		}
		if p != nil || q != nil {
			curr.Next = &ListNode{}
			curr = curr.Next
		}
	}
	if carry > 0 {
		curr.Next = &ListNode{carry, nil}
	}
	return dummy
}
