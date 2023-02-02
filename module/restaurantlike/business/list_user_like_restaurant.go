package restaurantlikebusiness

import (
	"context"
	"food_delivery/common"
	restaurantlikemodel "food_delivery/module/restaurantlike/model"
)

type ListUserLikeRestaurantStore interface {
	ListUserLikeRestaurant(ctx context.Context,
		conditions map[string]interface{},
		filter *restaurantlikemodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]common.SimpleUser, error)
}

type listUserLikeRestaurantBusiness struct {
	store ListUserLikeRestaurantStore
}

func NewListUserLikeRestaurantBusiness(store ListUserLikeRestaurantStore) *listUserLikeRestaurantBusiness {
	return &listUserLikeRestaurantBusiness{store: store}
}

func (bsn *listUserLikeRestaurantBusiness) ListUserLikeRestaurantBusiness(ctx context.Context,
	conditions map[string]interface{},
	filter *restaurantlikemodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]common.SimpleUser, error) {
	result, err := bsn.store.ListUserLikeRestaurant(ctx, conditions, filter, paging)
	if err != nil {
		return nil, err
	}

	return result, err
}
