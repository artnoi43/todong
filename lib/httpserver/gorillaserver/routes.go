package gorillaserver

import (
	"net/http"

	"github.com/artnoi43/todong/enums"
	"github.com/artnoi43/todong/lib/handler"
	"github.com/artnoi43/todong/lib/middleware"
	"github.com/artnoi43/todong/lib/utils"
)

func (s *gorillaServer) SetUpRoutes(conf *middleware.Config, h handler.Adaptor) {

	authenticator := utils.NewAuthenticator([]byte(conf.SecretKey))
	usersApi := s.router.PathPrefix("/users").Subrouter()
	todosApi := s.router.PathPrefix("/todos").Subrouter()
	todosApi.Use(authenticator.AuthMiddleware)

	usersApi.
		Handle("/register", h.Gorilla(enums.HandlerRegister)).
		Methods(http.MethodPost)
	usersApi.
		Handle("/login", h.Gorilla(enums.HandlerLogin)).
		Methods(http.MethodPost)
}
