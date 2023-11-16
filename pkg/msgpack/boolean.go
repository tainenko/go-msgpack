package msgpack

type Boolean struct {
	value bool
}

func (b Boolean) Encode() []byte {
	if b.value == true {
		return []byte{TrueFormat}
	} else {
		return []byte{FalseFormat}
	}
}
