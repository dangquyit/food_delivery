package restaurantrepository

import (
	"context"
	"food_delivery/common"
	restaurantmodel "food_delivery/module/restaurant/model"
)

type ListRestaurantStore interface {
	List(ctx context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKey ...string) ([]restaurantmodel.Restaurant, error)
}

//type LikeRestaurantStore interface {
//	GetListRestaurantLikes(ctx context.Context, ids []int) (map[int]int, error)
//}

type listRestaurantRepo struct {
	store ListRestaurantStore
	//likeStore LikeRestaurantStore
}

func NewListRestaurantRepo(store ListRestaurantStore) *listRestaurantRepo {
	return &listRestaurantRepo{store: store}
}

func (repo *listRestaurantRepo) ListRestaurant(ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging) ([]restaurantmodel.Restaurant, error) {
	result, err := repo.store.List(ctx, filter, paging, "User")
	if err != nil {
		return nil, common.ErrCannotGetEntity(restaurantmodel.EntityName, err)
	}

	//ids := make([]int, len(result))
	//
	//for i := range ids {
	//	ids[i] = result[i].Id
	//}
	//
	//likeMap, err := repo.likeStore.GetListRestaurantLikes(ctx, ids)
	//
	//if err != nil {
	//	log.Println(err)
	//	return result, nil
	//}
	//
	//for i, v := range result {
	//	result[i].LikeCount = likeMap[v.Id]
	//}

	return result, nil

}
