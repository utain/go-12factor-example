package todos

import (
	"github.com/utain/go/example/internal/core/models"

	"github.com/google/uuid"
)

// Command use to insert todo
type AddTodoDto struct {
	Title       string
	Description string
}

// Command use to update todo status by id
type UpdateTodoStatusDto struct {
	ID     uuid.UUID
	Status models.TodoStatus
}

// Command use to delete todo by id
type DeleteTodoDto struct {
	ID uuid.UUID
}

// Query use to get first todo matched with id
type GetFirstDto struct {
	ID uuid.UUID
}

type FilterTodoStatus uint

const (
	FilterStatusAll FilterTodoStatus = iota
	FilterStatusPending
	FilterStatusDoing
	FilterStatusDone
)

// Query use to search and filter todo list
type FilterTodoDto struct {
	Title  string
	Status FilterTodoStatus
}
