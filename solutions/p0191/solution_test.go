package solution

import "testing"

func TestHammingWeight(t *testing.T) {
	tests := []struct {
		x    uint32
		want int
	}{
		{
			0b00000000000000000000000000001011,
			3,
		},
		{
			0b00000000000000000000000010000000,
			1,
		},
		{
			0b11111111111111111111111111111101,
			31,
		},
	}
	for _, tt := range tests {
		if got := hammingWeight(tt.x); got != tt.want {
			t.Errorf("hammingWeight() = %v, want %v", got, tt.want)
		}
	}
}

func TestHammingWeightLoop(t *testing.T) {
	tests := []struct {
		num  uint32
		want int
	}{
		{
			0b00000000000000000000000000001011,
			3,
		},
		{
			0b00000000000000000000000010000000,
			1,
		},
		{
			0b11111111111111111111111111111101,
			31,
		},
	}
	for _, tt := range tests {
		if got := hammingWeightLoop(tt.num); got != tt.want {
			t.Errorf("hammingWeight() = %v, want %v", got, tt.want)
		}
	}
}

func BenchmarkHammingWeight(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hammingWeight(0b00000010100101000001111010011100)
	}
}

func BenchmarkHammingWeightLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hammingWeightLoop(0b00000010100101000001111010011100)
	}
}
