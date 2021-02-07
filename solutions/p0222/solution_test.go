package solution

import "testing"

func Test_countNodes(t *testing.T) {

	tests := []struct {
		root *TreeNode
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := countNodes(tt.root); got != tt.want {
			t.Errorf("countNodes() = %v, want %v", got, tt.want)
		}
	}
}
