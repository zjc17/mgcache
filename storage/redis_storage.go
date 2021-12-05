package storage

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/zjc17/mgcache/codec"
	"time"
)

type (
	// RedisClientInterface defines the functions used by redisStorage
	RedisClientInterface interface {
		Get(ctx context.Context, key string) *redis.StringCmd
		Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
		Del(ctx context.Context, keys ...string) *redis.IntCmd
	}

	redisStorage struct {
		client RedisClientInterface
		next   IFallbackStorage
		codec  codec.ICodec
	}
)

// NewRedisStorage initializes the redisStorage
func NewRedisStorage(redisClient RedisClientInterface, next IFallbackStorage) IStorage {
	return &redisStorage{
		client: redisClient,
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
		bytes, err = r.Refresh(key)
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
	return r.next.Invalidate(key)
}

func (r redisStorage) Refresh(key string) (bytes []byte, err error) {
	if r.next == nil {
		return nil, ErrRefreshUnsupported
	}
	if bytes, err = r.next.GetBytes(key); err != nil {
		return
	}
	err = r.Set(key, bytes)
	return
}
