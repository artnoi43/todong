package model

import (
	"time"
)

// User represents users' data model in SQL databases
type User struct {
	UUID      string    `json:"uuid" gorm:"primaryKey;column:uuid"`
	Username  string    `json:"username" gorm:"column:username"`
	Password  []byte    `json:"-" gorm:"column:password"`
	Todos     []Todo    `json:"todos" gorm:"foreignKey:UserUUID;references:UUID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt time.Time `json:"-" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"-" gorm:"autoUpdateTime"`
}
