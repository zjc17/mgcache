package mgcache

type (
	// EncodeFunc decode func
	EncodeFunc func(value interface{}) (bytes []byte, err error)
	// DecodeFunc encode func
	DecodeFunc func(bytes []byte, valuePtr interface{}) (err error)

	customCodec struct {
		encodeFunc EncodeFunc
		decodeFunc DecodeFunc
	}
)

// NewCustomCodec TODO
func NewCustomCodec(encodeFunc EncodeFunc, decodeFunc DecodeFunc) ICodec {
	return &customCodec{
		encodeFunc: encodeFunc,
		decodeFunc: decodeFunc,
	}
}

func (c customCodec) Encode(value interface{}) (bytes []byte, err error) {
	return c.encodeFunc(value)
}

func (c customCodec) Decode(bytes []byte, valuePtr interface{}) (err error) {
	return c.decodeFunc(bytes, valuePtr)
}
