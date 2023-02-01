package restaurantginrestaurant

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	restaurantbusiness "food_delivery/module/restaurant/business"
	restaurantmodel "food_delivery/module/restaurant/model"
	restaurantrepository "food_delivery/module/restaurant/repository"
	restaurantstorage "food_delivery/module/restaurant/storage"
	restaurantlikestorage "food_delivery/module/restaurantlike/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListRestaurant(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			//c.JSON(http.StatusBadRequest, gin.H{
			//	"error": err.Error(),
			//})
			//return
			panic(common.ErrInvalidRequest(err))
		}

		var filter restaurantmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			//c.JSON(http.StatusBadRequest, gin.H{
			//	"error": err.Error(),
			//})
			//return
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSqlStore(db)
		likeStore := restaurantlikestorage.NewSQLStorage(db)
		repo := restaurantrepository.NewListRestaurantRepo(store, likeStore)
		bsn := restaurantbusiness.NewListRestaurantBusiness(repo)
		result, err := bsn.ListRestaurantBusiness(c, &filter, &paging)
		if err != nil {
			//c.JSON(http.StatusBadRequest, gin.H{
			//	"error": err.Error(),
			//})
			//return
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
