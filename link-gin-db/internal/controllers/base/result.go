package base

import (
	"link-gin-db/internal/pkg/errno"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func (r *Result) Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Result{
		Code: 1,
		Data: data,
		Msg:  "ok",
	})
}

func (r *Result) Failure(c *gin.Context, errno *errno.Errno) {
	c.JSON(http.StatusOK, Result{
		Code: 0,
		Data: map[string]interface{}{
			"code": errno.Code,
			"msg":  errno.Msg,
		},
		Msg: "error",
	})
}
