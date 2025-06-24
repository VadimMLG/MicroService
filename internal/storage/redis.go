package storage

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client
var ctx = context.Background()

func InitRedis(addr string) {
	rdb = redis.NewClient(&redis.Options{
		Addr: addr,
	})
}

func CachePostLink(link string, postID uint) error {
	return rdb.Set(ctx, "post_link:"+link, postID, 24*time.Hour).Err()
}

func IsLinkCached(link string) (bool, error) {
	exists, err := rdb.Exists(ctx, "post_link:"+link).Result()
	return exists == 1, err
}
