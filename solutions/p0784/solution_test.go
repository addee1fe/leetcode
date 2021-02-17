package solution

import (
	"reflect"
	"testing"
)

func TestLetterCasePermutation(t *testing.T) {
	tests := []struct {
		S    string
		want []string
	}{
		{
			"a1b2",
			[]string{"a1b2", "A1b2", "A1B2", "a1B2"},
		},
		{
			"3z4",
			[]string{"3z4", "3Z4"},
		},
		{
			"12345",
			[]string{"12345"},
		},
		{
			"0",
			[]string{"0"},
		},
	}
	for _, tt := range tests {
		if got := letterCasePermutation(tt.S); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("letterCasePermutation() = %v, want %v", got, tt.want)
		}
	}
}
