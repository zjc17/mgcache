package mgcache

import "time"

type (
	// OptionFunc option function for IStorage
	OptionFunc func(storageOption *StorageOption)

	// StorageOption hold options for storage client
	StorageOption struct {
		codec          ICodec
		timeToLive     time.Duration
		contextTimeout time.Duration
	}
)

// WithCodec customized codec for storage client
func WithCodec(c ICodec) OptionFunc {
	return func(storageOption *StorageOption) {
		storageOption.codec = c
	}
}

// WithTimeToLive customized time-to-live for storage client
func WithTimeToLive(d time.Duration) OptionFunc {
	return func(storageOption *StorageOption) {
		storageOption.timeToLive = d
	}
}

func WithContextTimeout(d time.Duration) OptionFunc {
	return func(storageOption *StorageOption) {
		storageOption.contextTimeout = d
	}
}
