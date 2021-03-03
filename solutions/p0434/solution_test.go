package solution

import "testing"

// Example 1:
// Input: s = "Hello, my name is John"
// Output: 5
// Explanation: The five segments are ["Hello,", "my", "name", "is", "John"]

// Example 2:
// Input: s = "Hello"
// Output: 1

// Example 3:
// Input: s = "love live! mu'sic forever"
// Output: 4

// Example 4:
// Input: s = ""
// Output: 0

func TestCountSegments(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{
			"Hello, my name is John",
			5,
		},
		{
			"Hello",
			1,
		},
		{
			"love live! mu'sic forever",
			4,
		},
		{
			"",
			0,
		},
	}
	for _, tt := range tests {
		if got := countSegments(tt.s); got != tt.want {
			t.Errorf("countSegments() = %v, want %v", got, tt.want)
		}
	}
}
