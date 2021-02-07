package solution

import "testing"

func Test_numIslands(t *testing.T) {
	tests := []struct {
		grid [][]byte
		want int
	}{
		{[][]byte{
			[]byte{'1', '1', '1', '1', '0'},
			[]byte{'1', '1', '0', '1', '0'},
			[]byte{'1', '1', '0', '0', '0'},
			[]byte{'0', '0', '0', '0', '0'}},
			1},
		{[][]byte{
			[]byte{'1', '1', '0', '0', '0'},
			[]byte{'1', '1', '0', '0', '0'},
			[]byte{'0', '0', '1', '0', '0'},
			[]byte{'0', '0', '0', '1', '1'}},
			3},
	}
	for _, tt := range tests {
		if got := numIslands(tt.grid); got != tt.want {
			t.Errorf("numIslands() = %v, want %v", got, tt.want)
		}
	}
}
