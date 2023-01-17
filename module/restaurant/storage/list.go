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
	if v := paging.FakeCursor; v != "" {
		uid, err := common.FromBase58(v)
		if err != nil {
			return nil, common.ErrDB(err)
		}
		db = db.Where("id < ?", uid.GetLocalID())
	} else {
		offSet := (paging.Page - 1) * paging.Limit
		db = db.Offset(offSet)
	}

	// Query data
	if err := db.Limit(paging.Limit).Order("id desc").Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if len(result) > 0 {
		last := result[len(result)-1]
		last.Mask(true)
		paging.NextCursor = last.FakeID.String()
	}

	return result, nil
}
