package solution

import "testing"

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
		{[]int{3, 2, 3}, 3},
	}
	for _, tt := range tests {
		if got := majorityElement(tt.nums); got != tt.want {
			t.Errorf("majorityElement() = %v, want %v", got, tt.want)
		}
	}
}

func TestMajorityElementMap(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
		{[]int{3, 2, 3}, 3},
	}
	for _, tt := range tests {
		if got := majorityElementMap(tt.nums); got != tt.want {
			t.Errorf("majorityElementMap() = %v, want %v", got, tt.want)
		}
	}
}

func TestMajorityElementSort(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
		{[]int{3, 2, 3}, 3},
	}
	for _, tt := range tests {
		if got := majorityElementSort(tt.nums); got != tt.want {
			t.Errorf("majorityElement() = %v, want %v", got, tt.want)
		}
	}
}

func BenchmarkMajorityElement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		majorityElement([]int{2, 2, 1, 1, 1, 2, 2})
	}
}

func BenchmarkMajorityElementMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		majorityElementMap([]int{2, 2, 1, 1, 1, 2, 2})
	}
}

func BenchmarkMajorityElementSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		majorityElementSort([]int{2, 2, 1, 1, 1, 2, 2})
	}
}
