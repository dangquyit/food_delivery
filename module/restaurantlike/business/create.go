package restaurantlikebusiness

import (
	"context"
	restaurantlikemodel "food_delivery/module/restaurantlike/model"
	"log"
)

type CreateRestaurantLikeStore interface {
	CreateRestaurantLike(ctx context.Context, data *restaurantlikemodel.Like) error
}

type InclikedCountResStore interface {
	IncreaseLikeCount(ctx context.Context, id int) error
}

type createRestaurantLikeBusiness struct {
	store    CreateRestaurantLikeStore
	incStore InclikedCountResStore
}

func NewCreateBusiness(store CreateRestaurantLikeStore, incStore InclikedCountResStore) *createRestaurantLikeBusiness {
	return &createRestaurantLikeBusiness{store: store, incStore: incStore}
}

func (bsn *createRestaurantLikeBusiness) CreateLikeRestaurant(ctx context.Context,
	data *restaurantlikemodel.Like) error {
	if err := bsn.store.CreateRestaurantLike(ctx, data); err != nil {
		return restaurantlikemodel.ErrCannotLikeRestaurant(err)
	}

	if err := bsn.incStore.IncreaseLikeCount(ctx, data.RestaurantId); err != nil {
		log.Println(err)
	}

	return nil
}
