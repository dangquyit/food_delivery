package common

const (
	DbTypeRestaurant = 1
	DbTYpeUser       = 2
)

const (
	TokenPayloadInJWTRequest = "tokenPayload"
)

const (
	TopicUserLikeRestaurant   = "TopicUserLikeRestaurant"
	TopicUserUnLikeRestaurant = "TopicUserUnLikeRestaurant"
)

type Requester interface {
	GetUserId() int
	GetRole() string
}
