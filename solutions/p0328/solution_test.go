package solution

import (
	"reflect"
	"testing"
)

func TestOddEvenList(t *testing.T) {

	tests := []struct {
		head *ListNode
		want *ListNode
	}{
		// Hardcoded tests
		{
			&ListNode{},
			&ListNode{},
		},
		{
			&ListNode{1, nil},
			&ListNode{1, nil},
		},
		{
			&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
			&ListNode{1, &ListNode{3, &ListNode{5, &ListNode{2, &ListNode{4, nil}}}}}},
		{
			&ListNode{2, &ListNode{1, &ListNode{3, &ListNode{5, &ListNode{6, &ListNode{4, &ListNode{7, nil}}}}}}},
			&ListNode{2, &ListNode{3, &ListNode{6, &ListNode{7, &ListNode{1, &ListNode{5, &ListNode{4, nil}}}}}}},
		},
	}
	for _, tt := range tests {
		if got := oddEvenList(tt.head); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("oddEvenList() = %v, want %v", got, tt.want)
		}
	}
}
