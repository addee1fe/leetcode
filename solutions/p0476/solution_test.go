package solution

import "testing"

func TestFindComplement(t *testing.T) {
	tests := []struct {
		N    int
		want int
	}{
		{5, 2},
		{7, 0},
		{10, 5},
	}
	for _, tt := range tests {
		if got := findComplement(tt.N); got != tt.want {
			t.Errorf("findComplement() = %v, want %v", got, tt.want)
		}
	}
}
