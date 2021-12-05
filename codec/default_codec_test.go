package codec

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestString(t *testing.T) {
	codec := NewDefaultCodec()
	value := "testing string value"
	bytes, err := codec.Encode(value)
	assert.Nil(t, err)
	var decodeValue string
	err = codec.Decode(bytes, &decodeValue)
	assert.Nil(t, err)
	assert.Equal(t, value, decodeValue)
}

func TestInt64(t *testing.T) {
	codec := NewDefaultCodec()
	var value int64 = 1
	bytes, err := codec.Encode(value)
	assert.Nil(t, err)

	var decodeValue int64
	err = codec.Decode(bytes, &decodeValue)
	assert.Nil(t, err)
	assert.Equal(t, value, decodeValue)
}

func TestMapStringString(t *testing.T) {
	codec := NewDefaultCodec()
	var value = map[string]string{"a": "b"}
	bytes, err := codec.Encode(value)
	assert.Nil(t, err)

	var decodeValue = make(map[string]string)
	err = codec.Decode(bytes, &decodeValue)
	assert.Nil(t, err)
	for k, v := range value {
		assert.Equal(t, v, decodeValue[k])
	}
}

func TestByteSlice(t *testing.T) {
	codec := NewDefaultCodec()
	var value = []byte{1, 2, 3}
	bytes, err := codec.Encode(value)
	assert.Nil(t, err)

	var decodeValue = make([]byte, 0)
	err = codec.Decode(bytes, &decodeValue)
	assert.Nil(t, err)
	assert.Equal(t, 3, len(decodeValue))
}
