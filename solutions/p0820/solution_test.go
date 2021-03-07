package solution

import "testing"

func TestMinimumLengthEncoding(t *testing.T) {
	tests := []struct {
		words []string
		want  int
	}{
		{
			[]string{"time", "me", "bell"},
			10,
		},
		{
			[]string{"t"},
			2,
		},
	}
	for _, tt := range tests {
		if got := minimumLengthEncoding(tt.words); got != tt.want {
			t.Errorf("minimumLengthEncoding() = %v, want %v", got, tt.want)
		}
	}
}
