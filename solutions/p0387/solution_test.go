package solution

import "testing"

func TestFirstUniqChar(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		s    string
		want int
	}{
		{
			"leetcode",
			0,
		},
		{
			"loveleetcode",
			2,
		},
	}
	for _, tt := range tests {
		if got := firstUniqChar(tt.s); got != tt.want {
			t.Errorf("firstUniqChar(%v) = %v, want %v", tt.s, got, tt.want)
		}
	}
}
