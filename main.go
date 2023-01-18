package main

import (
	"food_delivery/component/appctx"
	"food_delivery/middleware"
	restaurantginrestaurant "food_delivery/module/restaurant/transport/ginrestaurant"
	"food_delivery/module/upload/transport/ginupload"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:1234@tcp(127.0.0.1:3306)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db = db.Debug()

	r := gin.Default()
	if err != nil {
		return
	}
	appCtx := appctx.NewAppCtx(db)
	r.Use(middleware.Recover(appCtx))

	r.Static("/static", "./static")
	r.POST("/upload", ginupload.UploadImage(appCtx))

	restaurants := r.Group("/restaurants")
	restaurants.POST("", restaurantginrestaurant.CreateRestaurant(appCtx))

	restaurants.DELETE("/:id", restaurantginrestaurant.DeleteRestaurant(appCtx))

	restaurants.GET("", restaurantginrestaurant.ListRestaurant(appCtx))

	if err := r.Run(); err != nil {
		return
	}
}
