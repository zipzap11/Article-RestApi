package cache

import (
	"article/entity/response"
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisCacheImpl struct {
	Client  *redis.Client
	Host    string
	db      int
	expires time.Duration
	ctx     context.Context
}

func NewRedisCacheImpl(ctx context.Context, host string, db int, expires time.Duration) Cache {
	return &RedisCacheImpl{
		Host:    host,
		db:      db,
		expires: expires,
		ctx:     ctx,
		Client: redis.NewClient(&redis.Options{
			Addr:     host,
			Password: "",
			DB:       db,
		}),
	}
}

// func (cache *RedisCacheImpl) getClient() *redis.Client {
// 	return redis.NewClient(&redis.Options{
// 		Addr:     "localhost:6379",
// 		Password: "", // no password set
// 		DB:       0,
// 	})
// }

func (cache *RedisCacheImpl) Set(key string, value []response.ArticleGet) {
	client := cache.Client

	json, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}

	client.Set(cache.ctx, key, json, cache.expires)
}

func (cache *RedisCacheImpl) Get(key string) []response.ArticleGet {
	client := cache.Client

	val, err := client.Get(cache.ctx, key).Result()
	if err != nil {
		return nil
	}

	articles := []response.ArticleGet{}

	err = json.Unmarshal([]byte(val), &articles)
	if err != nil {
		panic(err)
	}

	return articles
}
