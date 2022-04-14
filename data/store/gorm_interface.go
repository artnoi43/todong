package store

import (
	"context"

	"gorm.io/gorm"
)

// GormDB represents *gorm.DB methods that GormStore wraps.
type GormDB interface {
	AutoMigrate(dst ...interface{}) error
	WithContext(context.Context) *gorm.DB
	Where(interface{}, ...interface{}) *gorm.DB
	First(interface{}, ...interface{}) *gorm.DB
	Find(interface{}, ...interface{}) *gorm.DB
	Create(interface{}) *gorm.DB
	Updates(interface{}) *gorm.DB
}

// GormStore wraps GormDB.
type GormStore interface {
	First(ctx context.Context, where interface{}, item interface{}) error
	Find(ctx context.Context, where interface{}, item interface{}) error
	Create(ctx context.Context, item interface{}) error
	Updates(ctx context.Context, where interface{}, item interface{}) error
	Delete(ctx context.Context, where interface{}, item interface{}) error
}

// gormStore is the actual implementation of GormStore.
type gormStore struct {
	db GormDB
}

func NewGormStore(db GormDB) GormStore {
	return &gormStore{
		db: db,
	}
}
