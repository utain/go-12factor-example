package todos

import (
	"context"
	"fmt"

	"github.com/utain/go/example/internal/adapters/gormadapter/entities"
	"github.com/utain/go/example/internal/core/models"
	"github.com/utain/go/example/internal/core/todos"
	"github.com/utain/go/example/internal/errs"
	"github.com/utain/go/example/internal/logs"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func WithGormPersistence(db *gorm.DB, log logs.Logging) todos.TodoPersistencePort {
	return &todoOrmRepository{db, log}
}

type todoOrmRepository struct {
	db  *gorm.DB
	log logs.Logging
}

func (r *todoOrmRepository) InsertTodo(ctx context.Context, cmd todos.AddTodoDto) (*models.Todo, error) {
	data := entities.TodoEntity{
		ID:          uuid.New(),
		Title:       cmd.Title,
		Description: cmd.Description,
		Status:      entities.TodoPendingStatus,
	}
	if err := r.db.WithContext(ctx).Create(&data).Error; err != nil {
		// should have no error
		return nil, err
	}
	out := data.ToDomain()
	return &out, nil
}

func (r *todoOrmRepository) UpdateTodoStatus(ctx context.Context, cmd todos.UpdateTodoStatusDto) (*models.Todo, error) {
	return nil, errs.ErrNotImplemented
}

// DeleteTodo(command DeleteTodoCommand) error {}
// GetAll(query QueryTodoList) ([]Todo, error){}
func (r *todoOrmRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Todo, error) {
	return nil, errs.ErrNotImplemented
}

func (r *todoOrmRepository) FindByNameAndStatus(ctx context.Context, query todos.FilterTodoDto) ([]models.Todo, error) {
	fmt.Println("Hello")
	out := []models.Todo{}
	list := []entities.TodoEntity{}
	if err := r.db.Debug().WithContext(ctx).Find(&list).Error; err != nil {
		// should have no error
		return out, err
	}
	for _, te := range list {
		out = append(out, te.ToDomain())
	}
	return out, nil
}
