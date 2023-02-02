package restaurantlikegin

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	restaurantlikebusiness "food_delivery/module/restaurantlike/business"
	restaurantlikemodel "food_delivery/module/restaurantlike/model"
	restaurantlikestorage "food_delivery/module/restaurantlike/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GET /restaurant/:id/like
func ListUserLikeRestaurantHandler(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		filter := restaurantlikemodel.Filter{
			RestaurantId: int(uid.GetLocalID()),
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		paging.Fulfill()

		store := restaurantlikestorage.NewSQLStorage(appCtx.GetMainDBConnection())
		bsn := restaurantlikebusiness.NewListUserLikeRestaurantBusiness(store)

		result, err := bsn.ListUserLikeRestaurantBusiness(c.Request.Context(),
			nil, &filter, &paging)

		if err != nil {
			panic(err)
		}

		for i, _ := range result {
			result[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))

	}
}
