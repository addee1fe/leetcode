package solution

import (
	"reflect"
	"testing"
)

func Test_floodFill(t *testing.T) {
	tests := []struct {
		image    [][]int
		sr       int
		sc       int
		newColor int
		want     [][]int
	}{
		{
			[][]int{[]int{1, 1, 1}, []int{1, 1, 0}, []int{1, 0, 1}},
			1,
			1,
			2,
			[][]int{[]int{2, 2, 2}, []int{2, 2, 0}, []int{2, 0, 1}},
		},
	}
	for _, tt := range tests {
		if got := floodFill(tt.image, tt.sr, tt.sc, tt.newColor); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("floodFill() = %v, want %v", got, tt.want)
		}
	}
}
