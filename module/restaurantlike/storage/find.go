package restaurantlikestorage

import (
	"context"
	"food_delivery/common"
	restaurantlikemodel "food_delivery/module/restaurantlike/model"
)

func (s *sqlStorage) Find(ctx context.Context, data *restaurantlikemodel.Like) (*restaurantlikemodel.Like, error) {
	if err := s.db.Where("restaurant_id = ? and user_id = ?",
		data.RestaurantId, data.UserId).First(&data).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return data, nil
}
