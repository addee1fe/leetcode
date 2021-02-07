package solution

import (
	"testing"
)

func TestIsBadVersion(t *testing.T) {
	tests := []struct {
		version int
		want    bool
	}{
		{4, true},
	}
	for _, tt := range tests {
		if got := isBadVersion(tt.version); got != tt.want {
			t.Errorf("isBadVersion() = %v, want %v", got, tt.want)
		}
	}
}

func TestFirstBadVersion(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{5, 4},
	}
	for _, tt := range tests {
		if got := firstBadVersion(tt.n); got != tt.want {
			t.Errorf("firstBadVersion() = %v, want %v", got, tt.want)
		}
	}
}
