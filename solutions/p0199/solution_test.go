package solution

import (
	"reflect"
	"testing"
)

func TestRightSideView(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := rightSideView(tt.root); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("rightSideView() = %v, want %v", got, tt.want)
		}
	}
}
