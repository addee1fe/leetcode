package solution

import "testing"

func TestFindNumbers(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			[]int{12, 345, 2, 6, 7896},
			2,
		},
		{
			[]int{555, 901, 482, 1771},
			1,
		},
	}
	for _, tt := range tests {

		if got := findNumbers(tt.nums); got != tt.want {
			t.Errorf("findNumbers() = %v, want %v", got, tt.want)
		}
	}
}

func TestCountDigits(t *testing.T) {
	tests := []struct {
		a    int
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
			12,
			2,
		},
		{
			123,
			3,
		},
		{
			1234,
			4,
		},
		{
			12345,
			5,
		},
	}
	for _, tt := range tests {
		if got := countDigits(tt.a); got != tt.want {
			t.Errorf("countDigits() = %v, want %v", got, tt.want)
		}
	}
}
