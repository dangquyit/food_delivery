package userbusiness

import (
	"context"
	"errors"
	"food_delivery/common"
	usermodel "food_delivery/module/user/model"
)

type FindUserStorage interface {
	FindUser(ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string) (*usermodel.User, error)
}

type findUserBusiness struct {
	userStorage FindUserStorage
}

func NewFindUserBusiness(userStorage FindUserStorage) *findUserBusiness {
	return &findUserBusiness{
		userStorage: userStorage,
	}
}

func (bsn *findUserBusiness) Find(ctx context.Context, data *usermodel.User) (*usermodel.User, error) {
	user, err := bsn.userStorage.FindUser(ctx, map[string]interface{}{"id": data.Id, "status": 1})
	if err != nil {
		return nil, common.NewCustomError(
			errors.New("cannot found id user"),
			"cannot found id user",
			"ErrNotFoundIdUser")
	}

	return user, nil
}
