package solution

import "testing"

func TestMinRemoveToMakeValid(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{
			"lee(t(c)o)de)",
			"lee(t(c)o)de",
		},
		{
			"a)b(c)d",
			"ab(c)d",
		},
		{
			"))((",
			"",
		},
		{
			"(a(b(c)d)",
			"a(b(c)d)",
		},
	}
	for _, tt := range tests {
		if got := minRemoveToMakeValid(tt.s); got != tt.want {
			t.Errorf("minRemoveToMakeValid() = %v, want %v", got, tt.want)
		}
	}
}
