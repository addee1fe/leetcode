package solution

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		s    string
		t    string
		want bool
	}{
		{
			"anagram",
			"nagaram",
			true,
		},
		{
			"car",
			"rat",
			false,
		},
	}
	for _, tt := range tests {
		if got := isAnagram(tt.s, tt.t); got != tt.want {
			t.Errorf("isAnagram() = %v, want %v", got, tt.want)
		}
	}
}
