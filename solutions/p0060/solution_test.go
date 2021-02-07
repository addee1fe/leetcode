package solution

import "testing"

func Test_getPermutation(t *testing.T) {
	tests := []struct {
		n    int
		k    int
		want string
	}{
		{3, 3, "213"},
		{4, 9, "2314"},
	}
	for _, tt := range tests {
		if got := getPermutation(tt.n, tt.k); got != tt.want {
			t.Errorf("getPermutation() = %v, want %v", got, tt.want)
		}
	}
}
