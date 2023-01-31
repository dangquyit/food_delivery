package ginuser

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/component/hasher"
	"food_delivery/component/tokenprovider/jwt"
	userbusiness "food_delivery/module/user/business"
	usermodel "food_delivery/module/user/model"
	userstorage "food_delivery/module/user/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data usermodel.UserLogin

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		storage := userstorage.NewSQLStorage(appCtx.GetMainDBConnection())
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())
		md5 := hasher.NewMd5Hash()
		business := userbusiness.NewLoginBusiness(storage, tokenProvider, md5, 60*60*24*30)

		account, err := business.Login(c.Request.Context(), &data)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
