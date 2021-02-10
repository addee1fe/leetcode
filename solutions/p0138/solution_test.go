package solution

import (
	"reflect"
	"testing"
)

func TestCopyRandomList(t *testing.T) {
	tests := []struct {
		head *Node
		want *Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := copyRandomList(tt.head); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("copyRandomList() = %v, want %v", got, tt.want)
		}
	}
}
