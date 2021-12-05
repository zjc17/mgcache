package bigcache

import (
	"fmt"
	"github.com/allegro/bigcache/v3"
	"math"
	"testing"
	"time"
)

func BenchmarkBigcacheSet_Bytes(b *testing.B) {
	client, _ := bigcache.NewBigCache(bigcache.DefaultConfig(5 * time.Minute))
	store := NewBigCacheStorage(client, nil)
	for k := 0.; k <= 10; k++ {
		n := int(math.Pow(2, k))
		b.Run(fmt.Sprintf("%d", n), func(b *testing.B) {
			for i := 0; i < b.N*n; i++ {
				key := fmt.Sprintf("key-%d", n)
				value := []byte(fmt.Sprintf("value-%d", n))
				_ = store.Set(key, value)
			}
		})
	}
}

func BenchmarkBigcacheGetBytes(b *testing.B) {
	client, _ := bigcache.NewBigCache(bigcache.DefaultConfig(5 * time.Minute))
	store := NewBigCacheStorage(client, nil)

	key := "test"
	_ = store.Set(key, []byte("value"))

	for k := 0.; k <= 10; k++ {
		n := int(math.Pow(2, k))
		b.Run(fmt.Sprintf("%d", n), func(b *testing.B) {
			for i := 0; i < b.N*n; i++ {
				_, _ = store.GetBytes(key)
			}
		})
	}
}

func BenchmarkBigcacheGet(b *testing.B) {
	client, _ := bigcache.NewBigCache(bigcache.DefaultConfig(5 * time.Minute))
	store := NewBigCacheStorage(client, nil)

	key := "test"
	_ = store.Set(key, []byte("value"))
	var value []byte
	for k := 0.; k <= 10; k++ {
		n := int(math.Pow(2, k))
		b.Run(fmt.Sprintf("%d", n), func(b *testing.B) {
			for i := 0; i < b.N*n; i++ {
				_ = store.Get(key, &value)
			}
		})
	}
}
