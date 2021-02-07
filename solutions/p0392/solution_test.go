package solution

import "testing"

func TestIsSubsequence(t *testing.T) {
	tests := []struct {
		s    string
		t    string
		want bool
	}{
		{"abc", "ahbgdc", true},
		{"", "abc", true},
		{"axc", "ahbgdc", false},
	}
	for _, tt := range tests {
		if got := isSubsequence(tt.s, tt.t); got != tt.want {
			t.Errorf("isSubsequence(%v, %v) = %v, want %v", tt.s, tt.t, got, tt.want)
		}
	}
}
