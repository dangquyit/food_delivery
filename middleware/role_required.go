package middleware

import (
	"errors"
	"food_delivery/common"
	"food_delivery/component/appctx"
	"github.com/gin-gonic/gin"
)

func RequiredRoles(appCtx appctx.AppContext, allRoles ...string) func(c *gin.Context) {
	return func(c *gin.Context) {
		tokenPayload := c.MustGet(common.TokenPayloadInJWTRequest).(common.Requester)

		for _, v := range allRoles {
			if tokenPayload.GetRole() == v {
				c.Next()
				return
				
			}
		}

		panic(common.ErrNoPermission(errors.New("user has been deleted or banned")))
	}
}
