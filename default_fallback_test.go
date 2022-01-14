package mgcache

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestInvalidate(t *testing.T) {
	var loadFunc LoadFunc = func(key string) (interface{}, error) {
		return fmt.Sprintf("value_of_key[%s]", key), nil
	}
	defaultFallbackStorage := NewDefaultFallbackStorage(loadFunc, WithTimeToLive(1 * time.Second))
	if err := defaultFallbackStorage.Invalidate("a"); err != nil {
		t.Error("test failed")
	}
}

func TestGet_Success(t *testing.T) {
	var defaultFallbackStorage IFallbackStorage
	var loadFunc LoadFunc = func(key string) (interface{}, error) {
		return "value", nil
	}
	defaultFallbackStorage = NewDefaultFallbackStorage(loadFunc)

	bytes, err := defaultFallbackStorage.GetBytes("")
	assert.Nil(t, err)
	assert.NotNil(t, bytes)
}

func TestGet_Failed(t *testing.T) {
	var defaultFallbackStorage IFallbackStorage
	customErr := errors.New("custom error")
	var loadFunc LoadFunc = func(key string) (interface{}, error) {
		return nil, customErr
	}

	defaultFallbackStorage = NewDefaultFallbackStorage(loadFunc)

	_, err := defaultFallbackStorage.GetBytes("")
	assert.Equal(t, customErr, err)
}
