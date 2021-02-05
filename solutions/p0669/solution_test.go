package solution

import (
	"reflect"
	"testing"
)

func TestTrimBST(t *testing.T) {
	tests := []struct {
		root *TreeNode
		low  int
		high int
		want *TreeNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := trimBST(tt.root, tt.low, tt.high); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("trimBST() = %v, want %v", got, tt.want)
		}
	}
}
