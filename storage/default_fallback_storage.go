package storage

import "github.com/zjc17/mgcache/codec"

type (
	// LoadFunc defines the function to load resource
	LoadFunc func(key string) (interface{}, error)

	defaultFallbackStorage struct {
		loadFunc LoadFunc
		codec    codec.ICodec
	}
)

// NewDefaultFallbackStorage initializes the defaultFallbackStorage
func NewDefaultFallbackStorage(loadFunc LoadFunc) IFallbackStorage {
	return &defaultFallbackStorage{
		loadFunc: loadFunc,
		codec:    codec.NewDefaultCodec(),
	}
}

func (d defaultFallbackStorage) GetBytes(key string) (bytes []byte, err error) {
	var loadedValue interface{}
	loadedValue, err = d.loadFunc(key)
	return d.codec.Encode(loadedValue)
}

func (d defaultFallbackStorage) Invalidate(_ string) (err error) {
	return
}
