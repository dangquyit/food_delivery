package restaurantlikestorage

import (
	"context"
	"food_delivery/common"
	restaurantlikemodel "food_delivery/module/restaurantlike/model"
)

func (s *sqlStorage) DeleteRestaurantLike(ctx context.Context, data *restaurantlikemodel.Like) error {
	db := s.db
	if err := db.Table(restaurantlikemodel.Like{}.TableName()).
		Where("user_id = ? and restaurant_id = ?", data.UserId, data.RestaurantId).
		Delete(nil).
		Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
