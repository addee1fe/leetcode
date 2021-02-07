package solution

import (
	"testing"
)

func TestStockSpannerNext(t *testing.T) {
	s := Constructor()
	tests := []struct {
		next int
		want int
	}{
		{100, 1},
		{80, 1},
		{60, 1},
		{70, 2},
		{60, 1},
		{75, 4},
		{85, 6},
	}
	for _, tt := range tests {
		if got := s.Next(tt.next); got != tt.want {
			t.Errorf("StockSpanner.Next() = %v, want %v", got, tt.want)
		}
	}
}
