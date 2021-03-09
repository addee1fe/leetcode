package solution

import (
	"reflect"
	"testing"
)

// hardcoded
func TestAddOneRow(t *testing.T) {
	tests := []struct {
		root *TreeNode
		v    int
		d    int
		want *TreeNode
	}{
		{
			&TreeNode{4, &TreeNode{2, &TreeNode{Val: 3}, &TreeNode{Val: 1}}, &TreeNode{6, &TreeNode{Val: 5}, nil}},
			1,
			2,
			&TreeNode{4, &TreeNode{1, &TreeNode{2, &TreeNode{Val: 3}, &TreeNode{Val: 1}}, nil}, &TreeNode{1, nil, &TreeNode{6, &TreeNode{Val: 5}, nil}}},
		},
		{
			&TreeNode{4, &TreeNode{2, &TreeNode{Val: 3}, &TreeNode{Val: 1}}, nil},
			1,
			3,
			&TreeNode{4, &TreeNode{2, &TreeNode{1, &TreeNode{Val: 3}, nil}, &TreeNode{1, nil, &TreeNode{Val: 1}}}, nil},
		},
	}
	for _, tt := range tests {
		if got := addOneRow(tt.root, tt.v, tt.d); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("addOneRow() = %v, want %v", got, tt.want)
		}
	}
}
