package mgcache

import (
	"github.com/allegro/bigcache/v3"
)

type (
	// BigcacheInterface defines the functions used by bigCacheStorage
	BigcacheInterface interface {
		Get(key string) ([]byte, error)
		Set(key string, entry []byte) error
		Delete(key string) error
	}

	bigCacheStorage struct {
		client BigcacheInterface
		next   IFallbackStorage
		codec  ICodec
	}
)

// NewBigCacheStorage initializes the bigCacheStorage
func NewBigCacheStorage(client BigcacheInterface,
	next IFallbackStorage,
	opts ...IStorageOption) IStorage {
	opt := options{
		codec: NewDefaultCodec(),
	}
	for _, o := range opts {
		o.apply(&opt)
	}
	return bigCacheStorage{
		client: client,
		next:   next,
		codec:  opt.codec,
	}
}

func (b bigCacheStorage) Get(key string, valuePtr interface{}) (err error) {
	var bytes []byte
	if bytes, err = b.GetBytes(key); err != nil {
		return
	}
	return b.codec.Decode(bytes, valuePtr)
}

func (b bigCacheStorage) GetBytes(key string) (bytes []byte, err error) {
	bytes, err = b.client.Get(key)
	if err == bigcache.ErrEntryNotFound {
		if b.next == nil {
			return nil, ErrNil
		}
		bytes, err = b.Refresh(key)
	}
	return
}

func (b bigCacheStorage) Set(key string, value interface{}) (err error) {
	var bytes []byte
	bytes, err = b.codec.Encode(value)
	err = b.client.Set(key, bytes)
	return
}

func (b bigCacheStorage) Invalidate(key string) (err error) {
	err = b.client.Delete(key)
	if err == bigcache.ErrEntryNotFound {
		err = nil
	}
	if err != nil {
		return
	}
	// invalid next storage layer if exist
	if b.next == nil {
		return
	}
	return b.next.Invalidate(key)
}

func (b bigCacheStorage) Refresh(key string) (bytes []byte, err error) {
	if b.next == nil {
		return nil, ErrRefreshUnsupported
	}
	if bytes, err = b.next.GetBytes(key); err != nil {
		return
	}
	err = b.Set(key, bytes)
	return
}
