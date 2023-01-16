package restaurantstorage

import (
	"context"
	"food_delivery/common"
	restaurantmodel "food_delivery/module/restaurant/model"
)

func (s *sqlStore) List(ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKey ...string) ([]restaurantmodel.Restaurant, error) {

	var result []restaurantmodel.Restaurant

	db := s.db.Table(restaurantmodel.Restaurant{}.TableName()).Where("status = 1")

	// Check filter
	if filter != nil {
		if filter.Addr != "" {
			db = db.Where("addr = ?", filter.Addr)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	// Check paging
	paging.Fulfill()
	db = db.Limit(paging.Limit).Offset((paging.Page - 1) * paging.Limit)

	// Query data
	if err := db.Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
