package solution

import (
	"reflect"
	"testing"
)

func TestAverageOfLevels(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want []float64
	}{
		{
			&TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}},
			[]float64{3, 14.5, 11},
		},
	}
	for _, tt := range tests {
		if got := averageOfLevels(tt.root); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("averageOfLevels() = %v, want %v", got, tt.want)
		}
	}
}
