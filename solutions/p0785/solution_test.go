package solution

import "testing"

func TestIsBipartite(t *testing.T) {
	tests := []struct {
		graph [][]int
		want  bool
	}{
		{
			[][]int{
				{1, 3},
				{0, 2},
				{1, 3},
				{0, 2},
			},
			true,
		},
		{
			[][]int{
				{1, 2, 3},
				{0, 2},
				{0, 1, 3},
				{0, 2},
			},
			false,
		},
	}
	for _, tt := range tests {
		if got := isBipartite(tt.graph); got != tt.want {
			t.Errorf("isBipartite() = %v, want %v", got, tt.want)
		}
	}
}
