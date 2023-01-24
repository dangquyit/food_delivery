package uploadbusiness

import (
	"context"
	"food_delivery/common"
)

type CreateImageStorage interface {
	CreateImage(context context.Context, data *common.Image) error
}

type uploadBusiness struct {
}
