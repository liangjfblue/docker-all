package user

import (
	"fmt"
	"link-gin-db/internal/controllers/base"
	"link-gin-db/internal/models"
	"link-gin-db/pkg/errno"
	"strconv"

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
		fmt.Println(err.Error())
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
