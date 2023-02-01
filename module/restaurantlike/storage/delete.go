package restaurantlikestorage

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	restaurantlikemodel "food_delivery/module/restaurantlike/model"
)

func (s *sqlStorage) DeleteRestaurantLike(ctx appctx.AppContext, userId, restaurantId int) error {
	db := s.db
	if err := db.Table(restaurantlikemodel.Like{}.TableName()).
		Where("user_id = ? and restaurant_id = ?", userId, restaurantId).
		Delete(nil).
		Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
