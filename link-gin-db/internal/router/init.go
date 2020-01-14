package router

import (
	"link-gin-db/config"
	"link-gin-db/internal/controllers"
	"link-gin-db/internal/db/redis"
	"link-gin-db/internal/router/mid"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Router struct {
	Config *config.Config
	G      *gin.Engine
	Cache  *redis.Client
}

func NewRouter(config *config.Config, cache *redis.Client) *Router {
	return &Router{
		Config: config,
		G:      gin.New(),
		Cache:  cache,
	}
}

func (r *Router) Init() {
	gin.SetMode(r.Config.HTTPConf.RunMode)

	r.G.Use(gin.Recovery())

	r.G.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route")
	})

	r.initRouter()
}

func (r *Router) initRouter() {
	handles := controllers.NewHandles(r.Config, r.Cache)

	r.G.POST("/v1/user", handles.User.Create)
	r.G.POST("/login", handles.User.Login)

	user := r.G.Group("/v1/user")
	user.Use(mid.AuthMid())
	{

		user.GET("/:uid", handles.User.Get)
	}

	f := r.G.Group("/v1/func")
	f.Use()
	{
		f.GET("/logintotal", handles.User.LoginTotal)
	}
}
