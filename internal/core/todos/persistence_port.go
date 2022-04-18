package todos

import (
	"context"

	"github.com/utain/go/example/internal/core/models"

	"github.com/google/uuid"
)

type TodoPersistencePort interface {
	InsertTodo(ctx context.Context, cmd AddTodoDto) (*models.Todo, error)
	UpdateTodoStatus(ctx context.Context, cmd UpdateTodoStatusDto) (*models.Todo, error)
	// DeleteTodo(command DeleteTodoCommand) error
	// GetAll(query QueryTodoList) ([]Todo, error)
	GetByID(ctx context.Context, id uuid.UUID) (*models.Todo, error)
	FindByNameAndStatus(ctx context.Context, query FilterTodoDto) ([]models.Todo, error)
}
