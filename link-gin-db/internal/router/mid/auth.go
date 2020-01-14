package mid

import (
	"link-gin-db/internal/controllers/base"
	"link-gin-db/pkg/errno"
	"link-gin-db/pkg/token"

	"github.com/gin-gonic/gin"
)

func AuthMid() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			err    error
			t      *token.Context
			result base.Result
		)

		if t, err = token.ParseRequest(c); err != nil {
			result.Failure(c, errno.ErrCheckToken)
			c.Abort()
			return
		}

		if t.UID == 0 {
			result.Failure(c, errno.ErrUserNotLogin)
			c.Abort()
			return
		}

		c.Next()
	}
}
