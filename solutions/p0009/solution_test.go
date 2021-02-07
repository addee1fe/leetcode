package solution

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		x    int
		want bool
	}{
		{
			121,
			true,
		},
		{
			-121,
			false,
		},
		{
			10,
			false,
		},
		{
			-101,
			false,
		},
	}
	for _, tt := range tests {
		if got := isPalindrome(tt.x); got != tt.want {
			t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
		}
	}
}
