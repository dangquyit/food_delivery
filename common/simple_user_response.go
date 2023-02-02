package common

type SimpleUser struct {
	SQLModel
	LastName  string `json:"last_name" gorm:"column:last_name"`
	FirstName string `json:"first_name" gorm:"column:first_name"`
	Status    int    `json:"status" gorm:"column:status;default:1"`
	Role      string `json:"role" gorm:"column:role"`
	Avatar    *Image `json:"avatar,omitempty" gorm:"column:avatar;type=json"`
}

func (SimpleUser) TableName() string {
	return "users"
}

func (u *SimpleUser) Mask(isAdmin bool) {
	u.GenUID(DbTYpeUser)
}
