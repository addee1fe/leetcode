package solution

import "testing"

func TestIsSameTree(t *testing.T) {
	tests := []struct {
		p    *TreeNode
		q    *TreeNode
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := isSameTree(tt.p, tt.q); got != tt.want {
			t.Errorf("isSameTree() = %v, want %v", got, tt.want)
		}
	}
}
