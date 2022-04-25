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

		// serviceIdentifier identifier for prometheus cacheBehaviorMetric
		serviceIdentifier string

		metricCollector IMetricCollector
	}
)

const (
	defaultServiceIdentifier = "mgcache"
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

// WithContextTimeout customized context timeout for storage client
func WithContextTimeout(d time.Duration) OptionFunc {
	return func(storageOption *StorageOption) {
		storageOption.contextTimeout = d
	}
}

// WithMetricCollector customized metric collector for storage client
// if not set, default metric collector defaultMetricCollector will be used
func WithMetricCollector(c IMetricCollector) OptionFunc {
	return func(storageOption *StorageOption) {
		storageOption.metricCollector = c
	}
}

// WithServiceIdentifier customized service identifier for storage client
func WithServiceIdentifier(s string) OptionFunc {
	return func(storageOption *StorageOption) {
		storageOption.serviceIdentifier = s
	}
}
