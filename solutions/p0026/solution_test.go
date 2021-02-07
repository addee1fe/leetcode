package solution

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 1, 2}, 2},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
	}
	for _, tt := range tests {
		if got := removeDuplicates(tt.nums); got != tt.want {
			t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
		}

	}
}
