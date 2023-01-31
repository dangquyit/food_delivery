package jwt

import (
	"food_delivery/component/tokenprovider"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type jwtProvider struct {
	secret string
}

func NewTokenJWTProvider(secret string) *jwtProvider {
	return &jwtProvider{secret: secret}
}

type myClaims struct {
	Playload tokenprovider.TokenPayload `json:"playload"`
	jwt.StandardClaims
}

func (j *jwtProvider) Generate(data tokenprovider.TokenPayload, expiry int) (*tokenprovider.Token, error) {
	t := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		myClaims{
			data,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Local().Add(time.Second * time.Duration(expiry)).Unix(),
				IssuedAt:  time.Now().Local().Unix(),
			},
		})
	myToken, err := t.SignedString([]byte(j.secret))
	if err != nil {
		return nil, err
	}

	return &tokenprovider.Token{
		Token:   myToken,
		Expiry:  expiry,
		Created: time.Now(),
	}, nil
}
