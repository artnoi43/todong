package store

import (
	"context"

	"github.com/artnoi43/todong/data/model"
	"github.com/artnoi43/todong/domain/usecase"
)

type testDataGateway struct{}

func NewTestDataGateway() usecase.DataGateway { return &testDataGateway{} }

func (g *testDataGateway) DeleteUser(
	ctx context.Context,
	where *model.User,
) error {
	return nil
}

func (g *testDataGateway) DeleteTodo(
	ctx context.Context,
	where *model.Todo,
) error {
	return nil
}

func (g *testDataGateway) CreateUser(
	ctx context.Context,
	user *model.User,
) error {
	return nil
}

func (g *testDataGateway) CreateTodo(
	ctx context.Context,
	todo *model.Todo,
) error {
	return nil
}

// This function is called in handler.Login and handler.Register
// So calling this method will fail either of the tests.
func (g *testDataGateway) GetUserByUuid(
	ctx context.Context,
	uuid string,
	dst *model.User,
) error {
	return nil
}

func (g *testDataGateway) GetUserByUsername(
	ctx context.Context,
	username string,
	dst *model.User,
) error {
	return nil
}

func (g *testDataGateway) GetOneTodo(
	ctx context.Context,
	where *model.Todo,
	dst *model.Todo,
) error {
	return nil
}

func (g *testDataGateway) GetUserTodos(
	ctx context.Context,
	where *model.Todo,
	dst interface{},
) error {
	return nil
}

func (g *testDataGateway) UpdateTodo(
	ctx context.Context,
	where *model.Todo,
	todo *model.Todo,
) error {
	return nil
}

func (g *testDataGateway) ChangePassword(
	ctx context.Context,
	where *model.User,
	hashedPassword []byte,
) error {
	return nil
}

func (g *testDataGateway) Shutdown() {}
