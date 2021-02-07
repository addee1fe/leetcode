package solution

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		head *ListNode
		want bool
	}{
		{
			&ListNode{1, &ListNode{2, nil}},
			false,
		},
		{
			&ListNode{1, &ListNode{2, &ListNode{2, &ListNode{1, nil}}}},
			true,
		},
	}
	for _, tt := range tests {
		if got := isPalindrome(tt.head); got != tt.want {
			t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
		}
	}
}
