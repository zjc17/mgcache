package codec

import (
	"github.com/vmihailenco/msgpack/v5"
)

type (
	defaultCodec struct {
	}
)

// NewDefaultCodec initializes the default codec
func NewDefaultCodec() ICodec {
	return &defaultCodec{}
}

// Encode encodes value pointed to by valuePtr
// into bytes []byte, which can be stored in cache
func (d defaultCodec) Encode(value interface{}) (bytes []byte, err error) {
	switch value.(type) {
	case []byte:
		bytes = value.([]byte)
		return
	}
	return msgpack.Marshal(value)
}

// Decode decodes bytes []byte from cache
// into value pointed to by valuePtr.
func (d defaultCodec) Decode(bytes []byte, valuePtr interface{}) (err error) {
	switch typ := valuePtr.(type) {
	case *[]byte:
		*typ = bytes
		return
	}
	return msgpack.Unmarshal(bytes, valuePtr)
}
