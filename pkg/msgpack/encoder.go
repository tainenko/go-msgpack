package msgpack

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
	return
}

func (Encoder) isFloat(v float64) bool {
	return v == float64(int64(v))
}

func (Encoder) EncodeFloat64(v float64) (buf []byte) {
	return
}
