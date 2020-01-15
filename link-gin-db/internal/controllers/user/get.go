package user

import (
	"log"
	"strconv"

	"github.com/liangjfblue/docker-all/link-gin-db/internal/controllers/base"
	"github.com/liangjfblue/docker-all/link-gin-db/internal/models"
	"github.com/liangjfblue/docker-all/link-gin-db/internal/pkg/errno"

	"github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"
)

func (u *User) Get(c *gin.Context) {
	var (
		err    error
		result base.Result
	)

	id := c.Param("uid")
	uid, _ := strconv.Atoi(id)

	user, err := models.GetUser(&models.TBUser{Model: gorm.Model{ID: uint(uid)}})
	if err != nil {
		log.Println(err.Error())
		result.Failure(c, errno.ErrUserNotFound)
		return
	}

	resp := GetResponse{
		UserName:    user.Username,
		UserEmail:   user.Email,
		UserPhone:   user.Phone,
		Sex:         user.Sex,
		Address:     user.Address,
		IsAvailable: user.IsAvailable,
	}

	result.Success(c, resp)
}
