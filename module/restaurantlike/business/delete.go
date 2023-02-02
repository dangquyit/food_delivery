package restaurantlikebusiness

import (
	"context"
	restaurantlikemodel "food_delivery/module/restaurantlike/model"
)

type DeleteRestaurantLikeStore interface {
	DeleteRestaurantLike(ctx context.Context, data *restaurantlikemodel.Like) error
}

type deleteRestaurantLikeBusiness struct {
	store DeleteRestaurantLikeStore
}

func NewDeleteBusiness(store DeleteRestaurantLikeStore) *deleteRestaurantLikeBusiness {
	return &deleteRestaurantLikeBusiness{store: store}
}

func (bsn *deleteRestaurantLikeBusiness) DeleteLikeRestaurant(ctx context.Context,
	data *restaurantlikemodel.Like) error {
	if err := bsn.store.DeleteRestaurantLike(ctx, data); err != nil {
		return restaurantlikemodel.ErrCannotUnlikeRestaurant(err)
	}
	return nil
}
