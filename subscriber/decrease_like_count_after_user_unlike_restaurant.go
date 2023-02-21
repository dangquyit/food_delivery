package subscriber

import (
	"context"
	"food_delivery/component/appctx"
	restaurantstorage "food_delivery/module/restaurant/storage"
	"food_delivery/pubsub"
)

//func DecreaseLikeCountAfterUserUnLikeRestaurant(appCtx appctx.AppContext, ctx context.Context) {
//	c, _ := appCtx.GetPubSub().Subscribe(ctx, common.TopicUserUnLikeRestaurant)
//
//	store := restaurantstorage.NewSqlStore(appCtx.GetMainDBConnection())
//
//	go func() {
//		for {
//			msg := <-c
//			likeData := msg.Data().(HasRestaurantId)
//			_ = store.DecreaseLikeCount(ctx, likeData.GetRestaurantId())
//		}
//	}()
//}

func DecreaseLikeCountAfterUserUnlikeRestaurant(appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "Decrease like count after user likes restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := restaurantstorage.NewSqlStore(appCtx.GetMainDBConnection())
			likeData := message.Data().(HasRestaurantId)

			return store.DecreaseLikeCount(ctx, likeData.GetRestaurantId())
		},
	}
}
