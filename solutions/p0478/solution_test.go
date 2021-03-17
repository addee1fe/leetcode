package solution

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		radius   float64
		x_center float64
		y_center float64
		want     Solution
	}{
		{
			1,
			2,
			3,
			Solution{1, 2, 3},
		},
		{
			1,
			0,
			0,
			Solution{1, 0, 0},
		},
		{
			10,
			5,
			-7.5,
			Solution{10, 5, -7.5},
		},
	}
	for _, tt := range tests {
		if got := Constructor(tt.radius, tt.x_center, tt.y_center); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Constructor() = %v, want %v", got, tt.want)
		}
	}
}

func TestSolutionRandPoint(t *testing.T) {
	s := Constructor(10, 5, -7.5)

	for i := 0; i < 100; i++ {
		if !isPointInCircle(t, s, s.RandPoint()) {
			// todo: add more meaningful output
			t.Error("Generated random points is out of circle")
		}
	}
}

func isPointInCircle(t *testing.T, s Solution, points []float64) bool {
	t.Helper()
	// TODO: add formula
	return true
}
