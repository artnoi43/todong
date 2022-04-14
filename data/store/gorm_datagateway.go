package store

import (
	"context"
	"log"

	"gorm.io/gorm"

	"github.com/artnoi43/todong/data/model"
	"github.com/artnoi43/todong/domain/usecase"
)

type gormDataGateway struct {
	store GormStore
}

func NewGormDataGateway(g GormStore) usecase.DataGateway {
	return &gormDataGateway{
		store: g,
	}
}

func (g *gormDataGateway) GetUserByUsername(
	ctx context.Context,
	username string,
	dst *model.User,
) error {
	where := &model.User{
		Username: username,
	}
	return g.store.First(ctx, where, dst)
}

func (g *gormDataGateway) GetUserByUuid(
	ctx context.Context,
	uuid string,
	dst *model.User,
) error {
	where := &model.User{
		UUID: uuid,
	}
	return g.store.First(ctx, where, dst)
}

func (g *gormDataGateway) DeleteUser(
	ctx context.Context,
	where *model.User,
) error {
	return g.store.Delete(ctx, where, &model.User{})
}

func (g *gormDataGateway) DeleteTodo(
	ctx context.Context,
	where *model.Todo,
) error {
	return g.store.Delete(ctx, where, &model.Todo{})
}

func (g *gormDataGateway) CreateUser(
	ctx context.Context,
	user *model.User,
) error {
	return g.store.Create(ctx, user)
}

func (g *gormDataGateway) CreateTodo(
	ctx context.Context,
	todo *model.Todo,
) error {
	return g.store.Create(ctx, todo)
}

func (g *gormDataGateway) GetOneTodo(
	ctx context.Context,
	where *model.Todo,
	dst *model.Todo,
) error {
	return g.store.First(ctx, where, dst)
}

func (g *gormDataGateway) GetUserTodos(
	ctx context.Context,
	where *model.Todo,
	dst interface{},
) error {
	return g.store.Find(ctx, where, dst)
}

func (g *gormDataGateway) UpdateTodo(
	ctx context.Context,
	where *model.Todo,
	todo *model.Todo,
) error {
	return g.store.Updates(ctx, where, todo)
}

func (g *gormDataGateway) ChangePassword(
	ctx context.Context,
	where *model.User,
	hashedPassword []byte,
) error {
	return g.store.Updates(ctx, where, &model.User{
		Password: hashedPassword,
	})
}

func (g *gormDataGateway) Shutdown() {
	switch store := g.store.(type) {
	case *gormStore:
		switch db := store.db.(type) {
		case *gorm.DB:
			sqlDB, err := db.DB()
			if err != nil {
				log.Println("failed to get sqlDB")
			}
			sqlDB.Close()
		}
	}
}
