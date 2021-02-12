package solution

import "testing"

func TestNumberOfSteps(t *testing.T) {
	tests := []struct {
		num  int
		want int
	}{
		{
			14,
			6,
		},
		{
			8,
			4,
		},
		{
			123,
			12,
		},
	}
	for _, tt := range tests {
		if got := numberOfSteps(tt.num); got != tt.want {
			t.Errorf("numberOfSteps() = %v, want %v", got, tt.want)
		}
	}
}
