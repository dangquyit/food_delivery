package ginupload

import (
	"fmt"
	"food_delivery/common"
	"food_delivery/component/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func UploadImage(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := c.SaveUploadedFile(fileHeader, fmt.Sprintf("static/%s", fileHeader.Filename)); err != nil {
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

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(common.Image{
			Id:        0,
			Url:       "http://localhost:8080/static/" + fileHeader.Filename,
			Width:     0,
			Height:    0,
			CloudName: "local",
			Extension: fileHeader.Filename[strings.Index(fileHeader.Filename, "."):len(fileHeader.Filename)],
		}))
	}
}
