package main

import (
	"github.com/liangjfblue/docker-all/link-gin-db/config"
	"github.com/liangjfblue/docker-all/link-gin-db/internal/db/redis"
	"github.com/liangjfblue/docker-all/link-gin-db/internal/router"
	"github.com/liangjfblue/docker-all/link-gin-db/internal/server"
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
