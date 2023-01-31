package middleware

import (
	"errors"
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/component/tokenprovider/jwt"
	"github.com/gin-gonic/gin"
	"strings"
)

var (
	ErrWrongAuthHeader = common.NewCustomError(errors.New("error wrong auth header"),
		"error wrong auth header",
		"ErrWrongAuthHeader")
)

func extractTokenFromHeaderString(s string) (string, error) {
	arr := strings.Split(s, " ") // Authorization "Bearer {token}"

	if arr[0] != "Bearer" || len(arr) < 2 {
		return "", ErrWrongAuthHeader
	}

	return arr[1], nil
}

func AuthenticateJWT(appCtx appctx.AppContext) func(c *gin.Context) {
	tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())
	return func(c *gin.Context) {
		token, err := extractTokenFromHeaderString(c.GetHeader("Authorization"))

		if err != nil {
			panic(err)
		}

		tokenPayload, err := tokenProvider.Validate(token)

		if err != nil {
			panic(err)
		}

		c.Set("tokenPayload", tokenPayload)
		c.Next()
	}
}
