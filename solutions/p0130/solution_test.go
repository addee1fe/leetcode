package solution

import (
	"reflect"
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		board [][]byte
		want  [][]byte
	}{
		// X X X X
		// X O O X
		// X X O X
		// X O X X

		{[][]byte{}, [][]byte{}},
		{
			[][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'O', 'X'},
				{'X', 'X', 'O', 'X'},
				{'X', 'O', 'X', 'X'}},
			[][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'X', 'X'},
			}},
	}
	for _, tt := range tests {
		solve(tt.board)
		if !reflect.DeepEqual(tt.board, tt.want) {
			t.Errorf("solve(%v) !=\n %v", tt.board, tt.want)
		}
	}
}
