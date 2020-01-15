package user

import (
	"log"

	"github.com/liangjfblue/docker-all/link-gin-db/internal/controllers/base"
	"github.com/liangjfblue/docker-all/link-gin-db/internal/db/redis"
	"github.com/liangjfblue/docker-all/link-gin-db/internal/models"
	"github.com/liangjfblue/docker-all/link-gin-db/internal/pkg/errno"

	"github.com/gin-gonic/gin"
	redis2 "github.com/gomodule/redigo/redis"
)

func (u *User) LoginTotal(c *gin.Context) {
	var (
		err    error
		count  int64
		result base.Result
	)

	reply, err := redis2.Int64(u.Cache.GET(redis.CreateUserTotalKey))
	if err != nil {
		log.Println(err)

		count, err = models.CountUser()
		if err != nil {
			log.Println(err)
			result.Failure(c, errno.ErrDatabase)
			return
		}

		if _, err = u.Cache.IncrBy(redis.CreateUserTotalKey, count); err != nil {
			log.Println(err)
			result.Failure(c, errno.ErrCache)
			return
		}
		reply = count
	}

	resp := LoginTotalResponse{
		Total: uint(reply),
	}

	result.Success(c, resp)
}
