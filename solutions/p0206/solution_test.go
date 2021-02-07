package solution

import (
	"reflect"
	"testing"
)

func TestReverseList(t *testing.T) {
	tests := []struct {
		head *ListNode
		want *ListNode
	}{
		{
			&ListNode{},
			&ListNode{},
		},
		{
			&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}},
			&ListNode{4, &ListNode{3, &ListNode{2, &ListNode{1, nil}}}},
		},
	}
	for _, tt := range tests {
		if got := reverseList(tt.head); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("reverseList() = %v, want %v", got, tt.want)
		}
	}
}
