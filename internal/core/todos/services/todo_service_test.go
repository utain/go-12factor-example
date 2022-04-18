package services_test

import (
	"context"
	"testing"

	"github.com/utain/go/example/internal/adapters/zapadapter"
	"github.com/utain/go/example/internal/core/models"
	"github.com/utain/go/example/internal/core/todos"
	"github.com/utain/go/example/internal/core/todos/services"
	"github.com/utain/go/example/internal/errs"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTodoPersistence struct {
	mock.Mock
}

func (m MockTodoPersistence) InsertTodo(ctx context.Context, cmd todos.AddTodoDto) (*models.Todo, error) {
	args := m.Called(ctx, cmd)
	v, ok := args.Get(0).(*models.Todo)
	if !ok {
		return nil, args.Error(1)
	}
	return v, args.Error(1)
}

func (m MockTodoPersistence) GetByID(ctx context.Context, id uuid.UUID) (*models.Todo, error) {
	args := m.Called(ctx, id)
	v, ok := args.Get(0).(*models.Todo)
	if !ok {
		return nil, args.Error(1)
	}
	return v, args.Error(1)
}

func (m MockTodoPersistence) UpdateTodoStatus(ctx context.Context, cmd todos.UpdateTodoStatusDto) (*models.Todo, error) {
	args := m.Called(ctx, cmd)
	v, ok := args.Get(0).(*models.Todo)
	if !ok {
		return nil, args.Error(1)
	}
	return v, args.Error(1)
}

func (m MockTodoPersistence) FindByNameAndStatus(ctx context.Context, filter todos.FilterTodoDto) ([]models.Todo, error) {
	args := m.Called(ctx, filter)
	v, ok := args.Get(0).([]models.Todo)
	if !ok {
		return nil, args.Error(1)
	}
	return v, args.Error(1)
}

func TestAddTodo(t *testing.T) {
	zaplog := zapadapter.ZapAdapter()
	t.Run("should error when todo title is empty", func(t *testing.T) {
		mock := &MockTodoPersistence{}
		command := todos.AddTodoDto{}
		serv := services.TodoService(zaplog, mock)

		out, err := serv.AddTodo(context.Background(), command)
		assert.Equal(t, errs.ErrInvalidTodoTitle, err)
		assert.Nil(t, out)
	})

	t.Run("should error when todo title is empty and description provided", func(t *testing.T) {
		mock := &MockTodoPersistence{}
		command := todos.AddTodoDto{Description: "Description"}
		serv := services.TodoService(zaplog, mock)

		out, err := serv.AddTodo(context.Background(), command)
		assert.Equal(t, errs.ErrInvalidTodoTitle, err)
		assert.Nil(t, out)
	})

	t.Run("should error when todo description is empty", func(t *testing.T) {
		mock := &MockTodoPersistence{}
		command := todos.AddTodoDto{Title: "Title"}
		serv := services.TodoService(zaplog, mock)

		out, err := serv.AddTodo(context.Background(), command)
		assert.Equal(t, errs.ErrInvalidTodoDescription, err)
		assert.Nil(t, out)
	})

	t.Run("should error when persistence error", func(t *testing.T) {
		mock := &MockTodoPersistence{}
		command := todos.AddTodoDto{Title: "Title", Description: "Description"}
		mock.On("InsertTodo", context.Background(), command).Return(nil, errs.ErrCannotInsertTodo)
		serv := services.TodoService(zaplog, mock)

		out, err := serv.AddTodo(context.Background(), command)
		assert.Equal(t, errs.ErrCannotInsertTodo, err)
		assert.Nil(t, out)
	})

	t.Run("should add todo without error", func(t *testing.T) {
		mock := &MockTodoPersistence{}
		command := todos.AddTodoDto{Title: "Title", Description: "Description"}
		need := &models.Todo{
			ID:          uuid.New(),
			Title:       "Title",
			Description: "Description",
			Status:      models.StatusPending,
		}
		mock.On("InsertTodo", context.Background(), command).Return(need, nil)
		serv := services.TodoService(zaplog, mock)

		out, err := serv.AddTodo(context.Background(), command)
		assert.Nil(t, err)
		assert.Equal(t, need, out)
	})
}

func TestGetTodoByID(t *testing.T) {
	zaplog := zapadapter.ZapAdapter()
	t.Run("should got error when id is empty", func(t *testing.T) {
		mock := &MockTodoPersistence{}
		serv := services.TodoService(zaplog, mock)
		query := todos.GetFirstDto{}
		out, err := serv.GetByID(context.Background(), query)
		assert.Equal(t, errs.ErrInvalidTodoID, err)
		assert.Nil(t, out)
	})

	t.Run("should got error when data not found", func(t *testing.T) {
		mock := &MockTodoPersistence{}
		query := todos.GetFirstDto{
			ID: uuid.New(),
		}
		mock.On("GetByID", context.Background(), query.ID).Return(nil, nil)
		serv := services.TodoService(zaplog, mock)

		out, err := serv.GetByID(context.Background(), query)
		assert.Equal(t, errs.ErrTodoNotFound, err)
		assert.Nil(t, out)
	})

	t.Run("should got data without error", func(t *testing.T) {
		mock := &MockTodoPersistence{}
		query := todos.GetFirstDto{
			ID: uuid.New(),
		}
		need := &models.Todo{
			ID:          query.ID,
			Title:       "Title",
			Description: "Description",
			Status:      models.StatusPending,
		}
		mock.On("GetByID", context.Background(), query.ID).Return(need, nil)
		serv := services.TodoService(zaplog, mock)

		out, err := serv.GetByID(context.Background(), query)
		assert.Equal(t, need, out)
		assert.Nil(t, err)
	})
}

func TestUpdateTodo(t *testing.T) {
	zaplog := zapadapter.ZapAdapter()
	t.Run("should error when todo id is empty", func(t *testing.T) {
		mock := &MockTodoPersistence{}
		serv := services.TodoService(zaplog, mock)

		out, err := serv.UpdateTodoStatus(context.Background(), todos.UpdateTodoStatusDto{})
		assert.Equal(t, errs.ErrInvalidTodoID, err)
		assert.Nil(t, out)
	})

	t.Run("should error when todo is missing", func(t *testing.T) {
		mock := &MockTodoPersistence{}
		serv := services.TodoService(zaplog, mock)
		cmd := todos.UpdateTodoStatusDto{
			ID:     uuid.New(),
			Status: models.StatusDoing,
		}
		mock.On("UpdateTodoStatus", context.Background(), cmd).Return(nil, nil)

		out, err := serv.UpdateTodoStatus(context.Background(), cmd)
		assert.Equal(t, errs.ErrTodoNotFound, err)
		assert.Nil(t, out)
	})

	t.Run("should update without any error", func(t *testing.T) {
		mock := &MockTodoPersistence{}
		serv := services.TodoService(zaplog, mock)
		status := models.StatusDoing
		cmd := todos.UpdateTodoStatusDto{
			ID:     uuid.New(),
			Status: status,
		}
		need := &models.Todo{
			ID:          uuid.New(),
			Title:       "Title",
			Description: "Description",
			Status:      status,
		}
		mock.On("UpdateTodoStatus", context.Background(), cmd).Return(need, nil)

		out, err := serv.UpdateTodoStatus(context.Background(), cmd)
		assert.Equal(t, need, out)
		assert.Nil(t, err)
	})
}

func TestFilterTodo(t *testing.T) {
	zaplog := zapadapter.ZapAdapter()
	t.Run("should show all when no filter", func(t *testing.T) {
		mock := &MockTodoPersistence{}
		serv := services.TodoService(zaplog, mock)
		filter := todos.FilterTodoDto{}
		need := []models.Todo{}
		mock.On("FindByNameAndStatus", context.Background(), filter).Return(need, nil)

		out, err := serv.Filter(context.Background(), filter)
		assert.Equal(t, need, out)
		assert.Nil(t, err)
	})

	t.Run("should error when persistence error", func(t *testing.T) {
		mock := &MockTodoPersistence{}
		serv := services.TodoService(zaplog, mock)
		filter := todos.FilterTodoDto{}
		need := []models.Todo{}
		mock.On("FindByNameAndStatus", context.Background(), filter).Return(need, nil)

		out, err := serv.Filter(context.Background(), filter)
		assert.Equal(t, need, out)
		assert.Error(t, errs.ErrUnknown, err)
	})
}
