package ginuser

import (
	"errors"
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/component/tokenprovider"
	userbusiness "food_delivery/module/user/business"
	usermodel "food_delivery/module/user/model"
	userstorage "food_delivery/module/user/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	ErrCannotGetPayload = common.NewCustomError(
		errors.New("error cannot get payload"),
		"error cannot get payload",
		"ErrCannotGetPayload")
)

func FindUser(ctx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		tokenPayload, ok := c.Get(common.TokenPayloadInJWTRequest)
		if !ok {
			panic(ErrCannotGetPayload)
		}
		payload := tokenPayload.(*tokenprovider.TokenPayload)
		var user usermodel.User
		user.Id = payload.UserId
		storage := userstorage.NewSQLStorage(ctx.GetMainDBConnection())
		business := userbusiness.NewFindUserBusiness(storage)
		userRes, err := business.Find(c.Request.Context(), &user)

		if err != nil {
			panic(err)
		}

		userRes.Mask()

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(userRes))
	}
}
