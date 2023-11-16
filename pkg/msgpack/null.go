package msgpack

type Nil struct {
}

func (n Nil) Encode() []byte {
	return []byte{0xc0}
}
