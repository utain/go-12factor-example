package mocktodos

import (
	"context"

	"github.com/utain/go/example/internal/core/models"
	"github.com/utain/go/example/internal/core/todos"

	"github.com/stretchr/testify/mock"
)

type MockTodoService struct {
	mock.Mock
}

func (m MockTodoService) GetByID(ctx context.Context, query todos.GetFirstDto) (*models.Todo, error) {
	args := m.Called(ctx, query)
	v, ok := args.Get(0).(*models.Todo)
	if !ok {
		return nil, args.Error(1)
	}
	return v, args.Error(1)
}
func (m MockTodoService) AddTodo(ctx context.Context, cmd todos.AddTodoDto) (*models.Todo, error) {
	args := m.Called(ctx, cmd)
	v, ok := args.Get(0).(*models.Todo)
	if !ok {
		return nil, args.Error(1)
	}
	return v, args.Error(1)
}
func (m MockTodoService) UpdateTodoStatus(ctx context.Context, cmd todos.UpdateTodoStatusDto) (*models.Todo, error) {
	args := m.Called(ctx, cmd)
	v, ok := args.Get(0).(*models.Todo)
	if !ok {
		return nil, args.Error(1)
	}
	return v, args.Error(1)
}
func (m MockTodoService) Filter(ctx context.Context, filter todos.FilterTodoDto) ([]models.Todo, error) {
	args := m.Called(ctx, filter)
	v, ok := args.Get(0).([]models.Todo)
	if !ok {
		return nil, args.Error(1)
	}
	return v, args.Error(1)
}
