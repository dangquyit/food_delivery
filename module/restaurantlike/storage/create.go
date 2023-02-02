package restaurantlikestorage

import (
	"context"
	"food_delivery/common"
	restaurantlikemodel "food_delivery/module/restaurantlike/model"
)

func (s *sqlStorage) CreateRestaurantLike(ctx context.Context, data *restaurantlikemodel.Like) error {
	db := s.db
	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
