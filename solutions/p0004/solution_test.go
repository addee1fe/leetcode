package solution

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
		want  float64
	}{
		{
			[]int{1, 3},
			[]int{2},
			2.0,
		},
		{
			[]int{1, 2},
			[]int{3, 4},
			2.5,
		},
		{
			[]int{0, 0},
			[]int{0, 0},
			0.0,
		},
		{
			[]int{},
			[]int{1},
			1.0,
		},
		{
			[]int{2},
			[]int{},
			2.0,
		},
	}
	for _, tt := range tests {
		if got := findMedianSortedArrays(tt.nums1, tt.nums2); got != tt.want {
			t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
		}
	}
}
