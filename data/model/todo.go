package model

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/artnoi43/todong/lib/enums"
)

// Todo represents to-dos data model in SQL databases
type Todo struct {
	UUID        string       `json:"uuid" gorm:"primaryKey;column:uuid"`
	UserUUID    string       `json:"user" gorm:"column:user_uuid"`
	Title       string       `json:"title" gorm:"column:title"`
	Description string       `json:"description" gorm:"column:description"`
	Status      enums.Status `json:"status" gorm:"column:status"`
	TodoDate    string       `json:"todoDate" gorm:"column:todo_date"` // time.Now().Format(time.RFC3339)
	CreatedAt   time.Time    `json:"createdAt" gorm:"autoUpdateTime;column:created_at"`
	Image       string       `json:"image" gorm:"column:image"`
}

// NewTodo returns new *model.Todo
func NewTodo(
	userUuid string,
	title string,
	description string,
	todoDate string,
	status enums.Status,
	image string,
) (*Todo, error) {
	if !status.IsValid() {
		return nil, fmt.Errorf("invalid status: %s", status)
	}
	return &Todo{
		UUID:        uuid.NewString(),
		UserUUID:    userUuid,
		Title:       title,
		Description: description,
		TodoDate:    todoDate,
		Status:      status,
		Image:       image,
	}, nil
}

func (t Todo) Marshal() string {
	b, _ := json.Marshal(t)
	return string(b)
}
