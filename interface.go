package mgcache

import (
	"errors"
)

type (
	// ICodec codec for cacheable objects
	ICodec interface {
		Encode(interface{}) ([]byte, error)
		Decode(bytes []byte, valuePtr interface{}) error
	}

	// IStorage defines a storage interface
	IStorage interface {
		IFallbackStorage
		// Set TODO
		Set(key string, value interface{}) (err error)
		// Refresh TODO
		Refresh(key string) (bytes []byte, err error)
		// Get loads the data into value pointed to by valuePtr
		Get(key string, value interface{}) (err error)
	}

	// IFallbackStorage is the cache in the layer lower than IStorage
	IFallbackStorage interface {
		// GetBytes returns the encoded value data stored in cache
		GetBytes(key string) (bytes []byte, err error)
		// Invalidate removes the cached key and propagates to the next layer of IStorage
		Invalidate(key string) (err error)
	}

	// IDistributedRefresher refreshes the cache for distributive scenarios' usage purposes.
	IDistributedRefresher interface {
		// Notify notifies all storage to refresh specific key
		Notify(key string) (err error)
	}
)

var (
	// ErrRefreshUnsupported occurs when the IStorage does not support Refresh function
	ErrRefreshUnsupported = errors.New("refresh unsupported due to no next storage set")
	// ErrNil occurs when get nil for the given key
	ErrNil = errors.New("no entry found for key")
)
