package solution

import "testing"

func TestFindMaxLength(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{0, 1}, 2},
		{[]int{0, 1, 0}, 2},
	}
	for _, tt := range tests {
		if got := findMaxLength(tt.nums); got != tt.want {
			t.Errorf("findMaxLength() = %v, want %v", got, tt.want)
		}

	}
}
