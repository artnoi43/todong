package utils

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
)

type Authenticator struct {
	JwtKey []byte
}

func NewAuthenticator(jwtKey []byte) *Authenticator {
	return &Authenticator{
		JwtKey: jwtKey,
	}
}

func (a *Authenticator) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("Authorization")
		if len(tokenStr) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Missing Authorization Header"))
			return
		}
		tokenStr = strings.Replace(tokenStr, "Bearer ", "", 1)
		claims, err := verifyJwt(tokenStr, a.JwtKey)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("error verifying JWT token: " + err.Error()))
			return
		}
		iss := claims.(jwt.MapClaims)["iss"].(string)
		exp := claims.(jwt.MapClaims)["exp"].(float64)

		r.Header.Set("iss", iss)
		r.Header.Set("exp", fmt.Sprintf("%f", exp))

		next.ServeHTTP(w, r)
	})
}

func verifyJwt(tokenStr string, key []byte) (jwt.Claims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse JWT token %s", tokenStr)
	}
	return token.Claims, nil
}
