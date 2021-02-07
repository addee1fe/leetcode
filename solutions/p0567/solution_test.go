package solution

import (
	"testing"
)

func TestCheckInclusion(t *testing.T) {
	tests := []struct {
		s1   string
		s2   string
		want bool
	}{
		{"", "eidbaooo", true},
		{"ab", "", false},
		{"ab", "eidbaooo", true},
		{"ab", "eidboaoo", false},
	}
	for _, tt := range tests {
		if got := checkInclusion(tt.s1, tt.s2); got != tt.want {
			t.Errorf("checkInclusion() = %v, want %v", got, tt.want)
		}
	}
}

func TestMatch(t *testing.T) {
	tests := []struct {
		a1   [26]int
		a2   [26]int
		want bool
	}{
		{[26]int{1, 2, 3, 4, 5, 6}, [26]int{1, 2, 3, 4, 5, 6}, true},
		{[26]int{1, 2, 3, 4, 5, 6}, [26]int{1, 2, 3, 4, 5}, false},
	}
	for _, tt := range tests {
		if got := match(tt.a1, tt.a2); got != tt.want {
			t.Errorf("match() = %v, want %v", got, tt.want)
		}
	}
}
