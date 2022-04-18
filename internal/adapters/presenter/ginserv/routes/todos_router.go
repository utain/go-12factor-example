package routes

import (
	"fmt"
	"net/http"

	"github.com/utain/go/example/internal/core/todos"
	"github.com/utain/go/example/internal/errs"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func TodoRouter(router *gin.RouterGroup, port todos.TodoServicePort) {
	todoCtl := todoCtl{port}
	// query
	router.GET("/todos/:id", todoCtl.GetByID)
	router.GET("/todos", todoCtl.SearchTodos)

	// mutations
	router.POST("/todos", todoCtl.CreateTodo)
	router.PATCH("/todos/:id", todoCtl.UpdateInfo)
	router.PATCH("/todos/:id/status", todoCtl.UpdateTodoStatus)
	router.DELETE("/todos/:id", todoCtl.DeleteTodo)
}

type todoCtl struct {
	service todos.TodoServicePort
}

// query
func (ctl *todoCtl) GetByID(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.Error(errs.ErrInvalidTodoID)
		return
	}
	out, err := ctl.service.GetByID(ctx.Request.Context(), todos.GetFirstDto{ID: id})
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, out)
}

func (ctl *todoCtl) SearchTodos(ctx *gin.Context) {
	fmt.Println("Hello1x")
	search := ctx.Query("s")
	out, err := ctl.service.Filter(ctx.Request.Context(), todos.FilterTodoDto{Status: todos.FilterStatusAll, Title: search})
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, out)
}

// mutation
func (ctl *todoCtl) CreateTodo(ctx *gin.Context) {
	ctx.Error(errs.ErrNotImplemented)
}
func (ctl *todoCtl) UpdateInfo(ctx *gin.Context) {
	ctx.Error(errs.ErrNotImplemented)
}
func (ctl *todoCtl) UpdateTodoStatus(ctx *gin.Context) {
	ctx.Error(errs.ErrNotImplemented)
}
func (ctl *todoCtl) DeleteTodo(ctx *gin.Context) {
	ctx.Error(errs.ErrNotImplemented)
}
