package ginupload

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	uploadbusiness "food_delivery/module/upload/business"
	"github.com/gin-gonic/gin"
	"log"
)

func UploadImage(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		folder := c.DefaultPostForm("folder", "img")

		file, err := fileHeader.Open()
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		defer file.Close()

		dataBytes := make([]byte, fileHeader.Size)
		if _, err := file.Read(dataBytes); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		bsn := uploadbusiness.NewUploadBusiness(appCtx.UploadProvider(), nil)
		log.Println(folder, fileHeader.Filename)
		img, err := bsn.Upload(c.Request.Context(), dataBytes, folder, fileHeader.Filename)

		if err != nil {
			panic(err)
		}

		c.JSON(200, common.SimpleSuccessResponse(img))
	}
}
