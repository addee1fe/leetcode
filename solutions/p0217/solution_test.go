package solution

import "testing"

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		{
			[]int{1, 2, 3, 1},
			true,
		},
		{
			[]int{1, 2, 3, 4},
			false,
		},
		{
			[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			true,
		},
	}
	for _, tt := range tests {
		if got := containsDuplicate(tt.nums); got != tt.want {
			t.Errorf("containsDuplicate() = %v, want %v", got, tt.want)
		}
	}
}
