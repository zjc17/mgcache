package mgcache

import (
	"errors"
	"github.com/allegro/bigcache/v3"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	mock "github.com/zjc17/mgcache/test/mock"
	"testing"
)

func TestNewBigcacheStorage(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	mockBigCache := mock.NewMockBigcacheInterface(ctrl)

	// When
	bigcacheStorage := NewBigCacheStorage(mockBigCache, nil)

	// Then
	assert.IsType(t, *new(bigCacheStorage), bigcacheStorage)
}

func TestBigcacheGet(t *testing.T) {
	var defaultCodec = NewDefaultCodec()
	// Given
	ctrl := gomock.NewController(t)
	mockBigCache := mock.NewMockBigcacheInterface(ctrl)
	storage := NewBigCacheStorage(mockBigCache, nil)

	cacheKey := "cache-key"
	cacheValue := "cache value"

	mockBigCache.EXPECT().Get(cacheKey).Return(defaultCodec.Encode(cacheValue))

	// When
	var value string
	err := storage.Get(cacheKey, &value)

	// Then
	assert.Nil(t, err)
	assert.Equal(t, cacheValue, value)
}

func TestBigcacheWhenError(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	mockBigCache := mock.NewMockBigcacheInterface(ctrl)
	storage := NewBigCacheStorage(mockBigCache, nil)

	cacheKey := "cache-key"
	expectedErr := errors.New("an unexpected error occurred")

	mockBigCache.EXPECT().Get(cacheKey).Return(nil, expectedErr)

	// When
	var value string
	err := storage.Get(cacheKey, &value)

	// Then
	assert.Equal(t, expectedErr, err)
}

func TestBigcacheWithNotFoundErr(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	mockBigCache := mock.NewMockBigcacheInterface(ctrl)
	storage := NewBigCacheStorage(mockBigCache, nil)

	cacheKey := "cache-key"

	mockBigCache.EXPECT().Get(cacheKey).Return(nil, bigcache.ErrEntryNotFound)

	// When
	var value string
	err := storage.Get(cacheKey, &value)

	// Then
	assert.Equal(t, ErrNil, err)
}

func TestBigcacheWithNotFoundAndFallback(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	mockBigCache := mock.NewMockBigcacheInterface(ctrl)
	mockNextStorage := mock.NewMockIFallbackStorage(ctrl)
	storage := NewBigCacheStorage(mockBigCache, mockNextStorage)

	cacheKey := "cache-key"
	expectedErr := errors.New("an unexpected error occurred")

	mockBigCache.EXPECT().Get(cacheKey).Return(nil, bigcache.ErrEntryNotFound)
	mockNextStorage.EXPECT().GetBytes(cacheKey).Return(nil, expectedErr)

	// When
	var value string
	err := storage.Get(cacheKey, &value)

	// Then
	assert.Equal(t, expectedErr, err)
}

func TestBigcacheWithNotFoundAndFallbackSuccess(t *testing.T) {
	var defaultCodec = NewDefaultCodec()
	// Given
	ctrl := gomock.NewController(t)
	mockBigCache := mock.NewMockBigcacheInterface(ctrl)
	mockNextStorage := mock.NewMockIFallbackStorage(ctrl)
	storage := NewBigCacheStorage(mockBigCache, mockNextStorage)

	cacheKey := "cache-key"
	cacheValue := "cache-value"
	bytes, _ := defaultCodec.Encode(cacheValue)

	mockBigCache.EXPECT().Get(cacheKey).Return(nil, bigcache.ErrEntryNotFound)
	mockNextStorage.EXPECT().GetBytes(cacheKey).Return(bytes, nil)
	mockBigCache.EXPECT().Set(cacheKey, bytes).Return(nil)

	// When
	var value string
	err := storage.Get(cacheKey, &value)

	// Then
	assert.Nil(t, err)
	assert.Equal(t, cacheValue, value)
}

