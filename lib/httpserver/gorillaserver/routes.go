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
	protectedUsersApi := usersApi.NewRoute().Subrouter()
	protectedUsersApi.Use(authenticator.AuthMiddleware)
	todosApi.Use(authenticator.AuthMiddleware)

	// /users
	usersApi.
		Handle("/register", h.Gorilla(enums.HandlerRegister)).
		Methods(http.MethodPost)
	usersApi.
		Handle("/login", h.Gorilla(enums.HandlerLogin)).
		Methods(http.MethodPost)
	protectedUsersApi.
		Handle("/", h.Gorilla(enums.HandlerDeleteUser)).
		Methods(http.MethodDelete)
	protectedUsersApi.
		Handle("/new-password", h.Gorilla(enums.HandlerNewPassword)).
		Methods(http.MethodPost)

	// /todos
	todosApi.
		Handle("/create", h.Gorilla(enums.HandlerCreateTodo)).
		Methods(http.MethodPost)
	// TODO: fix GetTodo for Gorilla - now it is 404 at "/"
	// so I register at "/all" instead for debugging
	todosApi.
		Handle("/all", h.Gorilla(enums.HandlerGetTodo)).
		Methods(http.MethodGet)
	todosApi.
		Handle("/{uuid}", h.Gorilla(enums.HandlerGetTodo)).
		Methods(http.MethodGet)
	todosApi.
		Handle("/{uuid}", h.Gorilla(enums.HandlerUpdateTodo)).
		Methods(http.MethodPost)
	todosApi.
		Handle("/{uuid}", h.Gorilla(enums.HandlerDeleteTodo)).
		Methods(http.MethodDelete)
}
