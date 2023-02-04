package restaurantlikestorage

import (
	"food_delivery/common"
	restaurantlikemodel "food_delivery/module/restaurantlike/model"
	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/net/context"
	"time"
)

const timeFormat = "2006-01-02 15:04:05.999999999"

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

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	paging.Fulfill()
	if f := paging.FakeCursor; f != "" {
		timeCreated, err := time.Parse(timeFormat, string(base58.Decode(f)))
		if err != nil {
			return nil, common.ErrDB(err)
		}

		db = db.Where("created_at < ?", timeCreated.Format(timeFormat))
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	db.Preload("User")

	if err := db.Limit(paging.Limit).Order("created_at desc").Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	users := make([]common.SimpleUser, len(result))
	
	for i, v := range result {
		result[i].User.CreatedAt = v.CreatedAt
		result[i].User.UpdatedAt = nil
		users[i] = *result[i].User
		if len(result)-1 == i {
			paging.NextCursor = base58.Encode([]byte(v.CreatedAt.Format(timeFormat)))
		}
	}

	return users, nil
}
