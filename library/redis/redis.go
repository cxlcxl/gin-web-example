package libredis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type Redis struct {
	*redis.Client
	Ctx    context.Context
	Prefix string
}

func NewRedis(addr, pwd, prefix string, db int) (r *Redis, err error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd,
		DB:       db,
		PoolSize: 20,
	})
	ctx := context.Background()
	if _, err = client.Ping(ctx).Result(); err != nil {
		return nil, err
	}

	return &Redis{
		Client: client,
		Ctx:    ctx,
		Prefix: prefix,
	}, nil
}
