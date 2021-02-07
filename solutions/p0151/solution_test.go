package solution

import "testing"

func TestReverseWords(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{" ", ""},
		{"the sky is blue", "blue is sky the"},
		{"  hello world!  ", "world! hello"},
		{"a good   example", "example good a"},
	}
	for _, tt := range tests {
		if got := reverseWords(tt.s); got != tt.want {
			t.Errorf("reverseWords() = %v, want %v", got, tt.want)
		}
	}
}
