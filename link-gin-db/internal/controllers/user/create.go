package user

import (
	"link-gin-db/internal/controllers/base"
	"link-gin-db/internal/db/redis"
	"link-gin-db/internal/models"
	"link-gin-db/internal/pkg/errno"
	"log"
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
		log.Println(err.Error())
		result.Failure(c, errno.ErrBind)
		return
	}

	if _, err := models.GetUser(&models.TBUser{Username: req.UserName}); err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			log.Println(err.Error())
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
		log.Println(err.Error())
		result.Failure(c, errno.ErrValidation)
		return
	}

	if err := user.Encrypt(); err != nil {
		log.Println(err.Error())
		result.Failure(c, errno.ErrEncrypt)
		return
	}

	if err := user.Create(); err != nil {
		log.Println(err.Error())
		result.Failure(c, errno.ErrDatabase)
		return
	}

	if _, err = u.Cache.Incr(redis.CreateUserTotalKey); err != nil {
		log.Println("INCR: " + err.Error())
		result.Failure(c, errno.ErrDatabase)
		return
	}

	result.Success(c, nil)
}
