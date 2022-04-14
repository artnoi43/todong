package middleware

import (
	ginjwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gofiber/fiber/v2"
	fiberjwt "github.com/gofiber/jwt/v3"
)

func AuthenticateGin(conf *Config) (*ginjwt.GinJWTMiddleware, error) {
	return ginjwt.New(&ginjwt.GinJWTMiddleware{
		Key: []byte(conf.SecretKey),
	})
}

func AuthenticateFiber(conf *Config) func(*fiber.Ctx) error {
	return fiberjwt.New(fiberjwt.Config{
		SigningKey: []byte(conf.SecretKey),
	})
}
