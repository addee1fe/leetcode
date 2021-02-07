package solution

import "testing"

func TestFindCheapestPrice(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		n       int
		flights [][]int
		src     int
		dst     int
		K       int
		want    int
	}{
		{3, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, 0, 2, 1, 200},
		{3, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, 0, 2, 0, 500},
	}
	for _, tt := range tests {
		if got := findCheapestPrice(tt.n, tt.flights, tt.src, tt.dst, tt.K); got != tt.want {
			t.Errorf("findCheapestPrice() = %v, want %v", got, tt.want)
		}
	}
}
