package solution

import "testing"

// Some hardcoded tests
func TestTree2str(t *testing.T) {
	tests := []struct {
		t    *TreeNode
		want string
	}{
		{
			&TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, nil}, &TreeNode{3, nil, nil}},
			"1(2(4))(3)",
		},
		{
			&TreeNode{1, &TreeNode{2, nil, &TreeNode{4, nil, nil}}, &TreeNode{3, nil, nil}},
			"1(2()(4))(3)",
		},
	}

	for _, tt := range tests {
		if got := tree2str(tt.t); got != tt.want {
			t.Errorf("tree2str() = %v, want %v", got, tt.want)
		}
	}
}
