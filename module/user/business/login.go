package userbusiness

import (
	"context"
	"errors"
	"food_delivery/common"
	"food_delivery/component/tokenprovider"
	usermodel "food_delivery/module/user/model"
)

type LoginStorage interface {
	FindUser(ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string) (*usermodel.User, error)
}

type loginBusiness struct {
	storeUser     LoginStorage
	tokenProvider tokenprovider.Provider
	hasher        Hasher
	expiri        int
}

func NewLoginBusiness(
	storeUser LoginStorage,
	tokenProvider tokenprovider.Provider,
	hasher Hasher, expiri int) *loginBusiness {
	return &loginBusiness{
		storeUser:     storeUser,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		expiri:        expiri,
	}
}

func (bsn *loginBusiness) Login(ctx context.Context, data *usermodel.UserLogin) (*tokenprovider.Token, error) {
	user, err := bsn.storeUser.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if err != nil {
		return nil, usermodel.ErrUsernameOrPassword
	}

	passHashed := bsn.hasher.Hash(data.Password + user.Salt)

	if user.Password != passHashed {
		return nil, usermodel.ErrUsernameOrPassword
	}

	payload := tokenprovider.TokenPayload{
		UserId: user.Id,
		Role:   user.Role,
	}

	accessToken, err := bsn.tokenProvider.Generate(payload, bsn.expiri)

	if err != nil {
		return nil, common.ErrInternal(err)
	}

	return accessToken, nil
}

var (
	ErrEmailOrPasswordInvalid = common.NewCustomError(
		errors.New("email or password invalid"),
		"email or password invalid",
		"ErrEmailOrPasswordInvalid")

	ErrEmailExisted = common.NewCustomError(
		errors.New("email has already existed"),
		"email has already existed",
		"ErrEmailExisted")
)
