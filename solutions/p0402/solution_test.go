package solution

import "testing"

func TestRemoveKdigits(t *testing.T) {
	tests := []struct {
		num  string
		k    int
		want string
	}{
		{"12345", 0, "12345"},
		{"12345", -1, "12345"},
		{"10", 2, "0"},
		{"", 10, ""},
		{"1432219", 3, "1219"},
		{"10200", 1, "200"},
	}
	for _, tt := range tests {
		if got := removeKdigits(tt.num, tt.k); got != tt.want {
			t.Errorf("removeKdigits() = %v, want %v", got, tt.want)
		}
	}
}
