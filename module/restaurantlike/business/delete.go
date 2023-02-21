package restaurantlikebusiness

import (
	"context"
	"food_delivery/common"
	restaurantlikemodel "food_delivery/module/restaurantlike/model"
	"food_delivery/pubsub"
	"log"
)

type DeleteRestaurantLikeStore interface {
	DeleteRestaurantLike(ctx context.Context, data *restaurantlikemodel.Like) error
}

type deleteRestaurantLikeBusiness struct {
	store DeleteRestaurantLikeStore
	ps    pubsub.Pubsub
}

func NewDeleteBusiness(store DeleteRestaurantLikeStore, ps pubsub.Pubsub) *deleteRestaurantLikeBusiness {
	return &deleteRestaurantLikeBusiness{store: store, ps: ps}
}

func (bsn *deleteRestaurantLikeBusiness) DeleteLikeRestaurant(ctx context.Context,
	data *restaurantlikemodel.Like) error {
	if err := bsn.store.DeleteRestaurantLike(ctx, data); err != nil {
		return restaurantlikemodel.ErrCannotUnlikeRestaurant(err)
	}

	if err := bsn.ps.Publish(ctx, common.TopicUserUnLikeRestaurant, pubsub.NewMessage(data)); err != nil {
		log.Println(err)
	}
	//j := asyncjob.NewJob(func(ctx context.Context) error {
	//	return bsn.decStore.DecreaseLikeCount(ctx, data.RestaurantId)
	//})
	//
	//if err := asyncjob.NewGroup(true, j).Run(ctx); err != nil {
	//	log.Println("Unlike restaurant err", err)
	//}

	return nil
}
