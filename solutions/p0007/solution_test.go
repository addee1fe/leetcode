package solution

import "testing"

func Test_reverse(t *testing.T) {
	tests := []struct {
		x    int
		want int
	}{
		{
			1,
			1,
		},
		{
			0,
			0,
		},
		{
			120,
			21,
		},
		{
			123,
			321,
		},
		{
			1534236469,
			0,
		},
		{
			-2147483648,
			0,
		},
	}
	for _, tt := range tests {
		if got := reverse(tt.x); got != tt.want {
			t.Errorf("reverse(%v) = %v, want %v", tt.x, got, tt.want)
		}
	}
}
