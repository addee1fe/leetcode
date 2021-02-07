package solution

import (
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		l1   *ListNode
		l2   *ListNode
		want *ListNode
	}{
		{
			&ListNode{1, &ListNode{3, &ListNode{5, &ListNode{7, nil}}}},
			&ListNode{2, &ListNode{4, &ListNode{6, &ListNode{8, nil}}}},
			&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, &ListNode{7, &ListNode{8, nil}}}}}}}},
		},
	}
	for _, tt := range tests {
		if got := mergeTwoLists(tt.l1, tt.l2); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
		}
	}
}
