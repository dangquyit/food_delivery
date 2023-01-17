package restaurantginrestaurant

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	restaurantbusiness "food_delivery/module/restaurant/business"
	restaurantmodel "food_delivery/module/restaurant/model"
	restaurantstorage "food_delivery/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var data restaurantmodel.RestaurantCreate

		if err := c.ShouldBind(&data); err != nil {
			//c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			//return
			panic(common.ErrInvalidRequest(err))

		}

		store := restaurantstorage.NewSqlStore(db)
		bsn := restaurantbusiness.NewCreateRestaurantBusiness(store)

		if err := bsn.CreateRestaurant(c.Request.Context(), &data); err != nil {
			//c.JSON(http.StatusBadRequest, err)
			//
			panic(err)
		}

		data.Mask(true)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeID))
	}
}
