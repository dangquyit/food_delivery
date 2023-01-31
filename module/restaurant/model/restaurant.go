package restaurantmodel

import (
	"errors"
	"food_delivery/common"
	usermodel "food_delivery/module/user/model"
	"strings"
)

type Restaurant struct {
	common.SQLModel `json:",inline"`
	Name            string          `json:"name" gorm:"column:name;"`
	Addr            string          `json:"addr" gorm:"column:addr;"`
	Logo            *common.Image   `json:"logo" gorm:"column:logo"`
	Cover           *common.Images  `json:"cover" gorm:"column:cover"`
	UserId          int             `json:"-" gorm:"column:user_id"`
	User            *usermodel.User `json:"user" gorm:"preload:false;"`
}

const EntityName = "Restaurant"

func (Restaurant) TableName() string {
	return "restaurants"
}

func (r *Restaurant) Mask(isAdminOrOwner bool) {
	r.GenUID(common.DbTypeRestaurant)
	if u := r.User; u != nil {
		u.Mask(false)
	}
}

type RestaurantUpdate struct {
	Name  *string        `json:"name" gorm:"column:name;"`
	Addr  *string        `json:"addr" gorm:"column:addr"`
	Logo  *common.Image  `json:"logo" gorm:"column:logo"`
	Cover *common.Images `json:"cover" gorm:"column:cover"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

type RestaurantCreate struct {
	common.SQLModel
	Name   string         `json:"name" gorm:"column:name;"`
	Addr   string         `json:"addr" gorm:"column:addr"`
	UserId int            `json:"-" gorm:"column:user_id"`
	Logo   *common.Image  `json:"logo" gorm:"column:logo"`
	Cover  *common.Images `json:"cover" gorm:"column:cover"`
}

func (data *RestaurantCreate) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DbTypeRestaurant)
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}

func (data *RestaurantCreate) Validate() error {
	data.Name = strings.TrimSpace(data.Name)
	if data.Name == "" {
		return ErrNameIsEmpty
	}

	return nil
}

var (
	ErrNameIsEmpty = errors.New("name can not be empty")
)
