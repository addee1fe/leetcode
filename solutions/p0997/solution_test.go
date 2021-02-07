package solution

import "testing"

func Test_findJudge(t *testing.T) {
	tests := []struct {
		N     int
		trust [][]int
		want  int
	}{
		{1, [][]int{}, 1},
		{2, [][]int{[]int{1, 2}}, 2},
		{3, [][]int{[]int{1, 3}, []int{2, 3}}, 3},
		{3, [][]int{[]int{1, 3}, []int{2, 3}, []int{3, 1}}, -1},
		{3, [][]int{[]int{1, 2}, []int{2, 3}}, -1},
		{4, [][]int{[]int{1, 3}, []int{1, 4}, []int{2, 3}, []int{2, 4}, []int{4, 3}}, 3},
	}
	for _, tt := range tests {
		if got := findJudge(tt.N, tt.trust); got != tt.want {
			t.Errorf("findJudge() = %v, want %v", got, tt.want)
		}
	}
}
