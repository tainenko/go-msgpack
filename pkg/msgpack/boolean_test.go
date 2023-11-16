package msgpack

import (
	"reflect"
	"testing"
)

func TestBoolean_Encode(t *testing.T) {
	type fields struct {
		value bool
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{"Encode True", fields{true}, []byte{0xc3}},
		{"Encode False", fields{false}, []byte{0xc2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Boolean{
				value: tt.fields.value,
			}
			if got := b.Encode(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}
