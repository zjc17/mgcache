package storage

import "github.com/zjc17/mgcache/codec"

type (
	LoadFunc               func(key string) (interface{}, error)
	defaultFallbackStorage struct {
		loadFunc LoadFunc
		codec    codec.ICodec
	}
)

func NewDefaultFallbackStorage(loadFunc LoadFunc) IFallbackStorage {
	return &defaultFallbackStorage{
		loadFunc: loadFunc,
		codec:    codec.NewDefaultCodec(),
	}
}

func (d defaultFallbackStorage) Get(key string, valuePtr interface{}) (err error) {
	var loadedValue interface{}
	loadedValue, err = d.loadFunc(key)
	switch typ := valuePtr.(type) {
	case *[]byte:
		var bytes []byte
		bytes, err = d.codec.Encode(loadedValue)
		*typ = bytes
	default:
		err = ErrTypeUnsupported
	}
	return
}

func (d defaultFallbackStorage) GetBytes(key string) (bytes []byte, err error) {
	var loadedValue interface{}
	loadedValue, err = d.loadFunc(key)
	return d.codec.Encode(loadedValue)
}

func (d defaultFallbackStorage) Invalidate(_ string) (err error) {
	return
}
