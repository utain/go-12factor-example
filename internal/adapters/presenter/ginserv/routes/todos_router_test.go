package routes_test

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/utain/go/example/internal/adapters/presenter/ginserv/middleware"
	"github.com/utain/go/example/internal/adapters/presenter/ginserv/routes"
	"github.com/utain/go/example/internal/core/models"
	"github.com/utain/go/example/internal/core/todos"
	"github.com/utain/go/example/internal/core/todos/mocktodos"
	"github.com/utain/go/example/internal/errs"
	"github.com/utain/go/example/internal/logs"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/utain/httpheaders"
	"github.com/utain/httpheaders/mediatypes"
	"github.com/utain/httpheaders/methods"
)

func TestGetByID(t *testing.T) {
	// setup server
	serv := gin.New()
	serv.Use(middleware.GinErrorMiddleware(middleware.ErrorOptions{Log: logs.Nolog}))
	apiRoutes := serv.Group("/api")
	mock := &mocktodos.MockTodoService{}
	routes.TodoRouter(apiRoutes, mock)

	t.Run("should get todo with id without error", func(t *testing.T) {
		todoID := uuid.New()
		expected := models.Todo{
			ID:          todoID,
			Title:       "Todo",
			Description: "Todo Description",
			Status:      models.StatusDone,
		}
		mock.On("GetByID", context.Background(), todos.GetFirstDto{ID: todoID}).Return(&expected, nil)
		res := httptest.NewRecorder()
		req := httptest.NewRequest(methods.GET, "/api/todos/"+todoID.String(), nil)
		serv.ServeHTTP(res, req)

		assert.Equal(t, http.StatusOK, res.Result().StatusCode)
		out, err := ioutil.ReadAll(res.Body)
		assert.Nil(t, err)

		needJson, _ := json.Marshal(expected)
		assert.JSONEq(t, string(needJson), string(out))
	})

	t.Run("should get error when wrong uuid format", func(t *testing.T) {
		needErr := errs.ErrInvalidTodoID
		res := httptest.NewRecorder()
		req := httptest.NewRequest(methods.GET, "/api/todos/1", nil)
		serv.ServeHTTP(res, req)

		assert.Equal(t, needErr.Code, res.Result().StatusCode)

		out, err := ioutil.ReadAll(res.Body)
		assert.Nil(t, err)

		assert.JSONEq(t, needErr.String(), string(out))
	})

	t.Run("should get error when server layers error", func(t *testing.T) {
		id := uuid.New()
		needErr := errs.ErrTodoNotFound
		mock.On("GetByID", context.Background(), todos.GetFirstDto{ID: id}).Return(nil, needErr)
		res := httptest.NewRecorder()
		req := httptest.NewRequest(methods.GET, "/api/todos/"+id.String(), nil)
		serv.ServeHTTP(res, req)

		assert.Equal(t, needErr.Code, res.Result().StatusCode)

		out, err := ioutil.ReadAll(res.Body)
		assert.Nil(t, err)

		assert.JSONEq(t, needErr.String(), string(out))
	})
}

func TestCreateTodo(t *testing.T) {
	serv := gin.New()
	serv.Use(middleware.GinErrorMiddleware(middleware.ErrorOptions{Log: logs.Nolog}))
	apiRoutes := serv.Group("/api")
	mock := &mocktodos.MockTodoService{}
	routes.TodoRouter(apiRoutes, mock)

	t.Run("should create todo with error", func(t *testing.T) {
		needErr := errs.ErrNotImplemented
		reader := strings.NewReader(`
			"title": "title",
			"description": "description"
		`)
		res := httptest.NewRecorder()
		req := httptest.NewRequest(methods.POST, "/api/todos", reader)
		req.Header.Add(httpheaders.ContentType, mediatypes.ApplicationJson)
		serv.ServeHTTP(res, req)
		assert.Equal(t, needErr.Code, res.Result().StatusCode)

		out, err := ioutil.ReadAll(res.Body)
		assert.Nil(t, err)
		assert.JSONEq(t, needErr.String(), string(out))
	})
}

func TestDeleteTodo(t *testing.T) {
	serv := gin.New()
	serv.Use(middleware.GinErrorMiddleware(middleware.ErrorOptions{Log: logs.Nolog}))
	apiRoutes := serv.Group("/api")
	mock := &mocktodos.MockTodoService{}
	routes.TodoRouter(apiRoutes, mock)

	t.Run("should delete todo with error", func(t *testing.T) {
		needErr := errs.ErrNotImplemented

		res := httptest.NewRecorder()
		req := httptest.NewRequest(methods.DELETE, "/api/todos/1", nil)
		req.Header.Add(httpheaders.ContentType, mediatypes.ApplicationJson)
		serv.ServeHTTP(res, req)
		assert.Equal(t, needErr.Code, res.Result().StatusCode)

		out, err := ioutil.ReadAll(res.Body)
		assert.Nil(t, err)
		assert.JSONEq(t, needErr.String(), string(out))
	})
}
func TestUpdateTodo(t *testing.T) {
	serv := gin.New()
	serv.Use(middleware.GinErrorMiddleware(middleware.ErrorOptions{Log: logs.Nolog}))
	apiRoutes := serv.Group("/api")
	mock := &mocktodos.MockTodoService{}
	routes.TodoRouter(apiRoutes, mock)

	t.Run("should update todo with error", func(t *testing.T) {
		needErr := errs.ErrNotImplemented
		res := httptest.NewRecorder()
		req := httptest.NewRequest(methods.PATCH, "/api/todos/1", nil)
		req.Header.Add(httpheaders.ContentType, mediatypes.ApplicationJson)
		serv.ServeHTTP(res, req)
		assert.Equal(t, needErr.Code, res.Result().StatusCode)

		out, err := ioutil.ReadAll(res.Body)
		assert.Nil(t, err)
		assert.JSONEq(t, needErr.String(), string(out))
	})

	t.Run("should update todo status with error", func(t *testing.T) {
		needErr := errs.ErrNotImplemented
		res := httptest.NewRecorder()
		req := httptest.NewRequest(methods.PATCH, "/api/todos/1/status", nil)
		req.Header.Add(httpheaders.ContentType, mediatypes.ApplicationJson)
		serv.ServeHTTP(res, req)
		assert.Equal(t, needErr.Code, res.Result().StatusCode)

		out, err := ioutil.ReadAll(res.Body)
		assert.Nil(t, err)
		assert.JSONEq(t, needErr.String(), string(out))
	})
}
