package solution

import "testing"

func Test_singleNumber(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{2, 2, 3, 2}, 3},
		{[]int{0, 1, 0, 1, 0, 1, 99}, 99},
	}
	for _, tt := range tests {
		if got := singleNumber(tt.nums); got != tt.want {
			t.Errorf("singleNumber() = %v, want %v", got, tt.want)
		}
	}
}
