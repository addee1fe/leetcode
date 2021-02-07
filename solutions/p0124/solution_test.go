package solution

import (
	"testing"
)

func TestMaxPathSum(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := maxPathSum(tt.root); got != tt.want {
			t.Errorf("maxPathSum() = %v, want %v", got, tt.want)
		}
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		a, b int
		want int
	}{
		{1, 2, 2},
		{3, 2, 3},
	}
	for _, tt := range tests {
		if got := max(tt.a, tt.b); got != tt.want {
			t.Errorf("max() = %v, want %v", got, tt.want)
		}
	}
}
