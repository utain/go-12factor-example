package todos

import (
	"context"

	"github.com/utain/go/example/internal/core/models"
)

type TodoServicePort interface {
	AddTodo(ctx context.Context, cmd AddTodoDto) (*models.Todo, error)
	GetByID(ctx context.Context, query GetFirstDto) (*models.Todo, error)
	UpdateTodoStatus(ctx context.Context, cmd UpdateTodoStatusDto) (*models.Todo, error)
	Filter(ctx context.Context, filter FilterTodoDto) ([]models.Todo, error)
}
