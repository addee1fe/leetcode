package solution

import "testing"

func Test_intToRoman(t *testing.T) {
	tests := []struct {
		num  int
		want string
	}{
		{
			3,
			"III",
		},
		{
			4,
			"IV",
		},
		{
			9,
			"IX",
		},
		{
			58,
			"LVIII",
		},
		{
			1994,
			"MCMXCIV",
		},
	}
	for _, tt := range tests {
		if got := intToRoman(tt.num); got != tt.want {
			t.Errorf("intToRoman() = %v, want %v", got, tt.want)
		}
	}
}
