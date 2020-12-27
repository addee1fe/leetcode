package solution

import (
	"reflect"
	"testing"
)

// Hardcoded tests
// TODO: maybe should implement linkedlist generator and use it intead
// but for simply testing purpose this solution is ok
func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		l1   *ListNode
		l2   *ListNode
		want *ListNode
	}{
		{
			l1:   &ListNode{0, nil},
			l2:   &ListNode{0, nil},
			want: &ListNode{0, nil},
		},
		{
			l1:   &ListNode{2, &ListNode{4, &ListNode{3, nil}}},
			l2:   &ListNode{5, &ListNode{6, &ListNode{4, nil}}},
			want: &ListNode{7, &ListNode{0, &ListNode{8, nil}}},
		},
		{
			l1:   &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}}},
			l2:   &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}},
			want: &ListNode{8, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{1, nil}}}}}}}},
		},
	}
	for _, tt := range tests {
		if got := addTwoNumbers(tt.l1, tt.l2); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got: %v, want: %v", got, tt.want)
		}
	}
}
