package solution

import (
	"reflect"
	"testing"
)

func TestBstToGst(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want *TreeNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := getSuccessor(tt.root); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("bstToGst() = %v, want %v", got, tt.want)
		}
	}
}
