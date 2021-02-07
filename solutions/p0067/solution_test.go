package solution

import "testing"

func TestAddBinary(t *testing.T) {
	tests := []struct {
		a    string
		b    string
		want string
	}{
		{"", "", ""},
		{"11", "1", "100"},
		{"1010", "1011", "10101"},
	}
	for _, tt := range tests {
		if got := addBinary(tt.a, tt.b); got != tt.want {
			t.Errorf("addBinary() = %v, want %v", got, tt.want)
		}
	}
}
