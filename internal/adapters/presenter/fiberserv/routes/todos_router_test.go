package routes_test

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/utain/go/example/internal/adapters/presenter/fiberserv/middleware"
	"github.com/utain/go/example/internal/adapters/presenter/fiberserv/routes"
	"github.com/utain/go/example/internal/core"
	"github.com/utain/go/example/internal/core/models"
	"github.com/utain/go/example/internal/core/todos"
	"github.com/utain/go/example/internal/core/todos/mocktodos"
	"github.com/utain/go/example/internal/errs"
	"github.com/utain/go/example/internal/logs"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/utain/httpheaders"
	"github.com/utain/httpheaders/mediatypes"
	"github.com/utain/httpheaders/methods"
)

var matchedByContext = mock.MatchedBy(func(_ context.Context) bool {
	return true
})

func TestGetByID(t *testing.T) {
	// setup server
	serv := fiber.New(
		fiber.Config{
			ErrorHandler: middleware.FiberErrorMiddleware(middleware.ErrorMiddlewareOpts{Log: logs.Nolog}),
		},
	)
	apiRoutes := serv.Group("/api")
	mocked := &mocktodos.MockTodoService{}
	routes.TodoRouter(apiRoutes, core.ServicesContainer{
		Log:             logs.Nolog,
		TodoServicePort: mocked,
	})

	t.Run("should get todo with id without error", func(t *testing.T) {
		todoID := uuid.New()
		expected := models.Todo{
			ID:          todoID,
			Title:       "Todo",
			Description: "Todo Description",
			Status:      models.StatusDone,
		}
		mocked.On("GetByID", matchedByContext, todos.GetFirstDto{ID: todoID}).Return(&expected, nil)
		req := httptest.NewRequest(methods.GET, "/api/todos/"+todoID.String(), nil)
		res, err := serv.Test(req)
		assert.Nil(t, err)

		assert.Equal(t, http.StatusOK, res.StatusCode)
		out, err := ioutil.ReadAll(res.Body)
		assert.Nil(t, err)

		needJson, _ := json.Marshal(expected)
		assert.JSONEq(t, string(needJson), string(out))
	})

	t.Run("should get error when wrong uuid format", func(t *testing.T) {
		needErr := errs.ErrInvalidTodoID

		req := httptest.NewRequest(methods.GET, "/api/todos/1", nil)
		res, err := serv.Test(req)
		assert.Nil(t, err)

		assert.Equal(t, needErr.Code, res.StatusCode)

		out, err := ioutil.ReadAll(res.Body)
		assert.Nil(t, err)

		assert.JSONEq(t, needErr.String(), string(out))
	})

	t.Run("should get error when server layers error", func(t *testing.T) {
		id := uuid.New()
		needErr := errs.ErrTodoNotFound
		mocked.On("GetByID", matchedByContext, todos.GetFirstDto{ID: id}).Return(nil, needErr)

		req := httptest.NewRequest(methods.GET, "/api/todos/"+id.String(), nil)
		res, err := serv.Test(req)
		assert.Nil(t, err)

		assert.Equal(t, needErr.Code, res.StatusCode)

		out, err := ioutil.ReadAll(res.Body)
		assert.Nil(t, err)

		assert.JSONEq(t, needErr.String(), string(out))
	})
}

func TestCreateTodo(t *testing.T) {
	serv := fiber.New(
		fiber.Config{
			ErrorHandler: middleware.FiberErrorMiddleware(middleware.ErrorMiddlewareOpts{Log: logs.Nolog}),
		},
	)
	apiRoutes := serv.Group("/api")
	mock := &mocktodos.MockTodoService{}
	routes.TodoRouter(apiRoutes, core.ServicesContainer{
		Log:             logs.Nolog,
		TodoServicePort: mock,
	})

	t.Run("should create todo with error", func(t *testing.T) {
		needErr := errs.ErrNotImplemented
		reader := strings.NewReader(`
			"title": "title",
			"description": "description"
		`)
		req := httptest.NewRequest(methods.POST, "/api/todos", reader)
		req.Header.Add(httpheaders.ContentType, mediatypes.ApplicationJson)
		res, err := serv.Test(req)
		assert.Nil(t, err)

		assert.Equal(t, needErr.Code, res.StatusCode)

		out, err := ioutil.ReadAll(res.Body)
		assert.Nil(t, err)
		assert.JSONEq(t, needErr.String(), string(out))
	})
}

