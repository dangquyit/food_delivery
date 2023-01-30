package main

import (
	"food_delivery/component/appctx"
	"food_delivery/component/uploadprovider"
	"food_delivery/middleware"
	restaurantginrestaurant "food_delivery/module/restaurant/transport/ginrestaurant"
	"food_delivery/module/upload/transport/ginupload"
	"food_delivery/module/user/transport/ginuser"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func main() {
	dsn := "root:1234@tcp(127.0.0.1:3306)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db = db.Debug()

	s3BucketName := os.Getenv("S3BucketName")
	s3Region := os.Getenv("S3Region")
	s3APIKey := os.Getenv("S3APIKey")
	s3SecretKey := os.Getenv("S3SecretKey")
	s3Domain := os.Getenv("S3Domain")
	r := gin.Default()
	if err != nil {
		return
	}

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)

	appCtx := appctx.NewAppCtx(db, s3Provider)
	r.Use(middleware.Recover(appCtx))

	r.POST("/upload", ginupload.UploadImage(appCtx))

	r.POST("/register", ginuser.CreateUser(appCtx))

	restaurants := r.Group("/restaurants")
	restaurants.POST("", restaurantginrestaurant.CreateRestaurant(appCtx))

	restaurants.DELETE("/:id", restaurantginrestaurant.DeleteRestaurant(appCtx))

	restaurants.GET("", restaurantginrestaurant.ListRestaurant(appCtx))

	if err := r.Run(); err != nil {
		return
	}
}
