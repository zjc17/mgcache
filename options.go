package mgcache

type (
	// IStorageOption options for IStorage
	IStorageOption interface {
		apply(opts *options)
	}

	options struct {
		codec ICodec
	}
)

func (o options) apply(opts *options) {
	opts.codec = o.codec
}

// WithCodec TODO
func WithCodec(c ICodec) IStorageOption {
	return options{codec: c}
}
