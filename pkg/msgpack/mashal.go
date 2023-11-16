package msgpack

import (
	"encoding/json"
	"fmt"
)

type Marshaller struct {
	buf []byte
}

func (m *Marshaller) Encode(v interface{}) {
	return
}

func (m *Marshaller) EncodeJSON(bin []byte) []byte {
	var obj interface{}
	if err := json.Unmarshal(bin, &obj); err != nil {
		panic(fmt.Sprintf("Error unmarshalling json: '%v'", err))
	}
	m.Encode(obj)
	return m.buf
}
