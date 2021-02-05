package solution

import "testing"

func TestIsValidSequence(t *testing.T) {
	type args struct {
		root *TreeNode
		arr  []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidSequence(tt.args.root, tt.args.arr); got != tt.want {
				t.Errorf("isValidSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHelper(t *testing.T) {
	type args struct {
		node *TreeNode
		arr  []int
		i    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := helper(tt.args.node, tt.args.arr, tt.args.i); got != tt.want {
				t.Errorf("helper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsLeaf(t *testing.T) {
	type args struct {
		node *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isLeaf(tt.args.node); got != tt.want {
				t.Errorf("isLeaf() = %v, want %v", got, tt.want)
			}
		})
	}
}
