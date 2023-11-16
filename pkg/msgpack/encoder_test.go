package msgpack

import (
	"reflect"
	"testing"
)

func TestEncoder_EncodeBool(t *testing.T) {
	type args struct {
		v bool
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"Encode True", args{true}, []byte{0xc3}},
		{"Encode False", args{false}, []byte{0xc2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			en := Encoder{}
			if got := en.EncodeBool(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EncodeBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncoder_EncodeNil(t *testing.T) {
	tests := []struct {
		name string
		want []byte
	}{
		{"Encode nil", []byte{0xc0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			en := Encoder{}
			if got := en.EncodeNil(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EncodeNil() = %v, want %v", got, tt.want)
			}
		})
	}
}
