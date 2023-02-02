package restaurantlikestorage

import (
	"food_delivery/common"
	restaurantlikemodel "food_delivery/module/restaurantlike/model"
	"golang.org/x/net/context"
	"log"
)

func (s *sqlStorage) ListUserLikeRestaurant(ctx context.Context,
	conditions map[string]interface{},
	filter *restaurantlikemodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]common.SimpleUser, error) {
	var result []restaurantlikemodel.Like
	db := s.db

	db = db.Table(restaurantlikemodel.Like{}.TableName()).Where(conditions)

	if filter != nil {
		if filter.RestaurantId > 0 {
			db.Where("restaurant_id = ?", filter.RestaurantId)
		}
	}

	db.Preload("User")

	paging.Fulfill()

	if err := db.Limit(paging.Limit).Order("created_at desc").Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	users := make([]common.SimpleUser, len(result))

	log.Println(result)

	for i, v := range result {
		result[i].User.CreatedAt = v.CreatedAt
		result[i].User.UpdatedAt = nil
		users[i] = *result[i].User
	}

	return users, nil
}
