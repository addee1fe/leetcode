package solution

import (
	"reflect"
	"testing"
)

func TestShortestToChar(t *testing.T) {
	tests := []struct {
		s    string
		c    byte
		want []int
	}{
		{
			"loveleetcode",
			'c',
			[]int{3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0},
		},
		{
			"aaab",
			'b',
			[]int{3, 2, 1, 0},
		},
		{
			"aaba",
			'b',
			[]int{2, 1, 0, 1},
		},
	}
	for _, tt := range tests {
		if got := shortestToChar(tt.s, tt.c); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("shortestToChar() = %v, want %v", got, tt.want)
		}
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		a, b int
		want int
	}{
		{1, 2, 1},
		{2, 3, 2},
		{3, 4, 3},
		{4, 5, 4},
		{7, 6, 6},
		{10, 8, 8},
	}
	for _, tt := range tests {
		if got := min(tt.a, tt.b); got != tt.want {
			t.Errorf("min() = %v, want %v", got, tt.want)
		}
	}
}
