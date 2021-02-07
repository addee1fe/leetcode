package solution

import (
	"reflect"
	"testing"
)

func TestReconstructQueue(t *testing.T) {

	tests := []struct {
		people [][]int
		want   [][]int
	}{
		{[][]int{{1, 1}}, [][]int{{1, 1}}},
		{[][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}, [][]int{{5, 0}, {7, 0}, {5, 2}, {6, 1}, {4, 4}, {7, 1}}},
	}
	for _, tt := range tests {
		if got := reconstructQueue(tt.people); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("reconstructQueue() = %v, want %v", got, tt.want)
		}
	}
}
