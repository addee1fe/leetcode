package solution

import (
	"reflect"
	"testing"
)

func TestSwapNodes(t *testing.T) {
	tests := []struct {
		head *ListNode
		k    int
		want *ListNode
	}{
		{
			&ListNode{1, nil},
			1,
			&ListNode{1, nil},
		},
		{
			&ListNode{1, &ListNode{2, nil}},
			1,
			&ListNode{2, &ListNode{1, nil}},
		},
		{
			&ListNode{1, &ListNode{2, &ListNode{3, nil}}},
			2,
			&ListNode{1, &ListNode{2, &ListNode{3, nil}}},
		},
		{
			&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
			2,
			&ListNode{1, &ListNode{4, &ListNode{3, &ListNode{2, &ListNode{5, nil}}}}},
		},
		{
			&ListNode{7, &ListNode{9, &ListNode{6, &ListNode{6, &ListNode{7, &ListNode{8, &ListNode{3, &ListNode{0, &ListNode{9, &ListNode{5, nil}}}}}}}}}},
			5,
			&ListNode{7, &ListNode{9, &ListNode{6, &ListNode{6, &ListNode{8, &ListNode{7, &ListNode{3, &ListNode{0, &ListNode{9, &ListNode{5, nil}}}}}}}}}},
		},
	}
	for _, tt := range tests {
		if got := swapNodes(tt.head, tt.k); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("swapNodes() = %v, want %v", got, tt.want)
		}
	}
}
