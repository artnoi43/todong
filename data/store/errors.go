package store

import "github.com/pkg/errors"

var (
	// This error is returned when gorm.ErrRecordNotFound or redis.Nil is encountered
	ErrRecordNotFound = errors.New("record not found")
)
