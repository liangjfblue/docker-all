package redis

import (
	"link-gin-db/config"
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

func (c Client) Do(cmd, key string, arg ...interface{}) (interface{}, error) {
	if arg != nil {
		return c.pool.RedisPool.Get().Do(cmd, key, arg)
	} else {
		return c.pool.RedisPool.Get().Do(cmd, key)
	}
}

func (c Client) Expire(key string, arg ...interface{}) (interface{}, error) {
	return c.pool.RedisPool.Get().Do("EXPIRE", key, arg)
}
