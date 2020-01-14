package user

import (
	"fmt"
	"link-gin-db/internal/controllers/base"
	"link-gin-db/internal/db/redis"
	"link-gin-db/internal/models"
	"link-gin-db/pkg/errno"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"
)

func (u *User) Create(c *gin.Context) {
	var (
		err    error
		result base.Result
		req    CreateRequest
	)

	if err = c.BindJSON(&req); err != nil {
		fmt.Println(err.Error())
		result.Failure(c, errno.ErrBind)
		return
	}

	if _, err := models.GetUser(&models.TBUser{Username: req.UserName}); err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			fmt.Println(err.Error())
			result.Failure(c, errno.ErrUserHadFound)
			return
		}
	}

	user := models.TBUser{
		Username:    req.UserName,
		Password:    req.UserPwd,
		Email:       req.UserEmail,
		Phone:       req.UserPhone,
		Sex:         req.Sex,
		Address:     req.Address,
		IsAvailable: req.IsAvailable,
		LastLogin:   time.Now(),
		LoginIp:     c.Request.Host,
	}

	if err := user.Validate(); err != nil {
		fmt.Println(err.Error())
		result.Failure(c, errno.ErrValidation)
		return
	}

	if err := user.Encrypt(); err != nil {
		fmt.Println(err.Error())
		result.Failure(c, errno.ErrEncrypt)
		return
	}

	if err := user.Create(); err != nil {
		fmt.Println(err.Error())
		result.Failure(c, errno.ErrDatabase)
		return
	}

	if _, err = u.Cache.Do("INCR", redis.CreateUserTotalKey); err != nil {
		fmt.Println(err.Error())
		result.Failure(c, errno.ErrDatabase)
		return
	}

	result.Success(c, nil)
}
