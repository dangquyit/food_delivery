package common

const (
	DbTypeRestaurant = 1
	DbTYpeUser       = 2
)

const (
	TokenPayloadInJWTRequest = "tokenPayload"
)

type Requester interface {
	GetUserId() int
	GetRole() string
}
