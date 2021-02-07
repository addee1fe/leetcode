package solution

import "testing"

func TestPossibleBipartition(t *testing.T) {
	tests := []struct {
		N        int
		dislikes [][]int
		want     bool
	}{
		{0, [][]int{{1, 2}, {3, 4}}, false},
		{2, [][]int{}, true},
		{3, [][]int{{1, 2}, {1, 3}, {2, 3}}, false},
		{4, [][]int{{1, 2}, {1, 3}, {2, 4}}, true},
		{5, [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {1, 5}}, false},
	}
	for _, tt := range tests {
		if got := possibleBipartition(tt.N, tt.dislikes); got != tt.want {
			t.Errorf("possibleBipartition() = %v, want %v", got, tt.want)
		}

	}
}
