package solution

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
		{[]int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100}},
	}
	for _, tt := range tests {
		rotate(tt.nums, tt.k)
		if !reflect.DeepEqual(tt.nums, tt.want) {
			t.Errorf("rotate() = %v, want %v", tt.nums, tt.want)
		}
	}
}

func TestRotateSecond(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
		{[]int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100}},
	}
	for _, tt := range tests {
		rotateSecond(tt.nums, tt.k)
		if !reflect.DeepEqual(tt.nums, tt.want) {
			t.Errorf("rotateSecond() = %v, want %v", tt.nums, tt.want)
		}
	}
}

func BenchmarkRotate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
	}
}

func BenchmarkRotateSecond(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rotateSecond([]int{1, 2, 3, 4, 5, 6, 7}, 3)
	}
}
