package solution

import "testing"

func TestDefangIPaddr(t *testing.T) {

	tests := []struct {
		address string
		want    string
	}{
		{
			"1.1.1.1",
			"1[.]1[.]1[.]1",
		},
		{
			"255.100.50.0",
			"255[.]100[.]50[.]0",
		},
	}
	for _, tt := range tests {
		if got := defangIPaddr(tt.address); got != tt.want {
			t.Errorf("defangIPaddr() = %v, want %v", got, tt.want)
		}
	}
}
