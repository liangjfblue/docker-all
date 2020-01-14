package user

import (
	"link-gin-db/config"
	"link-gin-db/internal/db/redis"
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
