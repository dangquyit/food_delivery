package restaurantmodel

type Filter struct {
	Addr string `json:"addr,omitempty" form:"addr"`
}
