package todos

import (
	"context"

	"github.com/utain/go/example/internal/core/models"
	"github.com/utain/go/example/internal/errs"
	"github.com/utain/go/example/internal/logs"

	"github.com/google/uuid"
)

func TodoService(log logs.Logging, repo TodoPersistencePort) TodoServicePort {
	return &todoService{
		repo: repo,
		log:  log,
	}
}

type todoService struct {
	log  logs.Logging
	repo TodoPersistencePort
}

func (s *todoService) AddTodo(ctx context.Context, cmd AddTodoDto) (out *models.Todo, err error) {
	if len(cmd.Title) == 0 {
		return nil, errs.ErrInvalidTodoTitle
	}
	if len(cmd.Description) == 0 {
		return nil, errs.ErrInvalidTodoDescription
	}
	out, err = s.repo.InsertTodo(ctx, cmd)
	if err != nil {
		return nil, errs.ErrCannotInsertTodo
	}
	return
}

func (s *todoService) GetByID(ctx context.Context, query GetFirstDto) (out *models.Todo, err error) {
	defer func() {
		if err != nil {
			s.log.Error("GetByID", logs.F{"error": err, "id": query.ID})
		}
	}()

	if query.ID == (uuid.UUID{}) {
		return nil, errs.ErrInvalidTodoID
	}

	out, err = s.repo.GetByID(ctx, query.ID)
	if out == nil || err != nil {
		return nil, errs.ErrTodoNotFound
	}
	return out, nil
}

func (s *todoService) UpdateTodoStatus(ctx context.Context, cmd UpdateTodoStatusDto) (out *models.Todo, err error) {
	if cmd.ID == uuid.Nil {
		return nil, errs.ErrInvalidTodoID
	}
	out, err = s.repo.UpdateTodoStatus(ctx, cmd)
	if out == nil || err != nil {
		return nil, errs.ErrTodoNotFound
	}
	return out, err
}

func (s *todoService) Filter(ctx context.Context, filter FilterTodoDto) (out []models.Todo, err error) {
	out, err = s.repo.FindByNameAndStatus(ctx, filter)
	if err != nil {
		return out, errs.ErrUnknown.With("filter", filter).With("error", err)
	}
	return out, nil
}
