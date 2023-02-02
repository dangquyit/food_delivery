package restaurantlikebusiness

import (
	"context"
	restaurantlikemodel "food_delivery/module/restaurantlike/model"
)

type CreateRestaurantLikeStore interface {
	CreateRestaurantLike(ctx context.Context, data *restaurantlikemodel.Like) error
}

type createRestaurantLikeBusiness struct {
	store CreateRestaurantLikeStore
}

func NewCreateBusiness(store CreateRestaurantLikeStore) *createRestaurantLikeBusiness {
	return &createRestaurantLikeBusiness{store: store}
}

func (bsn *createRestaurantLikeBusiness) CreateLikeRestaurant(ctx context.Context,
	data *restaurantlikemodel.Like) error {
	if err := bsn.store.CreateRestaurantLike(ctx, data); err != nil {
		return restaurantlikemodel.ErrCannotLikeRestaurant(err)
	}

	return nil
}
