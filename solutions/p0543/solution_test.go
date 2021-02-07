package solution

import (
	"testing"
)

func TestDiameterOfBinaryTree(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := diameterOfBinaryTree(tt.root); got != tt.want {
			t.Errorf("diameterOfBinaryTree() = %v, want %v", got, tt.want)
		}
	}
}
