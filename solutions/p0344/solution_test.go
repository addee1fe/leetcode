package solution

import (
	"reflect"
	"testing"
)

func TestReverseString(t *testing.T) {
	tests := []struct {
		s    []byte
		want []byte
	}{
		{[]byte{}, []byte{}},
		{[]byte{'a'}, []byte{'a'}},
		{[]byte{'a', 'b'}, []byte{'b', 'a'}},
		{[]byte{'h', 'e', 'l', 'l', 'o'}, []byte{'o', 'l', 'l', 'e', 'h'}},
		{[]byte{'H', 'a', 'n', 'n', 'a', 'h'}, []byte{'h', 'a', 'n', 'n', 'a', 'H'}},
	}
	for _, tt := range tests {
		reverseString(tt.s)
		if !reflect.DeepEqual(tt.s, tt.want) {
			t.Errorf("reverseString(%v) != %v", tt.s, tt.want)
		}
	}
}

func TestReverseStringOneIndex(t *testing.T) {
	tests := []struct {
		s    []byte
		want []byte
	}{
		{[]byte{}, []byte{}},
		{[]byte{'a'}, []byte{'a'}},
		{[]byte{'a', 'b'}, []byte{'b', 'a'}},
		{[]byte{'h', 'e', 'l', 'l', 'o'}, []byte{'o', 'l', 'l', 'e', 'h'}},
		{[]byte{'H', 'a', 'n', 'n', 'a', 'h'}, []byte{'h', 'a', 'n', 'n', 'a', 'H'}},
	}
	for _, tt := range tests {
		reverseStringOneIndex(tt.s)
		if !reflect.DeepEqual(tt.s, tt.want) {
			t.Errorf("reverseStringOneIndex(%v) != %v", tt.s, tt.want)
		}
	}
}

func BenchmarkReverseString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverseString([]byte{'h', 'e', 'l', 'l', 'o', 'm', 'y', 'n', 'a', 'm', 'e', 'i', 's'})
	}
}

func BenchmarkReverseStringOneIndex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverseStringOneIndex([]byte{'h', 'e', 'l', 'l', 'o', 'm', 'y', 'n', 'a', 'm', 'e', 'i', 's'})
	}
}
