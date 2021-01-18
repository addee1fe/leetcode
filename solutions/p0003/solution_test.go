package solution

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{
			"",
			0,
		},
		{
			"abcabcbb",
			3,
		},
		{
			"bbbbb",
			1,
		},
		{
			"pwwkew",
			3,
		},
	}
	for _, tt := range tests {
		if got := lengthOfLongestSubstring(tt.s); got != tt.want {
			t.Errorf("lengthOfLongestSubstring(%v) = %v, want %v", tt.s, got, tt.want)
		}
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		a, b int
		want int
	}{
		{0, 0, 0},
		{1, -1, 1},
		{-1, 1, 1},
		{-2, -1, -1},
	}
	for _, tt := range tests {
		if got := max(tt.a, tt.b); got != tt.want {
			t.Errorf("max() = %v, want %v", got, tt.want)
		}
	}
}
