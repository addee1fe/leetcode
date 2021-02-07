package solution

import "testing"

func TestRangeBitwiseAnd(t *testing.T) {
	tests := []struct {
		a    int
		b    int
		want int
	}{
		{5, 7, 4},
		{0, 1, 0},
	}
	for _, tt := range tests {
		if got := rangeBitwiseAnd(tt.a, tt.b); got != tt.want {
			t.Errorf("rangeBitwiseAnd() = %v, want %v", got, tt.want)
		}
	}
}

func TestRangeBitwiseAndRecursive(t *testing.T) {
	tests := []struct {
		a    int
		b    int
		want int
	}{
		{5, 7, 4},
		{0, 1, 0},
	}
	for _, tt := range tests {
		if got := rangeBitwiseAndRecursive(tt.a, tt.b); got != tt.want {
			t.Errorf("rangeBitwiseAnd() = %v, want %v", got, tt.want)
		}
	}
}

func TestStolenRangeBitwiseAnd(t *testing.T) {
	tests := []struct {
		a    int
		b    int
		want int
	}{
		{5, 7, 4},
		{0, 1, 0},
	}
	for _, tt := range tests {
		if got := stolenRangeBitwiseAnd(tt.a, tt.b); got != tt.want {
			t.Errorf("rangeBitwiseAnd() = %v, want %v", got, tt.want)
		}
	}
}

func BenchmarkRangeBitwise(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rangeBitwiseAnd(1, 1000)
	}
}

func BenchmarkRangeBitwiseAndRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rangeBitwiseAndRecursive(1, 1000)
	}
}

func BenchmarkStolenRangeBitwise(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stolenRangeBitwiseAnd(1, 1000)
	}
}
