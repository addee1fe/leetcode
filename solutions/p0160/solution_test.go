package solution

import (
	"reflect"
	"testing"
)

func TestGetIntersectionNode(t *testing.T) {
	tests := []struct {
		headA *ListNode
		headB *ListNode
		want  *ListNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := getIntersectionNode(tt.headA, tt.headB); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("getIntersectionNode() = %v, want %v", got, tt.want)
		}
	}
}
