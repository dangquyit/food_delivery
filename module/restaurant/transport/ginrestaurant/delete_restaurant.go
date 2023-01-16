package restaurantginrestaurant

import (
	"fmt"
	restaurantbusiness "food_delivery/module/restaurant/business"
	restaurantstorage "food_delivery/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func DeleteRestaurant(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
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

		c.JSON(http.StatusOK, gin.H{
			"data": 1,
		})
		return
	}
}
