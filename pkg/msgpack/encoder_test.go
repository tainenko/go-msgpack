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

func TestEncoder_EncodeFloat64(t *testing.T) {
	type args struct {
		v float64
	}
	tests := []struct {
		name    string
		args    args
		wantBuf []byte
	}{
		{"Encode float, 1.1", args{1.1}, []byte{0xCB, 0x3F, 0xF1, 0x99, 0x99, 0x99, 0x99, 0x99, 0x9A}},
		{"Encode float, 5.1", args{5.1}, []byte{0xCB, 0x40, 0x14, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Encoder{}
			if gotBuf := e.EncodeFloat64(tt.args.v); !reflect.DeepEqual(gotBuf, tt.wantBuf) {
				t.Errorf("EncodeFloat64() = %v, want %v", gotBuf, tt.wantBuf)
			}
		})
	}
}

func TestEncoder_EncodeStr(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name    string
		args    args
		wantBuf []byte
	}{
		{"Encode fixed string, 'abcde'", args{"abcde"}, []byte{0xA5, 0x61, 0x62, 0x63, 0x64, 0x65}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			en := Encoder{}
			if gotBuf := en.EncodeStr(tt.args.v); !reflect.DeepEqual(gotBuf, tt.wantBuf) {
				t.Errorf("EncodeStr() = %v, want %v", gotBuf, tt.wantBuf)
			}
		})
	}
}

func TestEncoder_isFloat(t *testing.T) {
	type args struct {
		v float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1.1 is float", args{1.1}, true},
		{"1.0 is not float", args{1.0}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			en := Encoder{}
			if got := en.isFloat(tt.args.v); got != tt.want {
				t.Errorf("isFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}
