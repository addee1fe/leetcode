package solution

import (
	"reflect"
	"testing"
)

func TestPickIndex(t *testing.T) {

}

func TestConstructor(t *testing.T) {
	type args struct {
		w []int
	}
	tests := []struct {
		name string
		args args
		want Solution
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.w); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolution_PickIndex(t *testing.T) {
	type fields struct {
		prefixSums []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Solution{
				prefixSums: tt.fields.prefixSums,
			}
			if got := s.PickIndex(); got != tt.want {
				t.Errorf("Solution.PickIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
