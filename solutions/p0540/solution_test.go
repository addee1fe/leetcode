package solution

import "testing"

func TestSingleNonDuplicate(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1}, 1},
		{[]int{1, 1, 2, 3, 3, 4, 4, 8, 8}, 2},
		{[]int{3, 3, 7, 7, 10, 11, 11}, 10},
	}
	for _, tt := range tests {
		if got := singleNonDuplicate(tt.nums); got != tt.want {
			t.Errorf("singleNonDuplicate() = %v, want %v", got, tt.want)
		}
	}
}

func TestSingleNonDuplicateXor(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1}, 1},
		{[]int{1, 1, 2, 3, 3, 4, 4, 8, 8}, 2},
		{[]int{3, 3, 7, 7, 10, 11, 11}, 10},
	}
	for _, tt := range tests {
		if got := singleNonDuplicateXor(tt.nums); got != tt.want {
			t.Errorf("singleNonDuplicate() = %v, want %v", got, tt.want)
		}
	}
}

func BenchmarkSingleNonDuplicate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		singleNonDuplicate([]int{1, 1, 2, 3, 3, 4, 4, 8, 8})
	}
}

func BenchmarkSingleNonDuplicateXor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		singleNonDuplicateXor([]int{1, 1, 2, 3, 3, 4, 4, 8, 8})
	}
}
