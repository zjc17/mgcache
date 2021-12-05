package codec

type (
	// ICodec codec for cacheable objects
	ICodec interface {
		Encode(interface{}) ([]byte, error)
		Decode(bytes []byte, valuePtr interface{}) error
	}
)
