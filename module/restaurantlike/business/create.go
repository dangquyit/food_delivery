package restaurantlikebusiness

import (
	"context"
	"food_delivery/component/asyncjob"
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

	// Start routine
	j := asyncjob.NewJob(func(ctx context.Context) error {
		return bsn.incStore.IncreaseLikeCount(ctx, data.RestaurantId)
	})

	if err := asyncjob.NewGroup(true, j).Run(ctx); err != nil {
		log.Println("Like restaurant err", err)
	}
	// End routine

	return nil
}
