package solution

import "testing"

func TestRotateString(t *testing.T) {
	tests := []struct {
		A    string
		B    string
		want bool
	}{
		{"abcde", "cdeab", true},
		{"abcde", "abced", false},
	}
	for _, tt := range tests {
		if got := rotateString(tt.A, tt.B); got != tt.want {
			t.Errorf("rotateString() = %v, want %v", got, tt.want)
		}
	}
}
