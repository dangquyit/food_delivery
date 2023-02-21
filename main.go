package main

import (
	"food_delivery/component/appctx"
	"food_delivery/component/uploadprovider"
	localpb "food_delivery/pubsub/local_pubsub"
	"food_delivery/subscriber"
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

	secretKey := os.Getenv("SYSTEM_SECRET")

	r := gin.Default()
	if err != nil {
		return
	}

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)

	pb := localpb.NewPubSub()
	appCtx := appctx.NewAppCtx(db, s3Provider, secretKey, pb)

	// Setup subcribe
	//subscriber.Setup(appCtx, context.Background())
	
	_ = subscriber.NewEngine(appCtx).Start()
	setupRoute(appCtx, r.Group(""))
	setupAdminRoute(appCtx, r.Group(""))
	if err := r.Run(); err != nil {
		return
	}
}
