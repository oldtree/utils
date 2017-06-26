package bitmap

import (
	"strconv"
	"sync"
	"testing"
)

func TestBitMap_Set(t *testing.T) {
	type fields struct {
		Data []uint64
		Size uint64
		sync.RWMutex
	}
	type args struct {
		value uint64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{"0", fields{Data: make([]uint64, 1), Size: 1}, args{value: 0}},
		{"1", fields{Data: make([]uint64, 1), Size: 1}, args{value: 1}},
		{"3", fields{Data: make([]uint64, 1), Size: 1}, args{value: 3}},
		{"4", fields{Data: make([]uint64, 1), Size: 1}, args{value: 4}},
		{"5", fields{Data: make([]uint64, 1), Size: 1}, args{value: 5}},
		{"16", fields{Data: make([]uint64, 1), Size: 1}, args{value: 16}},
		{"32", fields{Data: make([]uint64, 1), Size: 1}, args{value: 32}},
		{"63", fields{Data: make([]uint64, 2), Size: 2}, args{value: 63}},
		{"64", fields{Data: make([]uint64, 2), Size: 2}, args{value: 64}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BitMap{
				Data:    tt.fields.Data,
				Size:    tt.fields.Size,
				RWMutex: tt.fields.RWMutex,
			}
			b.Set(tt.args.value)
			Name, _ := strconv.Atoi(tt.name)
			if tt.fields.Data[uint64(Name)/64] != 1<<uint64(uint64(Name)%64) {
				t.Fail()
			}
		})
	}
}
