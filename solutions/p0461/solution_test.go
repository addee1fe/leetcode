package solution

import "testing"

func TestHammingDistance(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		x    int
		y    int
		want int
	}{
		{
			1,
			4,
			2,
		},
	}
	for _, tt := range tests {
		if got := hammingDistance(tt.x, tt.y); got != tt.want {
			t.Errorf("hammingDistance() = %v, want %v", got, tt.want)
		}
	}
}

func TestPop(t *testing.T) {
	tests := []struct {
		x    int
		want int
	}{
		{
			0,
			0,
		},
		{
			1,
			1,
		},
		{
			2,
			1,
		},
		{
			3,
			2,
		},
		{
			4,
			1,
		},
		{
			5,
			2,
		},
		{
			6,
			2,
		},
		{
			7,
			3,
		},
	}
	for _, tt := range tests {
		if got := pop(tt.x); got != tt.want {
			t.Errorf("pop() = %v, want %v", got, tt.want)
		}
	}
}
