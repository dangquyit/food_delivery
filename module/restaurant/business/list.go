package restaurantbusiness

import (
	"context"
	"food_delivery/common"
	restaurantmodel "food_delivery/module/restaurant/model"
)

type ListRestaurantRepo interface {
	ListRestaurant(ctx context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging) ([]restaurantmodel.Restaurant, error)
}

type listRestaurantBusiness struct {
	repo ListRestaurantRepo
}

func NewListRestaurantBusiness(repo ListRestaurantRepo) *listRestaurantBusiness {
	return &listRestaurantBusiness{repo: repo}
}

func (bsn *listRestaurantBusiness) ListRestaurantBusiness(ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging) ([]restaurantmodel.Restaurant, error) {
	result, err := bsn.repo.ListRestaurant(ctx, filter, paging)
	if err != nil {
		return nil, common.ErrCannotGetEntity(restaurantmodel.EntityName, err)
	}

	return result, nil

}