func TestBigcacheInvalidate(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	mockBigCache := mock.NewMockBigcacheInterface(ctrl)
	storage := NewBigCacheStorage(mockBigCache, nil)

	cacheKey := "cache-key"

	mockBigCache.EXPECT().Delete(cacheKey).Return(nil)

	// When
	err := storage.Invalidate(cacheKey)

	// Then
	assert.Nil(t, err)
}

func TestBigcacheInvalidateWhenErr(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	mockBigCache := mock.NewMockBigcacheInterface(ctrl)
	storage := NewBigCacheStorage(mockBigCache, nil)

	cacheKey := "cache-key"
	expectedErr := errors.New("an unexpected error occurred")

	mockBigCache.EXPECT().Delete(cacheKey).Return(expectedErr)

	// When
	err := storage.Invalidate(cacheKey)

	// Then
	assert.Equal(t, expectedErr, err)
}

func TestInvalidateWithFallbackStorage(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	mockBigCache := mock.NewMockBigcacheInterface(ctrl)
	mockNextStorage := mock.NewMockIFallbackStorage(ctrl)
	storage := NewBigCacheStorage(mockBigCache, mockNextStorage)

	cacheKey := "cache-key"
	mockBigCache.EXPECT().Delete(cacheKey).Return(nil)
	mockNextStorage.EXPECT().Invalidate(cacheKey).Return(nil)

	// When
	err := storage.Invalidate(cacheKey)

	// Then
	assert.Nil(t, err)
}

func TestInvalidateWithExpiredAndFallbackStorage(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	mockBigCache := mock.NewMockBigcacheInterface(ctrl)
	mockNextStorage := mock.NewMockIFallbackStorage(ctrl)
	storage := NewBigCacheStorage(mockBigCache, mockNextStorage)

	cacheKey := "cache-key"
	mockBigCache.EXPECT().Delete(cacheKey).Return(bigcache.ErrEntryNotFound)
	mockNextStorage.EXPECT().Invalidate(cacheKey).Return(nil)

	// When
	err := storage.Invalidate(cacheKey)

	// Then
	assert.Nil(t, err)
}

func TestRefresh(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	mockBigCache := mock.NewMockBigcacheInterface(ctrl)
	bigcacheStorage := NewBigCacheStorage(mockBigCache, nil)

	cacheKey := "cache-key"

	// When
	_, err := bigcacheStorage.Refresh(cacheKey)

	// Then
	assert.Equal(t, ErrRefreshUnsupported, err)
}

func TestRefreshWithFallbackStorage(t *testing.T) {
	var defaultCodec = NewDefaultCodec()
	// Given
	ctrl := gomock.NewController(t)
	mockBigCache := mock.NewMockBigcacheInterface(ctrl)
	mockNextStorage := mock.NewMockIFallbackStorage(ctrl)
	storage := NewBigCacheStorage(mockBigCache, mockNextStorage)

	cacheKey := "cache-key"
	cacheValue := "cache value"
	bytes, _ := defaultCodec.Encode(cacheValue)

	mockNextStorage.EXPECT().GetBytes(cacheKey).Return(bytes, nil)
	mockBigCache.EXPECT().Set(cacheKey, bytes).Return(nil)

	// When
	_, err := storage.Refresh(cacheKey)
	assert.Nil(t, err)
}

func TestRefreshWithFallbackStorageWhenErr(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	mockBigCache := mock.NewMockBigcacheInterface(ctrl)
	mockNextStorage := mock.NewMockIFallbackStorage(ctrl)
	storage := NewBigCacheStorage(mockBigCache, mockNextStorage)

	cacheKey := "cache-key"
	expectedErr := errors.New("an unexpected error occurred")

	mockNextStorage.EXPECT().GetBytes(cacheKey).Return(nil, expectedErr)

	// When
	_, err := storage.Refresh(cacheKey)
	assert.Equal(t, expectedErr, err)
}
