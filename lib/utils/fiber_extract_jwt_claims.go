package utils

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"

	"github.com/artnoi43/todong/internal"
)

func ExtractAndDecodeJwtFiber(c *fiber.Ctx) (*internal.UserInfo, error) {
	user, ok := c.Locals("user").(*jwt.Token)
	if !ok {
		return nil, errors.New("interface{} conversion to github.com/golang-jwt/jwt/v4.Token failed")
	}
	claims, ok := user.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("interface{} conversion to github.com/golang-jwt/jwt/v4.MapClaims failed")
	}
	iss := claims["iss"]
	exp := claims["exp"]
	userUuid, ok := iss.(string)
	if !ok {
		return nil, fmt.Errorf("ExtractAndDecodeJwtFiber: interface{} conversion failed for: %v", iss)
	}
	expUnixFloat, ok := exp.(float64)
	if !ok {
		return nil, fmt.Errorf("jwt expiration extract: interface{} converstion to float64 failed for: %v %T", exp, exp)
	}
	return &internal.UserInfo{
		UserUuid:   userUuid,
		Expiration: time.Unix(int64(expUnixFloat), 0),
	}, nil
}
