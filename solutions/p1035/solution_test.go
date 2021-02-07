package solution

import "testing"

func Test_maxUncrossedLines(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		A    []int
		B    []int
		want int
	}{
		{[]int{}, []int{1, 2, 4}, 0},
		{[]int{1, 2, 4}, []int{}, 0},
		{[]int{1, 4, 2}, []int{1, 2, 4}, 2},
		{[]int{2, 5, 1, 2, 5}, []int{10, 5, 2, 1, 5, 2}, 3},
		{[]int{1, 3, 7, 1, 7, 5}, []int{1, 9, 2, 5, 1}, 2},
	}
	for _, tt := range tests {
		if got := maxUncrossedLines(tt.A, tt.B); got != tt.want {
			t.Errorf("maxUncrossedLines() = %v, want %v", got, tt.want)
		}
	}
}
