package restaurantlikestorage

import (
	"context"
	restaurantlikemodel "food_delivery/module/restaurantlike/model"
)

func (s *sqlStorage) GetListRestaurantLikes(ctx context.Context, ids []int) (map[int]int, error) {
	result := make(map[int]int)

	type sqlData struct {
		RestaurantId int `gorm:"column:restaurant_id"`
		LikeCount    int `gorm:"column:count"`
	}

	var listLike []sqlData

	if err := s.db.Table(restaurantlikemodel.Like{}.TableName()).
		Select("restaurant_id, count(restaurant_id) as count").
		Where("restaurant_id in (?)", ids).
		Group("restaurant_id").Find(&listLike).Error; err != nil {
		return nil, err
	}

	for _, v := range listLike {
		result[v.RestaurantId] = v.LikeCount
	}

	return result, nil
}
