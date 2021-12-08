package mgcache

type (
	// LoadFunc defines the function to load resource
	LoadFunc func(key string) (interface{}, error)

	defaultFallbackStorage struct {
		loadFunc LoadFunc
		codec    ICodec
	}
)

// NewDefaultFallbackStorage initializes the defaultFallbackStorage
func NewDefaultFallbackStorage(loadFunc LoadFunc,
	opts ...IStorageOption) IFallbackStorage {
	opt := options{
		codec: NewDefaultCodec(),
	}
	for _, o := range opts {
		o.apply(&opt)
	}
	return &defaultFallbackStorage{
		loadFunc: loadFunc,
		codec:    opt.codec,
	}
}

func (d defaultFallbackStorage) GetBytes(key string) (bytes []byte, err error) {
	var loadedValue interface{}
	loadedValue, err = d.loadFunc(key)
	return d.codec.Encode(loadedValue)
}

func (d defaultFallbackStorage) Invalidate(_ string) (err error) {
	return
}
