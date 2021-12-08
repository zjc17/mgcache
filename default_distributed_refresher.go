package mgcache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"time"
)

type (
	// PubSubClientInterface PubSubClient Interface
	PubSubClientInterface interface {
		Publish(ctx context.Context, channel string, message interface{}) *redis.IntCmd
		Subscribe(ctx context.Context, channels ...string) *redis.PubSub
	}

	defaultDistributedRefresher struct {
		pubSubClient  PubSubClientInterface
		storageClient IStorage
		topic         string
	}
)

// NewDefaultDistributedRefresher initializes the distributive cache refresher
func NewDefaultDistributedRefresher(topic string,
	redisClient PubSubClientInterface,
	storageClient IStorage) IDistributedRefresher {
	refreshEventChan := redisClient.Subscribe(context.Background(), topic).Channel()
	go func() {
		for key := range refreshEventChan {
			key := key
			go func() {
				if _, err := storageClient.Refresh(key.Payload); err != nil {
					zap.L().DPanic("failed to refresh review cache", zap.Error(err))
				}
			}()
		}
	}()
	return &defaultDistributedRefresher{
		pubSubClient:  redisClient,
		storageClient: storageClient,
		topic:         topic,
	}
}

func (d defaultDistributedRefresher) Notify(key string) (err error) {
	if err = d.storageClient.Invalidate(key); err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	_, err = d.pubSubClient.Publish(ctx, d.topic, key).Result()
	return
}
