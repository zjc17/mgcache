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
