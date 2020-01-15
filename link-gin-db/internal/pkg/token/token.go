package token

import (
	"errors"
	"link-gin-db/config"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Context struct {
	UID uint `json:"uid"`
}

var (
	secret     string
	secretTime int
)

func Init(conf *config.TokenConfig) {
	secret = conf.Secret
	secretTime = conf.SecretTime
}

func secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(secret), nil
	}
}

func Parse(tokenString string, secret string) (*Context, error) {
	ctx := &Context{}

	token, err := jwt.Parse(tokenString, secretFunc(secret))
	if err != nil {
		return ctx, err
	} else if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		ctx.UID = claims["uuid"].(uint)
		return ctx, nil
	} else {
		return ctx, err
	}
}

func ParseRequest(c *gin.Context) (*Context, error) {
	header := c.Request.Header.Get("Authorization")

	if len(header) == 0 {
		return &Context{}, errors.New("`Authorization` header is 0")
	}

	return Parse(header, secret)
}

func SignToken(c Context) (tokenString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uuid": c.UID,
		"nbf":  time.Now().Unix(),
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(time.Second * time.Duration(secretTime)).Unix(),
	})

	tokenString, err = token.SignedString([]byte(secret))
	return
}
