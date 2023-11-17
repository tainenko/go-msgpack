package msgpack

import (
	"encoding/binary"
	"math"
)

type Encoder struct{}

func (Encoder) EncodeNil() []byte {
	return []byte{NilFormat}
}

func (Encoder) EncodeBool(v bool) []byte {
	if v == true {
		return []byte{TrueFormat}
	} else {
		return []byte{FalseFormat}
	}
}

func (Encoder) EncodeStr(v string) (buf []byte) {
	length := len(v)
	if length < 32 {
		buf = make([]byte, 1+length)
		buf[0] = 0xa0 | byte(length)
		for i := 0; i < length; i++ {
			buf[i+1] = v[i]
		}
	} else if length <= 0xff {
		// str8, but only when not in compatibility mode
		buf = make([]byte, 2+length)
		buf[0] = 0xd9
		buf[1] = byte(length)
		for i := 0; i < length; i++ {
			buf[i+2] = v[i]
		}
	} else if length <= 0xffff {
		buf = make([]byte, 3+length)
		buf[0] = 0xda
		binary.BigEndian.PutUint16(buf[1:], uint16(length))
		for i := 0; i < length; i++ {
			buf[i+3] = v[i]
		}
	} else {
		buf = make([]byte, 5+length)
		buf[0] = 0xdb
		binary.BigEndian.PutUint32(buf[1:], uint32(length))
		for i := 0; i < length; i++ {
			buf[i+5] = v[i]
		}
	}
	return
}

func (Encoder) isFloat(v float64) bool {
	return v != float64(int64(v))
}

func (e Encoder) EncodeFloat64(v float64) (buf []byte) {
	buf = make([]byte, 9)
	buf[0] = Float64Format
	mem := math.Float64bits(v)
	binary.BigEndian.PutUint64(buf[1:], uint64(mem))
	return
}
