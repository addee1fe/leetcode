package solution

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want FirstUnique
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShowFirstUnique(t *testing.T) {
	tests := []struct {
		name string
		fu   *FirstUnique
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fu := &FirstUnique{}
			if got := fu.ShowFirstUnique(); got != tt.want {
				t.Errorf("FirstUnique.ShowFirstUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	type args struct {
		value int
	}
	tests := []struct {
		name string
		fu   *FirstUnique
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fu := &FirstUnique{}
			fu.Add(tt.args.value)
		})
	}
}
