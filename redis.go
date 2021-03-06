package mgcache

import (
	"context"
	"github.com/go-redis/redis/v8"
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
		client         RedisClientInterface
		next           IFallbackStorage
		codec          ICodec
		timeToLive     time.Duration
		contextTimeout time.Duration

		storeType         string
		serviceIdentifier string

		metricCollector IMetricCollector
	}
)

// NewRedisStorage initializes the redisStorage
func NewRedisStorage(redisClient RedisClientInterface,
	next IFallbackStorage,
	opts ...OptionFunc) IStorage {

	opt := StorageOption{
		codec:             NewDefaultCodec(),
		timeToLive:        1 * time.Hour,
		contextTimeout:    100 * time.Millisecond,
		serviceIdentifier: defaultServiceIdentifier,
		metricCollector:   defaultMetricCollector,
	}
	for _, o := range opts {
		o(&opt)
	}

	return &redisStorage{
		client: redisClient,
		next:   next,

		codec:             opt.codec,
		timeToLive:        opt.timeToLive,
		contextTimeout:    opt.contextTimeout,
		serviceIdentifier: opt.serviceIdentifier,
		metricCollector:   opt.metricCollector,

		storeType: "redis",
	}
}

func (r redisStorage) GetBytes(key string) (bytes []byte, err error) {
	func() {
		ctx, cancel := context.WithTimeout(context.Background(), r.contextTimeout)
		defer cancel()
		bytes, err = r.client.Get(ctx, key).Bytes()
	}()

	switch err {
	case nil:
		// Cache Hit
		r.metricCollector.CacheHit(r.serviceIdentifier, r.storeType)
	case redis.Nil:
		if r.next == nil {
			return nil, ErrNil
		}
		bytes, err = r.Refresh(key)
		// Cache Miss
		r.metricCollector.CacheMiss(r.serviceIdentifier, r.storeType)
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
	ctx, cancel := context.WithTimeout(context.Background(), r.contextTimeout)
	defer cancel()

	var bytes []byte
	bytes, err = r.codec.Encode(value)

	err = r.client.Set(ctx, key, bytes, r.timeToLive).Err()
	if err != nil {
		// Cache Set
		r.metricCollector.CacheSet(r.serviceIdentifier, r.storeType)
		return
	}
	return
}

func (r redisStorage) Invalidate(key string) (err error) {

	func() {
		ctx, cancel := context.WithTimeout(context.Background(), r.contextTimeout)
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
