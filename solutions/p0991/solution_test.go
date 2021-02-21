package solution

import "testing"

func TestBrokenCalc(t *testing.T) {
	tests := []struct {
		x    int
		y    int
		want int
	}{
		{
			2,
			3,
			2,
		},
		{
			5,
			8,
			2,
		},
		{
			3,
			10,
			3,
		},
		{
			1024,
			1,
			1023,
		},
	}
	for _, tt := range tests {
		if got := brokenCalc(tt.x, tt.y); got != tt.want {
			t.Errorf("brokenCalc() = %v, want %v", got, tt.want)
		}
	}
}
