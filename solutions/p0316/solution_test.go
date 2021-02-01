package solution

import "testing"

func TestRemoveDuplicateLetters(t *testing.T) {
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
			"bcabc",
			"abc",
		},
		{
			"cbacdcbc",
			"acdb",
		},
	}
	for _, tt := range tests {
		if got := removeDuplicateLetters(tt.s); got != tt.want {
			t.Errorf("removeDuplicateLetters(%v) = %v, want %v", tt.s, got, tt.want)
		}
	}
}
