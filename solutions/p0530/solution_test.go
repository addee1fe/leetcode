package solution

import "testing"

func TestDetectCapitalUse(t *testing.T) {
	tests := []struct {
		word string
		want bool
	}{
		{"Ukraine", true},
		{"uSa", false},
		{"FlaG", false},
		{"leetcode", true},
		{"Microsoft", true},
	}
	for _, tt := range tests {
		if got := detectCapitalUse(tt.word); got != tt.want {
			t.Errorf("detectCapitalUse(%v) = %v, want %v", tt.word, got, tt.want)
		}
	}
}

func TestDetectCapitalUseLoop(t *testing.T) {
	tests := []struct {
		word string
		want bool
	}{
		{"Ukraine", true},
		{"uSa", false},
		{"FlaG", false},
		{"leetcode", true},
		{"Microsoft", true},
	}
	for _, tt := range tests {
		if got := detectCapitalUseLoop(tt.word); got != tt.want {
			t.Errorf("detectCapitalUseLoop(%v) = %v, want %v", tt.word, got, tt.want)
		}
	}
}

func TestDetectCapitalUseRegex(t *testing.T) {
	tests := []struct {
		word string
		want bool
	}{
		{"Ukraine", true},
		{"uSa", false},
		{"FlaG", false},
		{"leetcode", true},
		{"Microsoft", true},
	}
	for _, tt := range tests {
		if got := detectCapitalUseRegex(tt.word); got != tt.want {
			t.Errorf("detectCapitalUseRegex(%v) = %v, want %v", tt.word, got, tt.want)
		}
	}
}
func BenchmarkDetectCapitalUse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		detectCapitalUse("IAmTooBored")
	}
}

func BenchmarkDetectCapitalUseLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		detectCapitalUseLoop("IAmTooBored")
	}
}

func BenchmarkDetectCapitalUseRegex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		detectCapitalUseRegex("IAmTooBored")
	}
}
