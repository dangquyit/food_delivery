package restaurantbusiness

import (
	"context"
	"errors"
	"fmt"
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
		return err
	}

	if oldData.Status == 0 {
		fmt.Println("Error")
		return errors.New("data has been deleted")
	}

	if err := bsn.store.Delete(context, id); err != nil {
		return err
	}
	return nil
}
