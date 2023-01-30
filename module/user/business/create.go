package userbusiness

import (
	"context"
	usermodel "food_delivery/module/user/model"
)

type CreateUser interface {
	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
	FindUser(ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string) (*usermodel.User, error)
}

type createUserBusiness struct {
	createUser CreateUser
}

func NewCreateUserBusiness(createUser CreateUser) *createUserBusiness {
	return &createUserBusiness{createUser: createUser}
}

func (bsn *createUserBusiness) CreateUser(ctx context.Context, data *usermodel.UserCreate) error {
	if user, _ := bsn.createUser.FindUser(ctx, map[string]interface{}{"email": data.Email}); user != nil {
		return usermodel.ErrEmailExisted
	}

	if err := bsn.createUser.CreateUser(ctx, data); err != nil {
		return err
	}

	return nil
}