func TestDeleteTodo(t *testing.T) {
	serv := fiber.New(
		fiber.Config{
			ErrorHandler: middleware.FiberErrorMiddleware(middleware.ErrorMiddlewareOpts{Log: logs.Nolog}),
		},
	)
	apiRoutes := serv.Group("/api")
	mock := &mocktodos.MockTodoService{}
	routes.TodoRouter(apiRoutes, core.ServicesContainer{
		Log:             logs.Nolog,
		TodoServicePort: mock,
	})

	t.Run("should delete todo with error", func(t *testing.T) {
		needErr := errs.ErrNotImplemented

		req := httptest.NewRequest(methods.DELETE, "/api/todos/1", nil)
		req.Header.Add(httpheaders.ContentType, mediatypes.ApplicationJson)
		res, err := serv.Test(req)
		assert.Nil(t, err)
		assert.Equal(t, needErr.Code, res.StatusCode)

		out, err := ioutil.ReadAll(res.Body)
		assert.Nil(t, err)
		assert.JSONEq(t, needErr.String(), string(out))
	})
}
func TestSearchTodoList(t *testing.T) {
	serv := fiber.New(
		fiber.Config{
			ErrorHandler: middleware.FiberErrorMiddleware(middleware.ErrorMiddlewareOpts{Log: logs.Nolog}),
		},
	)
	apiRoutes := serv.Group("/api")
	mock := &mocktodos.MockTodoService{}
	routes.TodoRouter(apiRoutes, core.ServicesContainer{
		Log:             logs.Nolog,
		TodoServicePort: mock,
	})

	t.Run("should search todos without error", func(t *testing.T) {
		expected := []models.Todo{}
		mock.On("Filter", matchedByContext, todos.FilterTodoDto{}).Return(expected, nil)
		req := httptest.NewRequest(methods.GET, "/api/todos?s=", nil)
		res, err := serv.Test(req)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, res.StatusCode)

		out, err := ioutil.ReadAll(res.Body)
		assert.Nil(t, err)
		assert.JSONEq(t, "[]", string(out))
	})
}
func TestUpdateTodo(t *testing.T) {
	serv := fiber.New(
		fiber.Config{
			ErrorHandler: middleware.FiberErrorMiddleware(middleware.ErrorMiddlewareOpts{Log: logs.Nolog}),
		},
	)
	apiRoutes := serv.Group("/api")
	mock := &mocktodos.MockTodoService{}
	routes.TodoRouter(apiRoutes, core.ServicesContainer{
		Log:             logs.Nolog,
		TodoServicePort: mock,
	})

	t.Run("should update todo with error", func(t *testing.T) {
		needErr := errs.ErrNotImplemented
		req := httptest.NewRequest(methods.PATCH, "/api/todos/1", nil)
		req.Header.Add(httpheaders.ContentType, mediatypes.ApplicationJson)
		res, err := serv.Test(req)
		assert.Nil(t, err)
		assert.Equal(t, needErr.Code, res.StatusCode)

		out, err := ioutil.ReadAll(res.Body)
		assert.Nil(t, err)
		assert.JSONEq(t, needErr.String(), string(out))
	})

	t.Run("should update todo status with error", func(t *testing.T) {
		needErr := errs.ErrNotImplemented
		req := httptest.NewRequest(methods.PATCH, "/api/todos/1/status", nil)
		req.Header.Add(httpheaders.ContentType, mediatypes.ApplicationJson)
		res, err := serv.Test(req)
		assert.Nil(t, err)
		assert.Equal(t, needErr.Code, res.StatusCode)

		out, err := ioutil.ReadAll(res.Body)
		assert.Nil(t, err)
		assert.JSONEq(t, needErr.String(), string(out))
	})
}
