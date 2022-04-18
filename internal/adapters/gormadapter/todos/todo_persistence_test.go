package todos_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/utain/go/example/internal/adapters/gormadapter/entities"
	"github.com/utain/go/example/internal/adapters/gormadapter/todos"
	"github.com/utain/go/example/internal/core/models"
	domain "github.com/utain/go/example/internal/core/todos"
	"github.com/utain/go/example/internal/logs"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setup(t *testing.T) (domain.TodoPersistencePort, *gorm.DB) {
	db, err := gorm.Open(sqlite.Open("file::memory:"))
	if err != nil {
		assert.FailNow(t, "Can't init database connection")
	}
	db.AutoMigrate(&entities.TodoEntity{})
	return todos.WithGormPersistence(db, logs.Nolog), db
}

func TestInsertTodoWithGormAdapter(t *testing.T) {
	t.Run("should insert todo without error", func(t *testing.T) {
		persistence, _ := setup(t)
		expected := domain.AddTodoDto{
			Title:       "Todo Title",
			Description: "Todo Description",
		}
		actual, err := persistence.InsertTodo(context.Background(), expected)
		assert.Nil(t, err)
		assert.Equal(t, expected.Title, actual.Title)
		assert.Equal(t, expected.Description, actual.Description)
		assert.NotEqual(t, actual.ID, uuid.UUID{})
	})
}

func TestSearchTodoWithGormAdapter(t *testing.T) {
	t.Run("should get empty list", func(t *testing.T) {
		persistence, _ := setup(t)
		actual, err := persistence.FindByNameAndStatus(context.Background(), domain.FilterTodoDto{})
		assert.Nil(t, err)
		assert.Equal(t, 0, len(actual))
	})

	t.Run("should found 1 element in the list", func(t *testing.T) {
		persistence, db := setup(t)
		// prepare test todos
		todos := []entities.TodoEntity{
			{
				ID:          uuid.New(),
				Title:       "Todo Title",
				Description: "Todo Description",
				Status:      entities.TodoPendingStatus,
			},
		}
		if err := db.CreateInBatches(&todos, len(todos)).Error; err != nil {
			assert.FailNow(t, "Can't setup test data", err)
		}
		// execute test
		actual, err := persistence.FindByNameAndStatus(context.Background(), domain.FilterTodoDto{})
		assert.Nil(t, err)
		assert.Equal(t, 1, len(actual))
		assert.Equal(t, todos[0].ID, actual[0].ID)
		assert.Equal(t, todos[0].Title, actual[0].Title)
		assert.Equal(t, todos[0].Description, actual[0].Description)
		assert.Equal(t, models.StatusPending, actual[0].Status)
	})
}
