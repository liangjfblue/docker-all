package user

import (
	"fmt"
	"link-gin-db/internal/controllers/base"
	"link-gin-db/internal/db/redis"
	"link-gin-db/pkg/errno"

	redis2 "github.com/gomodule/redigo/redis"

	"github.com/gin-gonic/gin"
)

func (u *User) LoginTotal(c *gin.Context) {
	var (
		err    error
		result base.Result
	)

	reply, err := redis2.Int(u.Cache.Do("GET", redis.CreateUserTotalKey))
	if err != nil {
		fmt.Println(err.Error())
		result.Failure(c, errno.ErrDatabase)
		return
	}

	resp := LoginTotalResponse{
		Total: uint(reply),
	}

	result.Success(c, resp)

}
