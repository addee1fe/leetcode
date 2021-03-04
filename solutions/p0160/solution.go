package solution

// ListNode is definition for singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	hA, hB := headA, headB

	for hA != hB {
		if hA == nil {
			hA = headB
		} else {
			hA = hA.Next
		}
		if hB == nil {
			hB = headA
		} else {
			hB = hB.Next
		}
	}
	return hA
}
