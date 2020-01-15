package user

import (
	"github.com/liangjfblue/docker-all/link-gin-db/config"
	"github.com/liangjfblue/docker-all/link-gin-db/internal/db/redis"
)

type User struct {
	Config *config.Config
	Cache  *redis.Client
}

func NewUserHandle(config *config.Config, cache *redis.Client) *User {
	return &User{
		Config: config,
		Cache:  cache,
	}
}
