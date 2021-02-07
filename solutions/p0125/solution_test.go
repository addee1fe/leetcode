package solution

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{
			"A man, a plan, a canal: Panama",
			true,
		},
		{
			"race a car",
			false,
		},
	}
	for _, tt := range tests {
		if got := isPalindrome(tt.s); got != tt.want {
			t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
		}
	}
}

func TestClean(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{
			"A man, a plan, a canal: Panama",
			"amanaplanacanalpanama",
		},
		{
			"race a car",
			"raceacar",
		},
	}
	for _, tt := range tests {
		if got := clean(tt.s); got != tt.want {
			t.Errorf("clean() = %v, want %v", got, tt.want)
		}
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{
			"A man, a plan, a canal: Panama",
			"amanaP :lanac a ,nalp a ,nam A",
		},
		{
			"race a car",
			"rac a ecar",
		},
	}
	for _, tt := range tests {
		if got := reverse(tt.s); got != tt.want {
			t.Errorf("reverse() = %v, want %v", got, tt.want)
		}
	}
}
