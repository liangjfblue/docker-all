package redis

import (
	"github.com/liangjfblue/docker-all/link-gin-db/config"

	"github.com/gomodule/redigo/redis"
)

type Client struct {
	pool *Pool
}

var (
	CreateUserTotalKey = "docker_link_create_user_total_prefix"
)

func NewClient(conf *config.RedisConfig) *Client {
	return &Client{
		pool: NewPool(conf),
	}
}

func (c Client) GET(key string) (interface{}, error) {
	return c.pool.RedisPool.Get().Do("GET", key)
}

func (c Client) Incr(key string) (int64, error) {
	return redis.Int64(c.pool.RedisPool.Get().Do("INCR", key))
}

func (c Client) IncrBy(key string, increment int64) (int64, error) {
	return redis.Int64(c.pool.RedisPool.Get().Do("INCRBY", key, increment))
}
