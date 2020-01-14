package main

import (
	"link-gin-db/config"
	"link-gin-db/internal/db/redis"
	"link-gin-db/internal/router"
	"link-gin-db/internal/server"
)

func main() {
	conf := config.NewConfig()

	cache := redis.NewClient(conf.RedisConf)

	r := router.NewRouter(conf, cache)
	r.Init()

	srv := server.NewServer(conf, r.G)
	srv.Init()
	srv.Start()
	defer srv.Stop()
}
