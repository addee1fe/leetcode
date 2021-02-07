package solution

import (
	"reflect"
	"testing"
)

func TestFindOrder(t *testing.T) {
	tests := []struct {
		numCourses    int
		prerequisites [][]int
		want          []int
	}{
		{2, [][]int{{1, 0}}, []int{0, 1}},
		{4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}, []int{0, 1, 2, 3}},
	}
	for _, tt := range tests {
		if got := findOrder(tt.numCourses, tt.prerequisites); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("findOrder() = %v, want %v", got, tt.want)
		}
	}
}
