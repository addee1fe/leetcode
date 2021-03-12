package solution

import "testing"

func TestHasAllCodes(t *testing.T) {
	tests := []struct {
		s    string
		k    int
		want bool
	}{
		{
			"00110110",
			2,
			true,
		},
		{
			"00110",
			2,
			true,
		},
		{
			"0110",
			1,
			true,
		},
		{
			"0110",
			2,
			false,
		},
		{
			"0000000001011100",
			4,
			false,
		},
	}
	for _, tt := range tests {
		if got := hasAllCodes(tt.s, tt.k); got != tt.want {
			t.Errorf("hasAllCodes() = %v, want %v", got, tt.want)
		}
	}
}
