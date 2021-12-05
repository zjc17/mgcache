package storage

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/zjc17/mgcache/codec"
	"time"
)

type (
	redisStorage struct {
		client redis.UniversalClient
		next   *IFallbackStorage
		codec  codec.ICodec
	}
)

func NewRedisStorage(client redis.UniversalClient, next *IFallbackStorage) IStorage {
	return &redisStorage{
		client: client,
		next:   next,
		codec:  codec.NewDefaultCodec(),
	}
}

func (r redisStorage) GetBytes(key string) (bytes []byte, err error) {

	func() {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
		defer cancel()
		bytes, err = r.client.Get(ctx, key).Bytes()
	}()

	switch err {
	case nil:
	case redis.Nil:
		if r.next == nil {
			return nil, ErrNil
		}
		if bytes, err = (*r.next).GetBytes(key); err != nil {
			return
		}
		err = r.Set(key, bytes)
	default:
		return
	}
	return
}

func (r redisStorage) Get(key string, valuePtr interface{}) (err error) {
	var bytes []byte
	if bytes, err = r.GetBytes(key); err != nil {
		return
	}
	return r.codec.Decode(bytes, valuePtr)
}

func (r redisStorage) Set(key string, value interface{}) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	var bytes []byte
	bytes, err = r.codec.Encode(value)

	err = r.client.Set(ctx, key, bytes, time.Hour).Err()
	return
}

func (r redisStorage) Invalidate(key string) (err error) {

	func() {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
		defer cancel()
		if err = r.client.Del(ctx, key).Err(); err != nil {
			return
		}
	}()

	// invalid next storage layer if exist
	if r.next == nil {
		return
	}
	return (*r.next).Invalidate(key)
}

func (r redisStorage) Refresh(key string) (err error) {
	if r.next == nil {
		return ErrRefreshUnsupported
	}
	var bytes []byte
	if err = (*r.next).Get(key, &bytes); err != nil {
		return
	}
	return r.Set(key, bytes)
}
