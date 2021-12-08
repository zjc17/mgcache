package mgcache

import (
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	mock "github.com/zjc17/mgcache/test/mock"
	"testing"
)

func TestNewRedisStorage(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	mockRedisClient := mock.NewMockRedisClientInterface(ctrl)

	// When
	redisStorageClient := NewRedisStorage(mockRedisClient, nil)

	// Then
	assert.IsType(t, new(redisStorage), redisStorageClient)
}

func TestGet(t *testing.T) {
	var (
		value []byte
		err error
		cmd *redis.StringCmd
		cacheKey = "cache-key"
		cacheValue = "cache-value"
	)

	// Given
	ctrl := gomock.NewController(t)
	mockRedisClient := mock.NewMockRedisClientInterface(ctrl)

	cmd = &redis.StringCmd{}
	cmd.SetVal(cacheValue)
	mockRedisClient.EXPECT().Get(gomock.Any(), cacheKey).Return(cmd)
	redisStorageClient := NewRedisStorage(mockRedisClient, nil)
	// When
	err = redisStorageClient.Get(cacheKey, &value)
	// Then
	assert.Nil(t, err)
	assert.Equal(t, []byte(cacheValue), value)

	// Given
	cmd = &redis.StringCmd{}
	cmd.SetErr(redis.Nil)
	mockRedisClient.EXPECT().Get(gomock.Any(), cacheKey).Return(cmd)
	// When
	value = nil
	err = redisStorageClient.Get(cacheKey, &value)
	// Then
	assert.Equal(t, ErrNil, err)
	assert.Nil(t, value)

	// Given
	expectedErr := errors.New("an unexpected error occurred")
	cmd = &redis.StringCmd{}
	cmd.SetErr(expectedErr)
	mockRedisClient.EXPECT().Get(gomock.Any(), cacheKey).Return(cmd)
	// When
	value = nil
	err = redisStorageClient.Get(cacheKey, &value)
	// Then
	assert.Equal(t, expectedErr, err)
	assert.Nil(t, value)
}

func TestGetWithNext(t *testing.T) {
	var (
		value []byte
		err error
		cmd *redis.StringCmd
		cacheKey = "cache-key"
		cacheValue = "cache-value"
	)

	// Given
	ctrl := gomock.NewController(t)
	mockNextStorage := mock.NewMockIFallbackStorage(ctrl)
	mockRedisClient := mock.NewMockRedisClientInterface(ctrl)
	redisStorageClient := NewRedisStorage(mockRedisClient, mockNextStorage)


	cmd = &redis.StringCmd{}
	cmd.SetErr(redis.Nil)
	mockRedisClient.EXPECT().Get(gomock.Any(), cacheKey).Return(cmd)
	mockNextStorage.EXPECT().GetBytes(cacheKey).Return([]byte(cacheValue), nil)
	mockRedisClient.EXPECT().Set(gomock.Any(), cacheKey, []byte(cacheValue), gomock.Any()).Return(&redis.StatusCmd{})

	// When
	value = nil
	err = redisStorageClient.Get(cacheKey, &value)
	// Then
	assert.Nil(t, err)
	assert.Equal(t, []byte(cacheValue), value)
}

func TestRedisInvalidate(t *testing.T) {
	var (
		err error
		cacheKey = "cache-key"
		cmd *redis.IntCmd
	)
	// Given
	ctrl := gomock.NewController(t)
	mockRedisClient := mock.NewMockRedisClientInterface(ctrl)
	redisStorageClient := NewRedisStorage(mockRedisClient, nil)
	cmd = &redis.IntCmd{}
	mockRedisClient.EXPECT().Del(gomock.Any(), cacheKey).Return(cmd)

	// When
	err = redisStorageClient.Invalidate(cacheKey)

	// Then
	assert.Nil(t, err)

	// Given
	expectedErr := errors.New("an unexpected error occurred")
	cmd = &redis.IntCmd{}
	cmd.SetErr(expectedErr)
	mockRedisClient.EXPECT().Del(gomock.Any(), cacheKey).Return(cmd)

	// When
	err = redisStorageClient.Invalidate(cacheKey)

	// Then
	assert.Equal(t, expectedErr, err)
}

func TestRedisInvalidateWithNext(t *testing.T) {
	var (
		err error
		cacheKey = "cache-key"
		cmd *redis.IntCmd
	)
	ctrl := gomock.NewController(t)
	mockNextStorage := mock.NewMockIFallbackStorage(ctrl)
	mockRedisClient := mock.NewMockRedisClientInterface(ctrl)
	redisStorageClient := NewRedisStorage(mockRedisClient, mockNextStorage)
	// Given
	cmd = &redis.IntCmd{}
	mockRedisClient.EXPECT().Del(gomock.Any(), cacheKey).Return(cmd)
	mockNextStorage.EXPECT().Invalidate(cacheKey).Return(nil)
	// When
	err = redisStorageClient.Invalidate(cacheKey)
	assert.Nil(t, err)
}

func TestRedisRefresh(t *testing.T) {
	var (
		err      error
		cacheKey = "cache-key"
		bytes []byte
	)
	ctrl := gomock.NewController(t)
	mockRedisClient := mock.NewMockRedisClientInterface(ctrl)
	redisStorageClient := NewRedisStorage(mockRedisClient, nil)
	// Given

	// When
	bytes, err =redisStorageClient.Refresh(cacheKey)
	// Then
	assert.Nil(t, bytes)
	assert.Equal(t, ErrRefreshUnsupported, err)
}

func TestRedisRefreshWithNext(t *testing.T) {
	var (
		err      error
		cacheKey = "cache-key"
		cacheValue = "cache-value"
		bytes []byte
		cmd *redis.StatusCmd
	)
	ctrl := gomock.NewController(t)
	mockNextStorage := mock.NewMockIFallbackStorage(ctrl)
	mockRedisClient := mock.NewMockRedisClientInterface(ctrl)
	redisStorageClient := NewRedisStorage(mockRedisClient, mockNextStorage)
	// Given
	cmd = &redis.StatusCmd{}
	mockNextStorage.EXPECT().GetBytes(cacheKey).Return([]byte(cacheValue), nil)
	mockRedisClient.EXPECT().Set(gomock.Any(), cacheKey, []byte(cacheValue), gomock.Any()).Return(cmd)
	// When
	bytes, err =redisStorageClient.Refresh(cacheKey)
	// Then
	assert.Nil(t, err)
	assert.Equal(t, []byte(cacheValue), bytes)

	// Given
	expectedErr := errors.New("an unexpected error occurred")
	//cmd = &redis.StatusCmd{}
	//cmd.SetErr(expectedErr)
	mockNextStorage.EXPECT().GetBytes(cacheKey).Return([]byte(cacheValue), expectedErr)
	// When
	bytes, err =redisStorageClient.Refresh(cacheKey)
	// Then
	assert.Equal(t, expectedErr, err)

}
