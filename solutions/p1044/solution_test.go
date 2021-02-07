package solution

import "testing"

func Test_longestDupSubstring(t *testing.T) {
	tests := []struct {
		S    string
		want string
	}{
		{"banana", "ana"},
		{"abcd", ""},
	}
	for _, tt := range tests {
		if got := longestDupSubstring(tt.S); got != tt.want {
			t.Errorf("longestDupSubstring() = %v, want %v", got, tt.want)
		}
	}
}
