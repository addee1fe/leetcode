package solution

import "testing"

func TestFindLongestWord(t *testing.T) {
	tests := []struct {
		s    string
		d    []string
		want string
	}{
		{
			"abpcplea",
			[]string{"ale", "apple", "monkey", "plea"},
			"apple",
		},
		{
			"abpcplea",
			[]string{"a", "b", "c"},
			"a",
		},
	}
	for _, tt := range tests {
		if got := findLongestWord(tt.s, tt.d); got != tt.want {
			t.Errorf("findLongestWord() = %v, want %v", got, tt.want)
		}
	}
}

func TestIsSubsequence(t *testing.T) {
	tests := []struct {
		x    string
		y    string
		want bool
	}{
		{
			"",
			"",
			true,
		},
		{
			"a",
			"abpcplea",
			true,
		},
		{
			"ale",
			"abpcplea",
			true,
		},
		{
			"apple",
			"abpcplea",
			true,
		},
		{
			"plea",
			"abpcplea",
			true,
		},
		{
			"monkey",
			"abpcplea",
			false,
		},
	}
	for _, tt := range tests {
		if got := isSubsequence(tt.x, tt.y); got != tt.want {
			t.Errorf("isSubsequence() = %v, want %v", got, tt.want)
		}
	}
}
