package userbusiness

import (
	"context"
	"food_delivery/common"
	usermodel "food_delivery/module/user/model"
)

type CreateUser interface {
	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
	FindUser(ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string) (*usermodel.User, error)
}

type Hasher interface {
	Hash(data string) string
}

type createUserBusiness struct {
	createUser CreateUser
	hasher     Hasher
}

func NewCreateUserBusiness(createUser CreateUser, hasher Hasher) *createUserBusiness {
	return &createUserBusiness{
		createUser: createUser,
		hasher:     hasher}
}

func (bsn *createUserBusiness) CreateUser(ctx context.Context, data *usermodel.UserCreate) error {
	if user, _ := bsn.createUser.FindUser(ctx, map[string]interface{}{"email": data.Email}); user != nil {
		return usermodel.ErrEmailExisted
	}

	salt := common.GenSalt(50)
	data.Password = bsn.hasher.Hash(data.Password + salt)
	data.Salt = salt
	data.Role = "user"
	if err := bsn.createUser.CreateUser(ctx, data); err != nil {
		return err
	}

	return nil
}
