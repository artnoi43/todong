package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

func NewJwtToken(iss string, key []byte) (token string, exp time.Time, err error) {
	// TODO: investigate if Local() is actually needed
	exp = time.Now().Add(24 * time.Hour).Local()
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    iss,
		ExpiresAt: exp.Unix(),
	})
	// Generate JWT token from claims
	token, err = claims.SignedString(key)
	if err != nil {
		return token, exp, errors.Wrapf(err, "failed to validate with key %s", key)
	}
	return token, exp, nil
}
