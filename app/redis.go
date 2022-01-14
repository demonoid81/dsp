package app

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

type Redis struct {
	Redis *redis.Client
}

func (app *Env) Client() error {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", app.Cfg.RedisHost, app.Cfg.RedisPort),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return err
	}

	app.Redis.Redis = client
	return nil
}

func (r *Redis) Set(ctx context.Context, key string, val string) error {
	err := r.Redis.Set(ctx, key, val, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *Redis) Get(ctx context.Context, key string) (string, error) {
	value, err := r.Redis.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}
