package solution

import "testing"

func Test_isHappy(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		n    int
		want bool
	}{
		{19, true},
		{64, false},
		{23, true},
		{28, true},
		{128, false},
	}
	for _, tt := range tests {
		if got := isHappy(tt.n); got != tt.want {
			t.Errorf("isHappy() = %v, want %v", got, tt.want)
		}
	}
}
