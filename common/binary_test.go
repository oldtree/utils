package common

import (
	"testing"
)

func TestIsPowerOf2(t *testing.T) {
	type args struct {
		value uint64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPowerOf2(tt.args.value); got != tt.want {
				t.Errorf("IsPowerOf2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountBit1Number(t *testing.T) {
	type args struct {
		value uint64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountBit1Number(tt.args.value); got != tt.want {
				t.Errorf("CountBit1Number() = %v, want %v", got, tt.want)
			}
		})
	}
}
