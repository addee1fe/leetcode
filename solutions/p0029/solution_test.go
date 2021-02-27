package solution

import "testing"

func TestDivide(t *testing.T) {
	tests := []struct {
		dividend int
		divisor  int
		want     int
	}{
		{
			10,
			3,
			3,
		},
		{
			7,
			-3,
			-2,
		},
		{
			0,
			1,
			0,
		},
		{
			1,
			1,
			1,
		},
		{
			-2147483648,
			-1,
			2147483647,
		},
	}
	for _, tt := range tests {
		if got := divide(tt.dividend, tt.divisor); got != tt.want {
			t.Errorf("divide() = %v, want %v", got, tt.want)
		}
	}
}
