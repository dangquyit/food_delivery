package ginuser

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	userbusiness "food_delivery/module/user/business"
	usermodel "food_delivery/module/user/model"
	userstorage "food_delivery/module/user/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		storage := userstorage.NewSQLStorage(appCtx.GetMainDBConnection())
		business := userbusiness.NewCreateUserBusiness(storage)

		data.Role = "user"
		data.Mask()

		if err := business.CreateUser(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeID))
	}
}
