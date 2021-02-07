package solution

import "testing"

func TestSearchInsert(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{}, 10, 0},
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 7, 4},
		{[]int{1, 3, 5, 6}, 0, 0},
	}
	for _, tt := range tests {
		if got := searchInsert(tt.nums, tt.target); got != tt.want {
			t.Errorf("searchInsert() = %v, want %v", got, tt.want)
		}
	}
}

func TestSearchInsertSort(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{}, 10, 0},
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 7, 4},
		{[]int{1, 3, 5, 6}, 0, 0},
	}
	for _, tt := range tests {
		if got := searchInsertSort(tt.nums, tt.target); got != tt.want {
			t.Errorf("searchInsertSort() = %v, want %v", got, tt.want)
		}
	}
}

func TestSearchInsertFor(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{}, 10, 0},
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 7, 4},
		{[]int{1, 3, 5, 6}, 0, 0},
	}
	for _, tt := range tests {
		if got := searchInsertFor(tt.nums, tt.target); got != tt.want {
			t.Errorf("searchInsertSort() = %v, want %v", got, tt.want)
		}
	}
}

func BenchmarkSearchInsertFor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		searchInsertFor([]int{1, 3, 5, 6}, 7)
	}
}
func BenchmarkSearchInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		searchInsert([]int{1, 3, 5, 6}, 7)
	}
}

func BenchmarkSearchInsertSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		searchInsertSort([]int{1, 3, 5, 6}, 7)
	}
}
