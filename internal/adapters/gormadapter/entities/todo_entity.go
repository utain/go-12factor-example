package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/utain/go/example/internal/core/models"
	"gorm.io/gorm"
)

type TodoStatus string

func (t TodoStatus) ToDomain() models.TodoStatus {
	switch t {
	case TodoDoingStatus:
		return models.StatusDoing
	case TodoDoneStatus:
		return models.StatusDone
	case TodoPendingStatus:
		return models.StatusPending
	default:
		return models.StatusPending
	}
}

const (
	TodoPendingStatus TodoStatus = "PENDING"
	TodoDoingStatus   TodoStatus = "DOING"
	TodoDoneStatus    TodoStatus = "DONE"
)

type TodoEntity struct {
	ID          uuid.UUID `gorm:"primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Title       string
	Description string
	Status      TodoStatus
}

func (TodoEntity) TableName() string {
	return "todos"
}

func (t TodoEntity) ToDomain() models.Todo {
	return models.Todo{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		Status:      t.Status.ToDomain(),
	}
}
