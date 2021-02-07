package solution

import "testing"

func TestIsPerfectSquare(t *testing.T) {
	tests := []struct {
		num  int
		want bool
	}{
		{16, true},
		{14, false},
	}
	for _, tt := range tests {
		if got := isPerfectSquare(tt.num); got != tt.want {
			t.Errorf("isPerfectSquare() = %v, want %v", got, tt.want)
		}
	}
}
