package utils

import (
	"fmt"
	"time"

	ginjwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/artnoi43/todong/internal"
)

func ExtractAndDecodeJwt(c *gin.Context) (*internal.UserInfo, error) {
	claims, exists := c.Get("JWT_PAYLOAD")
	if !exists {
		return nil, errors.New("claims does not exists")
	}
	var userUuid string
	var expUnixFloat float64
	if mc, ok := claims.(ginjwt.MapClaims); ok {
		iss := mc["iss"]
		exp := mc["exp"]
		if userUuid, ok = iss.(string); !ok {
			return nil, fmt.Errorf("jwt issuer extract: interface{} converstion to string failed for: %v %T", iss, iss)
		}
		if expUnixFloat, ok = exp.(float64); !ok {
			return nil, fmt.Errorf("jwt expiration extract: interface{} converstion to float64 failed for: %v %T", exp, exp)
		}
	}
	expiration := time.Unix(int64(expUnixFloat), 0)
	return &internal.UserInfo{
		UserUuid:   userUuid,
		Expiration: expiration,
	}, nil
}
