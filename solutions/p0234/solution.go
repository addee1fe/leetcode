package solution

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

// This is dumb idea but I like it more than using 2 pointers and reversing half of the linked list
func isPalindrome(head *ListNode) bool {
	values := make([]int, 0)
	for ; head != nil; head = head.Next {
		values = append(values, head.Val)
	}
	for left, right := 0, len(values)-1; left < right; left, right = left+1, right-1 {
		if values[left] != values[right] {
			return false
		}
	}
	return true
}
