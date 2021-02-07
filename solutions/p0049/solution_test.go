package solution

import (
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	got := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	t.Log(got)
}

func TestEncodeString(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"eat", "10001000000000000001000000"},
		{"tea", "10001000000000000001000000"},
		{"tan", "10000000000001000001000000"},
		{"ate", "10001000000000000001000000"},
		{"nat", "10000000000001000001000000"},
		{"bat", "11000000000000000001000000"},
	}
	for _, tt := range tests {
		if got := encodeString(tt.s); got != tt.want {
			t.Errorf("encodeString() = %v, want %v", got, tt.want)
		}

	}
}
