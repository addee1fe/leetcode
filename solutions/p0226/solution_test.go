package solution

import (
	"reflect"
	"testing"
)

func TestInvertTree(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want *TreeNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := invertTree(tt.root); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("invertTree() = %v, want %v", got, tt.want)
		}
	}
}
