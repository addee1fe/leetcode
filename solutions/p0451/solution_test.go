package solution

import "testing"

func TestFrequencySort(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"", ""},
		{"a", "a"},
		{"tree", "eert"},
		{"cccaaa", "aaaccc"},
		{"2a554442f544asfasssffffasss", "sssssssffffff44444aaaa55522"},
		{"Aabb", "bbAa"},
	}
	for _, tt := range tests {
		if got := frequencySort(tt.s); got != tt.want {
			t.Errorf("frequencySort() = %v, want %v", got, tt.want)
		}
	}
}

func TestFrequencySortMap(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		// {"", ""},
		// {"a", "a"},
		// {"tree", "eetr"},
		// {"cccaaa", "cccaaa"},
		// {"2a554442f544asfasssffffasss", "sssssssffffff44444aaaa55522"},
		// {"Aabb", "bbaA"},
	}
	for _, tt := range tests {
		if got := frequencySortMap(tt.s); got != tt.want {
			t.Errorf("frequencySort() = %v, want %v", got, tt.want)
		}
	}
}

func BenchmarkFrequencySortMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		frequencySortMap("2a554442f544asfasssffffasss")
	}
}

func BenchmarkFrequencySort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		frequencySort("2a554442f544asfasssffffasss")
	}
}
