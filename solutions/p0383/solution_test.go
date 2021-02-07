package solution

import "testing"

func TestCanConstruct(t *testing.T) {
	tests := []struct {
		ransomNote string
		magazine   string
		want       bool
	}{
		{"", "", true},
		{"a", "b", false},
		{"aa", "ab", false},
		{"aa", "aab", true},
		{"fffbfg", "effjfggbffjdgbjjhhdegh", true},
	}

	for _, tt := range tests {
		if got := canConstruct(tt.ransomNote, tt.magazine); got != tt.want {
			t.Errorf("canConstruct(%v, %v) = %v, want %v", tt.ransomNote, tt.magazine, got, tt.want)
		}
	}
}
