package solution

import "testing"

func TestReverseBits(t *testing.T) {
	tests := []struct {
		x    uint32
		want uint32
	}{
		{
			0b00000010100101000001111010011100,
			0b00111001011110000010100101000000,
		},
		{
			0b11111111111111111111111111111101,
			0b10111111111111111111111111111111,
		},
	}

	for _, tt := range tests {
		if got := reverseBits(tt.x); got != tt.want {
			t.Errorf("reverseBits() = %v, want %v", got, tt.want)
		}
	}
}

func TestReverseBitsLoop(t *testing.T) {
	tests := []struct {
		x    uint32
		want uint32
	}{
		{
			0b00000010100101000001111010011100,
			0b00111001011110000010100101000000,
		},
		{
			0b11111111111111111111111111111101,
			0b10111111111111111111111111111111,
		},
	}
	for _, tt := range tests {
		if got := reverseBits(tt.x); got != tt.want {
			t.Errorf("reverseBits() = %v, want %v", got, tt.want)
		}
	}
}

func BenchmarkReverseBits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverseBits(0b00000010100101000001111010011100)
	}
}

func BenchmarkReverseBitsLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverseBitsLoop(0b00000010100101000001111010011100)
	}
}
