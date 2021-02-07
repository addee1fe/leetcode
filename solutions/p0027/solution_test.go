package solution

import "testing"

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		nums []int
		val  int
		want int
	}{
		{[]int{3, 2, 2, 3}, 3, 2},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5},
	}
	for _, tt := range tests {
		if got := removeElement(tt.nums, tt.val); got != tt.want {
			t.Errorf("removeElement() = %v, want %v", got, tt.want)
		}
	}
}
