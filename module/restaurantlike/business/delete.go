package restaurantlikebusiness

import (
	"context"
	"food_delivery/component/asyncjob"
	restaurantlikemodel "food_delivery/module/restaurantlike/model"
	"log"
)

type DeclikedCountResStore interface {
	DecreaseLikeCount(ctx context.Context, id int) error
}

type DeleteRestaurantLikeStore interface {
	DeleteRestaurantLike(ctx context.Context, data *restaurantlikemodel.Like) error
}

type deleteRestaurantLikeBusiness struct {
	store    DeleteRestaurantLikeStore
	decStore DeclikedCountResStore
}

func NewDeleteBusiness(store DeleteRestaurantLikeStore, decStore DeclikedCountResStore) *deleteRestaurantLikeBusiness {
	return &deleteRestaurantLikeBusiness{store: store, decStore: decStore}
}

func (bsn *deleteRestaurantLikeBusiness) DeleteLikeRestaurant(ctx context.Context,
	data *restaurantlikemodel.Like) error {
	if err := bsn.store.DeleteRestaurantLike(ctx, data); err != nil {
		return restaurantlikemodel.ErrCannotUnlikeRestaurant(err)
	}

	j := asyncjob.NewJob(func(ctx context.Context) error {
		return bsn.decStore.DecreaseLikeCount(ctx, data.RestaurantId)
	})

	if err := asyncjob.NewGroup(true, j).Run(ctx); err != nil {
		log.Println("Unlike restaurant err", err)
	}

	return nil
}
