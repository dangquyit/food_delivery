package userstorage

import (
	"food_delivery/common"
	usermodel "food_delivery/module/user/model"
	"golang.org/x/net/context"
)

func (s *sqlStorage) FindUser(ctx context.Context,
	conditions map[string]interface{},
	moreInfo ...string) (*usermodel.User, error) {
	db := s.db

	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}

	var user usermodel.User

	if err := db.Table(usermodel.User{}.TableName()).Where(conditions).First(&user).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return &user, nil
}
