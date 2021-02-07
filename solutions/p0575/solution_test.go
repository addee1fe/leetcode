package solution

import "testing"

func Test_distributeCandies(t *testing.T) {
	tests := []struct {
		candies []int
		want    int
	}{
		{[]int{1, 1, 2, 2, 3, 3}, 3},
		{[]int{1, 1, 2, 3}, 2},
	}
	for _, tt := range tests {
		if got := distributeCandies(tt.candies); got != tt.want {
			t.Errorf("distributeCandies() = %v, want %v", got, tt.want)
		}
	}
}
