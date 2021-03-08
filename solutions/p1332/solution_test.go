package solution

import "testing"

func TestTemovePalindromeSub(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{
			"",
			0,
		},
		{
			"ababa",
			1,
		},
		{
			"abb",
			2,
		},
		{
			"baabb",
			2,
		},
	}
	for _, tt := range tests {
		if got := removePalindromeSub(tt.s); got != tt.want {
			t.Errorf("removePalindromeSub() = %v, want %v", got, tt.want)
		}
	}
}
