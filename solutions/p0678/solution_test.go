package solution

import "testing"

func TestCheckValidString(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"()", true},
		{"(((", false},
		{"(*)", true},
		{"**", true},
		{"***", true},
		{"(*))", true},
	}
	for _, tt := range tests {
		if got := checkValidString(tt.s); got != tt.want {
			t.Errorf("checkValidString() = %v, want %v", got, tt.want)
		}
	}
}
