package restaurantbusiness

import (
	"context"
	"food_delivery/common"
	restaurantmodel "food_delivery/module/restaurant/model"
)

type CreateRestaurantStore interface {
	Create(context context.Context, data *restaurantmodel.RestaurantCreate) error
}

type createRestaurantBusiness struct {
	store CreateRestaurantStore
}

func NewCreateRestaurantBusiness(store CreateRestaurantStore) *createRestaurantBusiness {
	return &createRestaurantBusiness{store: store}
}
func (bsn *createRestaurantBusiness) CreateRestaurant(context context.Context, data *restaurantmodel.RestaurantCreate) error {
	if err := data.Validate(); err != nil {
		return common.ErrInvalidRequest(err)
	}

	if err := bsn.store.Create(context, data); err != nil {
		return common.ErrCannotCreateEntity(restaurantmodel.EntityName, err)
	}

	return nil
}
