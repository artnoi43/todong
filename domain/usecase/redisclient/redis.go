package redisclient

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/go-redis/redis/v8"
)

func New(ctx context.Context, conf *Config) (redis.Cmdable, error) {
	j, _ := json.Marshal(conf)
	log.Printf("Redis client configuration:\n%s\n", j)
	cli := redis.NewClient(&redis.Options{DB: conf.DB})
	if cli != nil {
		return cli, nil
	}
	return nil, errors.New("nil redis client")
}
