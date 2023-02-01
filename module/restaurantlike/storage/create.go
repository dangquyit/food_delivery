package restaurantlikestorage

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	restaurantlikemodel "food_delivery/module/restaurantlike/model"
)

func (s *sqlStorage) CreateRestaurantLike(ctx appctx.AppContext, data *restaurantlikemodel.Like) error {
	db := s.db
	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
