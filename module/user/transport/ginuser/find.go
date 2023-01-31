package ginuser

import (
	"errors"
	"food_delivery/common"
	"food_delivery/component/appctx"
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
		requester := c.MustGet(common.TokenPayloadInJWTRequest).(common.Requester)
		var user usermodel.User
		user.Id = requester.GetUserId()
		storage := userstorage.NewSQLStorage(ctx.GetMainDBConnection())
		business := userbusiness.NewFindUserBusiness(storage)
		userRes, err := business.Find(c.Request.Context(), &user)

		if err != nil {
			panic(err)
		}

		userRes.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(userRes))
	}
}
