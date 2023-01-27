package uploadbusiness

import (
	"bytes"
	"context"
	"food_delivery/common"
	"food_delivery/component/uploadprovider"
	uploadmodel "food_delivery/module/upload/model"
	"image"
	"io"
	"log"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type CreateImageStorage interface {
	CreateImage(context context.Context, data *common.Image) error
}

type uploadBusiness struct {
	provider  uploadprovider.UploadProvider
	imagStore CreateImageStorage
}

func NewUploadBusiness(provider uploadprovider.UploadProvider, imgStore CreateImageStorage) *uploadBusiness {
	return &uploadBusiness{
		provider:  provider,
		imagStore: imgStore,
	}
}

func (bsn *uploadBusiness) Upload(ctx context.Context, data []byte, folder, fileName string) (*common.Image, error) {
	fileBytes := bytes.NewBuffer(data)

	w, h, err := getImageDimension(fileBytes)

	if err != nil {
		return nil, uploadmodel.ErrFileIsNotImage(err)
	}

	if strings.TrimSpace(folder) == "" {
		folder = "img"
	}

	fileExt := filepath.Ext(fileName) // img, png, .jpg
	fileName = strconv.Itoa(time.Now().Nanosecond()) + fileName

	img, err := bsn.provider.SaveFileUploaded(ctx, data, folder+"/"+fileName)

	if err != nil {
		return nil, uploadmodel.ErrCannotSaveFile(err)
	}

	img.Width = w
	img.Height = h
	img.Extension = fileExt
	return img, nil
}

func getImageDimension(reader io.Reader) (int, int, error) {
	img, _, err := image.DecodeConfig(reader)

	if err != nil {
		log.Println("err: ", err)
		return 0, 0, err
	}
	return img.Width, img.Height, nil
}
