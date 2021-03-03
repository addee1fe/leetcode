package solution

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		want FreqStack
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFreqStackPush(t *testing.T) {
	type fields struct {
		stacks [][]int
		freq   map[int]int
	}
	type args struct {
		x int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FreqStack{
				stacks: tt.fields.stacks,
				freq:   tt.fields.freq,
			}
			fs.Push(tt.args.x)
		})
	}
}

func TestFreqStackPop(t *testing.T) {
	type fields struct {
		stacks [][]int
		freq   map[int]int
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
			fs := &FreqStack{
				stacks: tt.fields.stacks,
				freq:   tt.fields.freq,
			}
			if got := fs.Pop(); got != tt.want {
				t.Errorf("FreqStack.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}
