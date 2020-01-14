package controllers

import (
	"link-gin-db/config"
	"link-gin-db/internal/controllers/user"
	"link-gin-db/internal/db/redis"
)

type Handles struct {
	User *user.User
}

func NewHandles(config *config.Config, cache *redis.Client) *Handles {
	h := new(Handles)
	h.User = user.NewUserHandle(config, cache)
	return h
}
