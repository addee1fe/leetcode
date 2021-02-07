package solution

import "testing"

func TestFindDuplicate(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 3, 4, 2, 2}, 2},
		{[]int{3, 1, 3, 4, 2}, 3},
	}
	for _, tt := range tests {
		if got := findDuplicate(tt.nums); got != tt.want {
			t.Errorf("findDuplicate() = %v, want %v", got, tt.want)
		}
	}
}
