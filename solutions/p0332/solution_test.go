package solution

import (
	"reflect"
	"testing"
)

func Test_findItinerary(t *testing.T) {
	tests := []struct {
		tickets [][]string
		want    []string
	}{
		{[][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}},
			[]string{"JFK", "MUC", "LHR", "SFO", "SJC"}},

		{[][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}},
			[]string{"JFK", "ATL", "JFK", "SFO", "ATL", "SFO"}},
	}
	for _, tt := range tests {
		if got := findItinerary(tt.tickets); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("findItinerary() = %v, want %v", got, tt.want)
		}
	}
}
