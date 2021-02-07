package solution

import "testing"

func TestValidUtf8(t *testing.T) {
	tests := []struct {
		data []int
		want bool
	}{
		{
			[]int{197, 130, 1},
			true,
		},
		{
			[]int{235, 140, 4},
			false,
		},
		{
			[]int{235, 140, 4},
			false,
		},
	}
	for _, tt := range tests {
		if got := validUtf8(tt.data); got != tt.want {
			t.Errorf("validUtf8() = %v, want %v", got, tt.want)
		}
	}
}
