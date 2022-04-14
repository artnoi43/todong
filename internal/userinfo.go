package internal

import "time"

// UserInfo is a type returned by utils.ExtractAndDecodeJwt
// Currently, this application only encodes issuer's UUID and token expiration in the JWT token
type UserInfo struct {
	UserUuid   string    `json:"userUuid"`
	Expiration time.Time `json:"expiration"`
}
