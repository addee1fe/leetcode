package solution

import "testing"

func TestAngleClock(t *testing.T) {
	tests := []struct {
		hour    int
		minutes int
		want    float64
	}{
		{12, 30, 165},
		{3, 30, 75},
		{3, 15, 7.5},
		{4, 50, 155},
		{12, 0, 0},
	}
	for _, tt := range tests {
		if got := angleClock(tt.hour, tt.minutes); got != tt.want {
			t.Errorf("angleClock() = %v, want %v", got, tt.want)
		}
	}
}
