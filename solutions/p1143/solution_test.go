package solution

import (
	"testing"
)

func TestLongestCommonSubsequence(t *testing.T) {
	type args struct {
		text1 string
		text2 string
	}
	tests := []struct {
		text1 string
		text2 string
		want  int
	}{
		{"abcde", "ace", 3},
		{"abc", "abc", 3},
		{"abc", "def", 0},
	}
	for _, tt := range tests {
		if got := longestCommonSubsequence(tt.text1, tt.text2); got != tt.want {
			t.Errorf("longestCommonSubsequence() = %v, want %v", got, tt.want)
		}
	}
}

func TestMax(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		a    int
		b    int
		want int
	}{
		{1, 2, 2},
		{2, 3, 3},
		{3, 5, 5},
		{5, 2, 5},
	}
	for _, tt := range tests {
		if got := max(tt.a, tt.b); got != tt.want {
			t.Errorf("max() = %v, want %v", got, tt.want)
		}
	}
}
