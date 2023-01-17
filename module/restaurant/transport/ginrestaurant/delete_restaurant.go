package restaurantginrestaurant

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	restaurantbusiness "food_delivery/module/restaurant/business"
	restaurantstorage "food_delivery/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		//id, err := strconv.Atoi(c.Param("id"))
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSqlStore(db)
		bsn := restaurantbusiness.NewDeleteRestaurantBusiness(store)

		if err := bsn.DeleteRestaurant(c.Request.Context(), int(uid.GetLocalID())); err != nil {
			//c.JSON(http.StatusBadRequest, gin.H{
			//	"error": err.Error(),
			//})
			//return
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
		return
	}
}
