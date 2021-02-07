package solution

import "testing"

func Test_numTrees(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{3, 5},
	}
	for _, tt := range tests {
		if got := numTrees(tt.n); got != tt.want {
			t.Errorf("numTrees() = %v, want %v", got, tt.want)
		}
	}
}
