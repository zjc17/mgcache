package bigcache

import (
	"github.com/allegro/bigcache/v3"
	"github.com/zjc17/mgcache/codec"
	"github.com/zjc17/mgcache/storage"
)

type (
	BigcacheInterface interface {
		Get(key string) ([]byte, error)
		Set(key string, entry []byte) error
		Delete(key string) error
	}

	bigCacheStorage struct {
		client BigcacheInterface
		next   storage.IFallbackStorage
		codec  codec.ICodec
	}
)

func NewBigCacheStorage(client BigcacheInterface,
	next storage.IFallbackStorage) storage.IStorage {
	return bigCacheStorage{
		client: client,
		next:   next,
		codec:  codec.NewDefaultCodec(),
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
			return nil, storage.ErrNil
		}
		// TODO refactor and use refresh
		if bytes, err = b.next.GetBytes(key); err != nil {
			return
		}
		err = b.Set(key, bytes)
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

func (b bigCacheStorage) Refresh(key string) (err error) {
	if b.next == nil {
		return storage.ErrRefreshUnsupported
	}
	var bytes []byte
	if err = b.next.Get(key, &bytes); err != nil {
		return
	}
	return b.Set(key, bytes)
}
