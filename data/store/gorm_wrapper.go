package store

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

func handleGormNotFound(err error) error {
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrRecordNotFound
	}
	return err
}

// First item with conditions
func (s *gormStore) First(
	ctx context.Context,
	where interface{},
	item interface{},
) error {
	tx := s.db.WithContext(ctx).Where(where).First(item)
	return handleGormNotFound(tx.Error)
}

// Find items with conditions
func (s *gormStore) Find(
	ctx context.Context,
	where interface{},
	item interface{},
) error {
	tx := s.db.WithContext(ctx).Where(where).Find(item)
	return handleGormNotFound(tx.Error)
}

// Create items
func (s *gormStore) Create(
	ctx context.Context,
	item interface{},
) error {
	tx := s.db.WithContext(ctx).Create(item)
	return handleGormNotFound(tx.Error)
}

// Update item with conditions
func (s *gormStore) Updates(
	ctx context.Context,
	where interface{},
	item interface{},
) error {
	tx := s.db.WithContext(ctx).Where(where).Updates(item)
	return handleGormNotFound(tx.Error)
}

// Delete item with conditions
func (s *gormStore) Delete(
	ctx context.Context,
	where interface{},
	item interface{},
) error {
	tx := s.db.WithContext(ctx).Where(where).Delete(item)
	return handleGormNotFound(tx.Error)
}
