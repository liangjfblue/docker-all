package redis

import (
	"link-gin-db/config"
	"time"

	"github.com/gomodule/redigo/redis"
)

type IClient interface {
	Do(cmd, key string, arg ...interface{}) (interface{}, error)
	Expire(key string, arg ...interface{}) (interface{}, error)
}

type Pool struct {
	conf      *config.RedisConfig
	RedisPool *redis.Pool
}

func NewPool(conf *config.RedisConfig) *Pool {
	p := new(Pool)
	p.conf = conf
	p.initRedisPool()
	return p
}

func (p *Pool) initRedisPool() {
	p.RedisPool = &redis.Pool{
		MaxIdle:     p.conf.MaxIdle,                                  //最大的空闲连接数，表示即使没有redis连接时依然可以保持N个空闲的连接，而不被清除，随时处于待命状态
		MaxActive:   p.conf.MaxActive,                                //最大的激活连接数，表示同时最多有N个连接 ，为0事表示没有限制
		IdleTimeout: time.Duration(p.conf.IdleTimeout) * time.Second, //最大的空闲连接等待时间，超过此时间后，空闲连接将被关闭
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", p.conf.Host+":"+p.conf.Port)
			if err != nil {
				return nil, err
			}
			return c, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}
