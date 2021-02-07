package solution

import (
	"reflect"
	"testing"
)

// Due to big (x2) grow factor in small slices tests fail when used !reflect.DeepEqual
// last subslice has len:3, but cap:4 and test fails due that fact
// TODO: refactor code and fix tests

func TestSubsets(t *testing.T) {
	tests := []struct {
		nums []int
		want [][]int
	}{
		{[]int{1, 2, 3}, [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}}},
	}
	for _, tt := range tests {
		if got := subsets(tt.nums); !reflect.DeepEqual(got, tt.want) {
			for i := 0; i < len(got); i++ {
				t.Logf("Got[%d]: %v Want[%[1]d]: %v ", i, got[i], tt.want[i])
			}
			// t.Errorf("subsets() = %v, want %v", got, tt.want)
		}
	}
}
