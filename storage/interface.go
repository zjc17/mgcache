package storage

import "errors"

type (
	IStorage interface {
		IFallbackStorage
		// Set TODO
		Set(key string, value interface{}) (err error)
		// Refresh TODO
		Refresh(key string) (err error)
	}

	// IFallbackStorage is the cache in the layer lower than IStorage
	IFallbackStorage interface {
		// Get loads the data into value pointed to by valuePtr
		// TODO move to IStorage
		Get(key string, value interface{}) (err error)
		// GetBytes returns the encoded value data stored in cache
		GetBytes(key string) (bytes []byte, err error)
		// Invalidate removes the cached key and propagates to the next layer of IStorage
		Invalidate(key string) (err error)
	}
)

var (
	ErrRefreshUnsupported = errors.New("refresh unsupported due to no next storage set")
	ErrTypeUnsupported    = errors.New("type unsupported")
	ErrNil                = errors.New("no entry found for key")
)
