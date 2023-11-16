package msgpack

import (
	"reflect"
	"testing"
)

func TestNil_Encode(t *testing.T) {
	tests := []struct {
		name string
		want []byte
	}{
		{"Encode nil", []byte{0xc0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := Nil{}
			if got := n.Encode(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}
