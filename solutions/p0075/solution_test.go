package solution

import (
	"reflect"
	"testing"
)

func TestSortColors(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
	}
	for _, tt := range tests {
		sortColors(tt.nums)
		if !reflect.DeepEqual(tt.nums, tt.want) {
			t.Errorf("sortColors(): %v, want: %v", tt.nums, tt.want)
		}
	}
}

func TestSortColorsDummy(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
	}
	for _, tt := range tests {
		sortColorsDummy(tt.nums)
		if !reflect.DeepEqual(tt.nums, tt.want) {
			t.Errorf("sortColors(): %v, want: %v", tt.nums, tt.want)
		}
	}
}

func BenchmarkSortColors(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sortColors([]int{2, 0, 2, 1, 1, 0})
	}
}
func BenchmarkSortColorsDummy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sortColorsDummy([]int{2, 0, 2, 1, 1, 0})
	}
}
