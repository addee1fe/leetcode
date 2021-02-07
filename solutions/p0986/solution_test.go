package solution

import (
	"reflect"
	"testing"
)

func Test_intervalIntersection(t *testing.T) {
	tests := []struct {
		A    [][]int
		B    [][]int
		want [][]int
	}{
		{
			[][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}},
			[][]int{},
			[][]int{},
		},
		{
			[][]int{},
			[][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}},
			[][]int{},
		},
		{
			[][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}},
			[][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}},
			[][]int{{1, 2}, {5, 5}, {8, 10}, {15, 23}, {24, 24}, {25, 25}},
		},
	}
	for _, tt := range tests {
		if got := intervalIntersection(tt.A, tt.B); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("intervalIntersection() = %v, want %v", got, tt.want)
		}
	}
}
