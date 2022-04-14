package usecase

import (
	"context"

	"github.com/artnoi43/todong/data/model"
)

// DataGateway abstracts high-level storage of datamode.User and model.Todo
type DataGateway interface {
	// Shutdown is needed for graceful shutdown
	Shutdown()

	CreateUser(
		ctx context.Context,
		user *model.User,
	) error
	CreateTodo(
		ctx context.Context,
		todo *model.Todo,
	) error
	GetUserByUuid(
		ctx context.Context,
		uuid string,
		dst *model.User,
	) error
	GetUserByUsername(
		ctx context.Context,
		username string,
		dst *model.User,
	) error
	GetOneTodo(
		ctx context.Context,
		where *model.Todo,
		dst *model.Todo,
	) error
	GetUserTodos(
		ctx context.Context,
		where *model.Todo,
		dst interface{},
	) error
	ChangePassword(
		ctx context.Context,
		where *model.User,
		hashedPassword []byte,
	) error
	UpdateTodo(
		ctx context.Context,
		where *model.Todo,
		todo *model.Todo,
	) error
	DeleteUser(
		ctx context.Context,
		where *model.User,
	) error
	DeleteTodo(
		ctx context.Context,
		where *model.Todo,
	) error
}
