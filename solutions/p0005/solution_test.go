package solution

import "testing"

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{
			"",
			"",
		},
		{
			"a",
			"a",
		},
		{
			"aa",
			"aa",
		},
		{
			"ab",
			"b",
		},
		{
			"babad",
			"aba",
		},
	}
	for _, tt := range tests {
		if got := longestPalindrome(tt.s); got != tt.want {
			t.Errorf("longestPalindrome(%v) = %v, want %v", tt.s, got, tt.want)
		}
	}
}
