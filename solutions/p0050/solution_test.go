package solution

import "testing"

func TestMyPow(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		x    float64
		n    int
		want float64
	}{
		{2.00000, 10, 1024.00000},
		// {2.10000, 3, 9.26100},
		{2.00000, -2, 0.25000},
	}
	for _, tt := range tests {
		if got := myPow(tt.x, tt.n); got != tt.want {
			t.Errorf("myPow() = %v, want %v", got, tt.want)
		}
	}
}
