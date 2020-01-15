package mid

import (
	"github.com/liangjfblue/docker-all/link-gin-db/internal/controllers/base"
	"github.com/liangjfblue/docker-all/link-gin-db/internal/pkg/errno"
	"github.com/liangjfblue/docker-all/link-gin-db/internal/pkg/token"

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
