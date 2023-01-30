package usermodel

import (
	"errors"
	"food_delivery/common"
)

const EntityName = "User"

type User struct {
	common.SQLModel
	Email     string        `json:"email" gorm:"column:email"`
	Password  string        `json:"-" gorm:"column:password"`
	Salt      string        `json:"-" gorm:"column:salt"`
	LastName  string        `json:"last_name" gorm:"column:last_name"`
	FirstName string        `json:"first_name" gorm:"column:first_name"`
	Phone     string        `json:"phone" gorm:"column:phone"`
	Role      string        `json:"role" gorm:"column:role"`
	Avatar    *common.Image `json:"avatar,omitempty" gorm:"column:avatar;type=json"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) Mask() {
	u.GenUID(common.DbTYpeUser)
}

type UserCreate struct {
	common.SQLModel
	Email     string        `json:"email" gorm:"column:email"`
	Password  string        `json:"password" gorm:"column:password"`
	LastName  string        `json:"last_name" gorm:"column:last_name"`
	FirstName string        `json:"first_name" gorm:"column:first_name"`
	Phone     string        `json:"phone" gorm:"column:phone"`
	Role      string        `json:"-" gorm:"column:role"`
	Salt      string        `json:"-" gorm:"column:salt"`
	Avatar    *common.Image `json:"avatar,omitempty" gorm:"column:avatar;type=json"`
}

func (UserCreate) TableName() string {
	return User{}.TableName()
}

func (u *UserCreate) Mask() {
	u.GenUID(common.DbTYpeUser)
}

type UserLogin struct {
	Email    string `json:"email" form:"email" gorm:"email`
	Password string `json:"password" form:"password" gorm:"password"`
}

func (UserLogin) TableName() string {
	return User{}.TableName()
}

var (
	ErrUsernameOrPassword = common.NewCustomError(
		errors.New("username or password invalid"),
		"username or password invalid",
		"ErrUsernameOrPasswordInvalid",
	)

	ErrEmailExisted = common.NewCustomError(
		errors.New("email has already existed"),
		"email has already existed",
		"ErrEmailExisted",
	)
)
