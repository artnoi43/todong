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
		HandleFunc("/register", h.Gorilla(enums.HandlerRegister)).
		Methods(http.MethodPost)
	usersApi.
		HandleFunc("/login", h.Gorilla(enums.HandlerLogin)).
		Methods(http.MethodPost)
	protectedUsersApi.
		HandleFunc("/", h.Gorilla(enums.HandlerDeleteUser)).
		Methods(http.MethodDelete)
	protectedUsersApi.
		HandleFunc("/new-password", h.Gorilla(enums.HandlerNewPassword)).
		Methods(http.MethodPost)

	// /todos
	todosApi.
		HandleFunc("/", h.Gorilla(enums.HandlerGetTodo)).
		Methods(http.MethodGet)
	todosApi.
		HandleFunc("/create", h.Gorilla(enums.HandlerCreateTodo)).
		Methods(http.MethodPost)
	todosApi.
		HandleFunc("/{uuid}", h.Gorilla(enums.HandlerGetTodo)).
		Methods(http.MethodGet)
	todosApi.
		HandleFunc("/{uuid}", h.Gorilla(enums.HandlerUpdateTodo)).
		Methods(http.MethodPost)
	todosApi.
		HandleFunc("/{uuid}", h.Gorilla(enums.HandlerDeleteTodo)).
		Methods(http.MethodDelete)
}
