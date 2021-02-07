package solution

import "testing"

func TestBackspaceCompare(t *testing.T) {
	tests := []struct {
		S    string
		T    string
		want bool
	}{
		{
			"ab#c",
			"ad#c",
			true,
		},
		{
			"ab##",
			"c#d#",
			true,
		},
		{
			"a##c",
			"#a#c",
			true,
		},
		{
			"a#c",
			"b",
			false,
		},
	}
	for _, tt := range tests {
		if got := backspaceCompare(tt.S, tt.T); got != tt.want {
			t.Errorf("backspaceCompare() = %v, want %v", got, tt.want)
		}
	}
}

func TestBuild(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{
			"ab#c",
			"ac",
		},
		{
			"ad#c",
			"ac",
		},
		{
			"ab##",
			"",
		},
		{
			"abc###",
			"",
		},
		{
			"a#b#c#",
			"",
		},
		{
			"a##c",
			"c",
		},
		{
			"#a#c",
			"c",
		},
		{
			"abc",
			"abc",
		},
	}
	for _, tt := range tests {
		if got := build(tt.s); got != tt.want {
			t.Errorf("build() = %v, want %v", got, tt.want)
		}
	}
}
