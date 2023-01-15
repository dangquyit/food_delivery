package main

import (
	restaurantginrestaurant "food_delivery/module/restaurant/transport/ginrestaurant"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:1234@tcp(127.0.0.1:3306)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	r := gin.Default()
	if err != nil {
		return
	}

	restaurants := r.Group("/restaurants")
	restaurants.POST("", restaurantginrestaurant.CreateRestaurant(db))

	if err := r.Run(); err != nil {
		return
	}
}
