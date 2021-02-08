package solution

import "testing"

func TestIntegerReplacement(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{
			8,
			3,
		},
		{
			7,
			4,
		},
		{
			4,
			2,
		},
		{
			3,
			2,
		},
	}
	for _, tt := range tests {
		if got := integerReplacement(tt.n); got != tt.want {
			t.Errorf("integerReplacement() = %v, want %v", got, tt.want)
		}
	}
}
