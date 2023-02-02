package main

import (
	"food_delivery/component/appctx"
	"food_delivery/middleware"
	restaurantginrestaurant "food_delivery/module/restaurant/transport/ginrestaurant"
	restaurantlikegin "food_delivery/module/restaurantlike/transport/gin"
	"food_delivery/module/upload/transport/ginupload"
	"food_delivery/module/user/transport/ginuser"
	"github.com/gin-gonic/gin"
)

func setupRoute(appCtx appctx.AppContext, r *gin.RouterGroup) {
	r.Use(middleware.Recover(appCtx))

	r.POST("/authenticate", ginuser.Login(appCtx))
	r.POST("/upload", ginupload.UploadImage(appCtx))
	r.GET("/profile", middleware.AuthenticateJWT(appCtx), ginuser.FindUser(appCtx))
	r.POST("/register", ginuser.CreateUser(appCtx))

	restaurants := r.Group("/restaurants", middleware.AuthenticateJWT(appCtx))
	restaurants.POST("", restaurantginrestaurant.CreateRestaurant(appCtx))

	restaurants.DELETE("/:id", restaurantginrestaurant.DeleteRestaurant(appCtx))

	restaurants.GET("", restaurantginrestaurant.ListRestaurant(appCtx))

	restaurants.POST("/:id/like", restaurantlikegin.UserLikeRestaurant(appCtx))
}
