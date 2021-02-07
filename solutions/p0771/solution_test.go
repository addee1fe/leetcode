package solution

import "testing"

func TestNumJewelsInStones(t *testing.T) {
	tests := []struct {
		J    string
		S    string
		want int
	}{
		{"", "", 0},
		{"", "Ab", 0},
		{"aA", "aAAbbbb", 3},
		{"z", "ZZ", 0},
	}
	for _, tt := range tests {
		if got := numJewelsInStones(tt.J, tt.S); got != tt.want {
			t.Errorf("numJewelsInStones() = %v, want %v", got, tt.want)
		}
	}
}
func TestRuneNumJewelsInStones(t *testing.T) {
	tests := []struct {
		J    string
		S    string
		want int
	}{
		{"", "", 0},
		{"", "Ab", 0},
		{"aA", "aAAbbbb", 3},
		{"z", "ZZ", 0},
	}
	for _, tt := range tests {
		if got := runeNumJewelsInStones(tt.J, tt.S); got != tt.want {
			t.Errorf("runeNumJewelsInStones() = %v, want %v", got, tt.want)
		}
	}
}

func TestDummyNumJewelsInStones(t *testing.T) {
	tests := []struct {
		J    string
		S    string
		want int
	}{
		{"", "", 0},
		{"", "Ab", 0},
		{"aA", "aAAbbbb", 3},
		{"z", "ZZ", 0},
	}
	for _, tt := range tests {
		if got := dummyNumJewelsInStones(tt.J, tt.S); got != tt.want {
			t.Errorf("dummyNumJewelsInStones() = %v, want %v", got, tt.want)
		}
	}
}

func TestSliceNumJewelsInStones(t *testing.T) {
	tests := []struct {
		J    string
		S    string
		want int
	}{
		{"", "", 0},
		{"", "Ab", 0},
		{"aA", "aAAbbbb", 3},
		{"z", "ZZ", 0},
	}
	for _, tt := range tests {
		if got := sliceNumJewelsInStones(tt.J, tt.S); got != tt.want {
			t.Errorf("sliceNumJewelsInStones() = %v, want %v", got, tt.want)
		}
	}
}

func TestBitJewelsInStones(t *testing.T) {
	tests := []struct {
		J    string
		S    string
		want int
	}{
		{"", "", 0},
		{"", "Ab", 0},
		{"aA", "aAAbbbb", 3},
		{"z", "ZZ", 0},
	}
	for _, tt := range tests {
		if got := bitNumJewelsInStones(tt.J, tt.S); got != tt.want {
			t.Errorf("bitnumJewelsInStones() = %v, want %v", got, tt.want)
		}
	}
}

func BenchmarkNumJewelsInStones(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numJewelsInStones("asd", "asdjahdjkahuiqhjsbdaHJBAJDHAKJSDHaJSDNAHDBHJUWHNSDANNSBd")
	}
}

func BenchmarkRuneNumJewelsInStones(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runeNumJewelsInStones("asd", "asdjahdjkahuiqhjsbdaHJBAJDHAKJSDHaJSDNAHDBHJUWHNSDANNSBd")
	}
}

func BenchmarkDummyNumJewelsInStones(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dummyNumJewelsInStones("asd", "asdjahdjkahuiqhjsbdaHJBAJDHAKJSDHaJSDNAHDBHJUWHNSDANNSBd")
	}
}

func BenchmarkSliceNumJewelsInStones(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sliceNumJewelsInStones("asd", "asdjahdjkahuiqhjsbdaHJBAJDHAKJSDHaJSDNAHDBHJUWHNSDANNSBd")
	}
}

func BenchmarkBitNumJewelsInStones(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bitNumJewelsInStones("asd", "asdjahdjkahuiqhjsbdaHJBAJDHAKJSDHaJSDNAHDBHJUWHNSDANNSBd")
	}
}
