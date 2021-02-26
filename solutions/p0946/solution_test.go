package solution

import "testing"

func TestValidateStackSequences(t *testing.T) {
	tests := []struct {
		pushed []int
		popped []int
		want   bool
	}{
		{
			[]int{1, 2, 3, 4, 5},
			[]int{4, 5, 3, 2, 1},
			true,
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]int{4, 3, 5, 1, 2},
			false,
		},
	}
	for _, tt := range tests {
		if got := validateStackSequences(tt.pushed, tt.popped); got != tt.want {
			t.Errorf("validateStackSequences() = %v, want %v", got, tt.want)
		}
	}
}
