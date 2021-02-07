package solution

import (
	"testing"
)

func TestMaximalSquare(t *testing.T) {
	type args struct {
		matrix [][]byte
	}
	tests := []struct {
		matrix [][]byte
		want   int
	}{
		{[][]byte{}, 0},
		{[][]byte{
			{1, 0, 1, 0, 1},
			{1, 0, 1, 0, 1},
			{1, 0, 1, 0, 1},
			{1, 0, 1, 0, 1},
			{1, 0, 1, 0, 1},
			{1, 0, 1, 0, 1},
		}, 1},
		{[][]byte{
			{1, 0, 1, 0, 0},
			{1, 0, 1, 1, 1},
			{1, 1, 1, 1, 1},
			{1, 0, 0, 1, 0},
		}, 4},
		{[][]byte{
			{1, 1, 1, 0, 0},
			{1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1},
			{1, 0, 0, 1, 0},
		}, 9},
		{[][]byte{
			{1, 1, 1, 1, 0},
			{1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1},
			{1, 1, 1, 1, 0},
		}, 16},
	}
	for _, tt := range tests {
		if got := maximalSquare(tt.matrix); got != tt.want {
			t.Errorf("maximalSquare() = %v, want %v", got, tt.want)
		}
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		a, b int
		want int
	}{
		{1, 2, 1},
		{2, 3, 2},
		{3, 4, 3},
		{4, 5, 4},
		{7, 6, 6},
		{10, 8, 8},
	}
	for _, tt := range tests {
		if got := min(tt.a, tt.b); got != tt.want {
			t.Errorf("min() = %v, want %v", got, tt.want)
		}
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		a, b int
		want int
	}{
		{1, 2, 2},
		{2, 3, 3},
		{3, 4, 4},
		{4, 5, 5},
		{7, 6, 7},
		{10, 8, 10},
	}
	for _, tt := range tests {
		if got := max(tt.a, tt.b); got != tt.want {
			t.Errorf("max() = %v, want %v", got, tt.want)
		}
	}
}
