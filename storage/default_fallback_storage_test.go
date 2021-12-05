package storage

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInvalidate(t *testing.T) {
	var loadFunc LoadFunc = func(key string) (interface{}, error) {
		return fmt.Sprintf("value_of_key[%s]", key), nil
	}
	defaultFallbackStorage := NewDefaultFallbackStorage(loadFunc)
	if err := defaultFallbackStorage.Invalidate("a"); err != nil {
		t.Error("test failed")
	}
}

func TestGet_Failed(t *testing.T) {
	var defaultFallbackStorage IFallbackStorage
	var loadFunc LoadFunc = func(key string) (interface{}, error) {
		return "value", nil
	}
	defaultFallbackStorage = NewDefaultFallbackStorage(loadFunc)
	var value string
	err := defaultFallbackStorage.Get("", &value)
	assert.Equal(t, ErrTypeUnsupported, err)
}

func TestGet_Success(t *testing.T) {
	var defaultFallbackStorage IFallbackStorage
	var loadFunc LoadFunc = func(key string) (interface{}, error) {
		return "value", nil
	}
	defaultFallbackStorage = NewDefaultFallbackStorage(loadFunc)
	var value = make([]byte, 0)
	if err := defaultFallbackStorage.Get("", &value); err != nil {
		t.Error("test failed")
	}
	assert.NotNil(t, value)
}
