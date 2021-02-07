package solution

// ListNode is definition for singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	head.Next = removeElements(head.Next, val)
	if head.Val == val {
		return head.Next
	}
	return head
}

// Iterative
// func removeElements(head *ListNode, val int) *ListNode {
//     for ; head!=nil && head.Val == val; {
//         head = head.Next
//     }
//     if head == nil {
//         return head
//     }
//     prev := head
//     cur := head.Next
//     for cur != nil {
//         if cur.Val == val {
//             prev.Next = cur.Next
//             cur = prev.Next
//         } else {
//             prev = prev.Next
//             cur = cur.Next
//         }
//     }
//     return head
// }
