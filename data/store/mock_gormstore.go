package store

// NOTE: might not be needed - we can use gomock to mock store.Store instead

import (
	"context"

	"github.com/artnoi43/todong/data/model"
	"github.com/artnoi43/todong/test"
)

type mockStore struct{}

func NewMockStore() GormStore {
	return &mockStore{}
}

func (m *mockStore) First(ctx context.Context, where interface{}, item interface{}) error {
	switch item := item.(type) {
	case []*model.Todo:
		switch where := where.(type) {
		case *model.Todo:
			item[0] = &model.Todo{
				UUID:     where.UUID,
				UserUUID: where.UserUUID,
			}
		}
	case *model.Todo:
		switch where := where.(type) {
		case *model.Todo:
			*item = model.Todo{
				UUID:     where.UUID,
				UserUUID: where.UserUUID,
			}
		}
	case *model.User:
		switch where := where.(type) {
		case *model.User:
			*item = model.User{
				UUID:     test.JwtIss,
				Username: where.Username,
				Password: test.HashedPW,
			}
		}
	}
	return nil
}
func (m *mockStore) Find(ctx context.Context, where interface{}, item interface{}) error {
	switch where := where.(type) {
	case *model.Todo:
		switch item := item.(type) {
		case *model.Todo:
			*item = model.Todo{
				UUID:     where.UUID,
				UserUUID: where.UserUUID,
			}
		case []*model.Todo:
			item[0] = &model.Todo{
				UUID:     where.UUID,
				UserUUID: where.UUID,
			}
		}
	case *model.User:
		switch item := item.(type) {
		case *model.User:
			*item = model.User{
				UUID:     test.JwtIss,
				Username: where.Username,
				Password: test.HashedPW,
			}
		}
	}
	return nil
}
func (m *mockStore) Create(ctx context.Context, item interface{}) error {
	return nil
}
func (m *mockStore) Updates(ctx context.Context, where interface{}, item interface{}) error {
	return nil
}
func (m *mockStore) Delete(ctx context.Context, where interface{}, item interface{}) error {
	return nil
}
