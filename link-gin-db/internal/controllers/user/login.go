package user

import (
	"fmt"
	"link-gin-db/internal/controllers/base"
	"link-gin-db/internal/models"
	"link-gin-db/pkg/auth"
	"link-gin-db/pkg/errno"
	"link-gin-db/pkg/token"

	"github.com/gin-gonic/gin"
)

func (u *User) Login(c *gin.Context) {
	var (
		err    error
		result base.Result
		req    LoginRequest
	)

	if err = c.BindJSON(&req); err != nil {
		fmt.Println(err.Error())
		result.Failure(c, errno.ErrBind)
		return
	}

	user, err := models.GetUser(&models.TBUser{Username: req.Username})
	if err != nil {
		fmt.Println(err.Error())
		result.Failure(c, errno.ErrUserNotFound)
		return
	}

	hashPassword, err := auth.Encrypt(req.Password)
	if err != nil {
		fmt.Println(err.Error())
		result.Failure(c, errno.ErrEncrypt)
		return

	}

	if err = auth.Compare(hashPassword, user.Password); err != nil {
		fmt.Println(err.Error())
		result.Failure(c, errno.ErrCompare)
		return
	}

	tokenStr, err := token.SignToken(token.Context{UID: user.ID})
	if err != nil {
		fmt.Println(err.Error())
		result.Failure(c, errno.ErrSignToken)
		return
	}

	resp := LoginResponse{
		UID:   user.ID,
		Token: tokenStr,
	}
	result.Success(c, resp)
}
