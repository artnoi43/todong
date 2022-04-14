package store

import (
	"github.com/go-redis/redis/v8"
)

type redisDB interface {
	// redis.Client implements this interface
	redis.Cmdable
}
