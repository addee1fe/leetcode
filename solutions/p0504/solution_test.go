package solution

import "testing"

func TestConvertToBase7(t *testing.T) {
	tests := []struct {
		num  int
		want string
	}{
		{
			100,
			"202",
		},
		{
			-7,
			"-10",
		},
	}
	for _, tt := range tests {
		if got := convertToBase7(tt.num); got != tt.want {
			t.Errorf("convertToBase7() = %v, want %v", got, tt.want)
		}
	}
}
