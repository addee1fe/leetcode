package solution

import (
	"testing"
)

func Test_minDistance(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		word1 string
		word2 string
		want  int
	}{

		{"", "", 0},
		{"", "abc", 3},
		{"abc", "", 3},
		{"horse", "ros", 3},
		{"intention", "execution", 5},
	}
	for _, tt := range tests {
		if got := minDistance(tt.word1, tt.word2); got != tt.want {
			t.Errorf("minDistance() = %v, want %v", got, tt.want)
		}
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		a    []int
		want int
	}{
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{1, 2, 3, 4, 5}, 1},
		{[]int{1, 5, 2, -1, 5}, -1},
	}
	for _, tt := range tests {
		if got := min(tt.a...); got != tt.want {
			t.Errorf("min() = %v, want %v", got, tt.want)
		}
	}
}
