package solution

import (
	"reflect"
	"testing"
)

func Test_bstFromPreorder(t *testing.T) {
	tests := []struct {
		preorder []int
		want     *TreeNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := bstFromPreorder(tt.preorder); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("bstFromPreorder() = %v, want %v", got, tt.want)
		}
	}
}
