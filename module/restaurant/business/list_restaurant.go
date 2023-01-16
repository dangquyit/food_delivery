package restaurantbusiness

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

type listRestaurantBusiness struct {
	store ListRestaurantStore
}

func NewListRestaurantBusiness(store ListRestaurantStore) *listRestaurantBusiness {
	return &listRestaurantBusiness{store: store}
}

func (bsn *listRestaurantBusiness) ListRestaurantBusiness(ctx context.Context, filter *restaurantmodel.Filter, paging *common.Paging) ([]restaurantmodel.Restaurant, error) {
	var result, err = bsn.store.List(ctx, filter, paging)
	if err != nil {
		return nil, err
	}

	return result, nil

}
