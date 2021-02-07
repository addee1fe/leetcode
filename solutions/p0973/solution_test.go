package solution

import (
	"reflect"
	"testing"
)

func TestKClosest(t *testing.T) {
	tests := []struct {
		points [][]int
		K      int
		want   [][]int
	}{
		{[][]int{{1, 3}, {-2, 2}}, 2, [][]int{{1, 3}, {-2, 2}}},
		{[][]int{{1, 3}, {-2, 2}}, 1, [][]int{{-2, 2}}},
		{[][]int{{3, 3}, {5, -1}, {-2, 4}}, 2, [][]int{{3, 3}, {-2, 4}}},
	}
	for _, tt := range tests {
		if got := kClosest(tt.points, tt.K); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("kClosest() = %v, want %v", got, tt.want)
		}
	}
}
