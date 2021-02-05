package solution

import (
	"reflect"
	"testing"
)

func TestKidsWithCandies(t *testing.T) {
	tests := []struct {
		candies      []int
		extraCandies int
		want         []bool
	}{
		{[]int{2, 3, 5, 1, 3}, 3, []bool{true, true, true, false, true}},
		{[]int{4, 2, 1, 1, 2}, 1, []bool{true, false, false, false, false}},
		{[]int{12, 1, 12}, 10, []bool{true, false, true}},
	}
	for _, tt := range tests {
		if got := kidsWithCandies(tt.candies, tt.extraCandies); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("kidsWithCandies() = %v, want %v", got, tt.want)
		}
	}
}
