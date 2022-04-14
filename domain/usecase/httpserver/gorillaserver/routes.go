package gorillaserver

import (
	"net/http"

	"github.com/artnoi43/todong/domain/usecase/handler"
	"github.com/artnoi43/todong/domain/usecase/middleware"
	"github.com/artnoi43/todong/lib/enums"
	"github.com/artnoi43/todong/lib/utils"
)

func (s *gorillaServer) SetUpRoutes(conf *middleware.Config, h handler.Adapter) {

	authenticator := utils.NewAuthenticator([]byte(conf.SecretKey))
	usersApi := s.router.PathPrefix("/users").Subrouter()
	todosApi := s.router.PathPrefix("/todos").Subrouter()
	protectedUsersApi := usersApi.NewRoute().Subrouter()
	protectedUsersApi.Use(authenticator.AuthMiddleware)
	todosApi.Use(authenticator.AuthMiddleware)

	// /users
	usersApi.
		HandleFunc("/register", h.NetHttp(enums.HandlerRegister)).
		Methods(http.MethodPost)
	usersApi.
		HandleFunc("/login", h.NetHttp(enums.HandlerLogin)).
		Methods(http.MethodPost)
	protectedUsersApi.
		HandleFunc("/", h.NetHttp(enums.HandlerDeleteUser)).
		Methods(http.MethodDelete)
	protectedUsersApi.
		HandleFunc("/new-password", h.NetHttp(enums.HandlerNewPassword)).
		Methods(http.MethodPost)

	// /todos
	todosApi.
		HandleFunc("/", h.NetHttp(enums.HandlerGetTodo)).
		Methods(http.MethodGet)
	todosApi.
		HandleFunc("/create", h.NetHttp(enums.HandlerCreateTodo)).
		Methods(http.MethodPost)
	todosApi.
		HandleFunc("/{uuid}", h.NetHttp(enums.HandlerGetTodo)).
		Methods(http.MethodGet)
	todosApi.
		HandleFunc("/{uuid}", h.NetHttp(enums.HandlerUpdateTodo)).
		Methods(http.MethodPost)
	todosApi.
		HandleFunc("/{uuid}", h.NetHttp(enums.HandlerDeleteTodo)).
		Methods(http.MethodDelete)
}
