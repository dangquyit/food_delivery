package restaurantbusiness

import (
	"context"
	"food_delivery/common"
	restaurantmodel "food_delivery/module/restaurant/model"
)

type DeleteRestaurantStore interface {
	Delete(context context.Context, id int) error
	Find(context context.Context,
		condition map[string]interface{},
		moreKeys ...string) (*restaurantmodel.Restaurant, error)
}

type deleteRestaurantBusiness struct {
	store DeleteRestaurantStore
}

func NewDeleteRestaurantBusiness(store DeleteRestaurantStore) *deleteRestaurantBusiness {
	return &deleteRestaurantBusiness{store: store}
}

func (bsn *deleteRestaurantBusiness) DeleteRestaurant(context context.Context, id int) error {

	oldData, err := bsn.store.Find(context, map[string]interface{}{"id": id})

	if err != nil {
		return common.ErrCannotGetEntity(restaurantmodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrCannotDeleteEntity(restaurantmodel.EntityName, err)
	}

	if err := bsn.store.Delete(context, id); err != nil {
		return common.ErrCannotDeleteEntity(restaurantmodel.EntityName, err)
	}
	return nil
}
