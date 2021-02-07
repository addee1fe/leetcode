package solution

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	tests := []struct {
		s    string
		p    string
		want []int
	}{
		{
			"cbaebabacd",
			"",
			[]int{},
		},
		{
			"",
			"abc",
			[]int{},
		},
		{
			"cbaebabacd",
			"abc",
			[]int{0, 6},
		},
		{
			"abab",
			"ab",
			[]int{0, 1, 2},
		},
	}
	for _, tt := range tests {
		if got := findAnagrams(tt.s, tt.p); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("findAnagrams() = %v, want %v", got, tt.want)
		}

	}
}
