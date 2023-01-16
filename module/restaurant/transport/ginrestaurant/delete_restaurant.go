package restaurantginrestaurant

import (
	"fmt"
	"food_delivery/common"
	"food_delivery/component/appctx"
	restaurantbusiness "food_delivery/module/restaurant/business"
	restaurantstorage "food_delivery/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DeleteRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		store := restaurantstorage.NewSqlStore(db)
		bsn := restaurantbusiness.NewDeleteRestaurantBusiness(store)

		if err := bsn.DeleteRestaurant(c.Request.Context(), id); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
		return
	}
}
