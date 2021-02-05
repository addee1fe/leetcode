package solution

import "testing"

func TestSimplifyPath(t *testing.T) {
	tests := []struct {
		path string
		want string
	}{
		{
			"/home/",
			"/home",
		},
		{
			"/../",
			"/",
		},
		{
			"/home//foo/",
			"/home/foo",
		},
		{
			"/a/./b/../../c/",
			"/c",
		},
	}
	for _, tt := range tests {
		if got := simplifyPath(tt.path); got != tt.want {
			t.Errorf("simplifyPath(%v) = %v, want %v", tt.path, got, tt.want)
		}
	}
}
