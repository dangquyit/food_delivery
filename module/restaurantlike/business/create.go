package restaurantlikebusiness

import (
	"context"
	"food_delivery/common"
	restaurantlikemodel "food_delivery/module/restaurantlike/model"
	"food_delivery/pubsub"
	"log"
)

type CreateRestaurantLikeStore interface {
	CreateRestaurantLike(ctx context.Context, data *restaurantlikemodel.Like) error
}

type createRestaurantLikeBusiness struct {
	store CreateRestaurantLikeStore
	ps    pubsub.Pubsub
}

func NewCreateBusiness(store CreateRestaurantLikeStore, ps pubsub.Pubsub) *createRestaurantLikeBusiness {
	return &createRestaurantLikeBusiness{store: store, ps: ps}
}

func (bsn *createRestaurantLikeBusiness) CreateLikeRestaurant(ctx context.Context,
	data *restaurantlikemodel.Like) error {
	if err := bsn.store.CreateRestaurantLike(ctx, data); err != nil {
		return restaurantlikemodel.ErrCannotLikeRestaurant(err)
	}

	// Send message
	if err := bsn.ps.Publish(ctx, common.TopicUserLikeRestaurant, pubsub.NewMessage(data)); err != nil {
		log.Println(err)
	}
	// Start routine
	//j := asyncjob.NewJob(func(ctx context.Context) error {
	//	return bsn.incStore.IncreaseLikeCount(ctx, data.RestaurantId)
	//})

	//if err := asyncjob.NewGroup(true, j).Run(ctx); err != nil {
	//	log.Println("Like restaurant err", err)
	//}
	// End routine

	return nil
}
