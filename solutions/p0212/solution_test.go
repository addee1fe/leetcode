package solution

import (
	"reflect"
	"testing"
)

func Test_findWords(t *testing.T) {
	tests := []struct {
		board [][]byte
		words []string
		want  []string
	}{
		{[][]byte{
			{'o', 'a', 'a', 'n'},
			{'e', 't', 'a', 'e'},
			{'i', 'h', 'k', 'r'},
			{'i', 'f', 'l', 'v'},
		},
			[]string{"oath", "pea", "eat", "rain"},
			[]string{"oath", "eat"},
		},
	}
	for _, tt := range tests {
		if got := findWords(tt.board, tt.words); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("findWords() = %v, want %v", got, tt.want)
		}
	}
}
