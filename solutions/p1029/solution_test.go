package solution

import "testing"

func TestTwoCitySchedCost(t *testing.T) {

	tests := []struct {
		costs [][]int
		want  int
	}{
		{[][]int{{10, 20}, {30, 200}, {400, 50}, {30, 20}}, 110},
	}
	for _, tt := range tests {
		if got := twoCitySchedCost(tt.costs); got != tt.want {
			t.Errorf("twoCitySchedCost() = %v, want %v", got, tt.want)
		}
	}
}
