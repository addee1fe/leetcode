package main

import "testing"

func TestMissingNumber(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			[]int{3, 0, 1},
			2,
		},
		{
			[]int{0, 1},
			2,
		},
		{
			[]int{9, 6, 4, 2, 3, 5, 7, 0, 1},
			8,
		},
		{
			[]int{0},
			1,
		},
	}
	for _, tt := range tests {
		if got := missingNumber(tt.nums); got != tt.want {
			t.Errorf("missingNumber() = %v, want %v", got, tt.want)
		}
	}
}
