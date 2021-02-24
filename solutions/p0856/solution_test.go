package solution

import "testing"

func TestScoreOfParentheses(t *testing.T) {
	tests := []struct {
		S    string
		want int
	}{
		{
			"()",
			1,
		},
		{
			"(())",
			2,
		},
		{
			"()()",
			2,
		},
		{
			"(()(()))",
			6,
		},
	}
	for _, tt := range tests {
		if got := scoreOfParentheses(tt.S); got != tt.want {
			t.Errorf("scoreOfParentheses() = %v, want %v", got, tt.want)
		}
	}
}
