package solution

import "testing"

func TestNumberOfArithmeticSlices(t *testing.T) {
	tests := []struct {
		A    []int
		want int
	}{
		{
			[]int{1, 2, 3, 4},
			3,
		},
	}
	for _, tt := range tests {
		if got := numberOfArithmeticSlices(tt.A); got != tt.want {
			t.Errorf("numberOfArithmeticSlices() = %v, want %v", got, tt.want)
		}
	}
}
