package solution

import (
	"reflect"
	"testing"
)

func TestPrisonAfterNDays(t *testing.T) {
	tests := []struct {
		cells []int
		N     int
		want  []int
	}{
		{[]int{0, 1, 0, 1, 1, 0, 0, 1}, 7, []int{0, 0, 1, 1, 0, 0, 0, 0}},
		{[]int{1, 0, 0, 1, 0, 0, 1, 0}, 1000000000, []int{0, 0, 1, 1, 1, 1, 1, 0}},
	}
	for _, tt := range tests {
		if got := prisonAfterNDays(tt.cells, tt.N); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("prisonAfterNDays() = %v, want %v", got, tt.want)
		}
	}
}
