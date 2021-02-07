package solution

import "testing"

func TestBitwiseComplement(t *testing.T) {
	tests := []struct {
		N    int
		want int
	}{
		{5, 2},
		{7, 0},
		{10, 5},
	}
	for _, tt := range tests {
		if got := bitwiseComplement(tt.N); got != tt.want {
			t.Errorf("bitwiseComplement() = %v, want %v", got, tt.want)
		}
	}
}
