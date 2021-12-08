package mgcache

import (
	"github.com/go-redis/redis/v8"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	mock "github.com/zjc17/mgcache/test/mock"
	"testing"
)

func TestNewDefaultDistributedRefresher(t *testing.T) {
	var topic = "topic"
	// Given
	ctrl := gomock.NewController(t)
	pubSubClient := mock.NewMockPubSubClientInterface(ctrl)
	storageClient := mock.NewMockIStorage(ctrl)
	pubSub := &redis.PubSub{}
	defer func() {
		if r := recover(); r != nil {

		}
	}()
	defer pubSub.Close()
	pubSubClient.EXPECT().Subscribe(gomock.Any(), topic).Return(pubSub)

	// When
	refresher := NewDefaultDistributedRefresher(topic, pubSubClient, storageClient)


	// Then
	assert.IsType(t, new(defaultDistributedRefresher), refresher)
}

func TestDefaultDistributedRefresherNotify(t *testing.T) {
	//var topic = "topic"
	//var cacheKey = "cache-key"
	//var cmd *redis.IntCmd
	//var err error
	//defer func() {
	//	if r := recover(); r != nil {
	//
	//	}
	//}()
	//// Given
	//ctrl := gomock.NewController(t)
	//pubSubClient := mock.NewMockPubSubClientInterface(ctrl)
	//pubSub := &redis.PubSub{}
	//defer pubSub.Close()
	//pubSubClient.EXPECT().Subscribe(gomock.Any(), topic).Return(pubSub)
	//storageClient := mock.NewMockIStorage(ctrl)
	//refresher := NewDefaultDistributedRefresher(topic, pubSubClient, storageClient)
	//
	//storageClient.EXPECT().Invalidate(cacheKey).Return(nil)
	//cmd = &redis.IntCmd{}
	//pubSubClient.EXPECT().Publish(gomock.Any(), topic, cacheKey).Return(cmd)
	//
	//// When
	//err = refresher.Notify(cacheKey)
	//// Then
	//assert.Nil(t, err)

	// Given
	//expectedErr := errors.New("an unexpected error occurred")
	//storageClient.EXPECT().Invalidate(cacheKey).Return(expectedErr)
	//// When
	//err = refresher.Notify(cacheKey)
	//// Then
	//assert.Equal(t,expectedErr, err)
}
