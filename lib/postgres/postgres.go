package postgres

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(conf *Config) (db *gorm.DB, err error) {
	// DB config
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s port=%s",
		conf.Host, conf.User, conf.Password, conf.Port)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	log.Println("Database successfully connected")

	// Performance config?
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}
