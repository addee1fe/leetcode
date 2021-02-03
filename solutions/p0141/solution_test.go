package solution

import "testing"

func TestHasCycle(t *testing.T) {
	tests := []struct {
		head *ListNode
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := hasCycle(tt.head); got != tt.want {
			t.Errorf("hasCycle() = %v, want %v", got, tt.want)
		}
	}
}
