package routes

import (
	"fmt"

	"github.com/utain/go/example/internal/core"
	"github.com/utain/go/example/internal/core/todos"
	"github.com/utain/go/example/internal/errs"
	"github.com/utain/go/example/internal/logs"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func TodoRouter(route fiber.Router, services core.ServicesContainer) {
	handler := todosRouter{
		log:     services.Log,
		service: services.TodoServicePort,
	}
	// query
	route.Get("/todos/:id", handler.GetByID)
	route.Get("/todos", handler.Search)

	// mutation
	route.Post("/todos", handler.Create)
	route.Patch("/todos/:id", handler.UpdateInfo)
	route.Patch("/todos/:id/status", handler.UpdateStatus)
	route.Delete("/todos/:id", handler.Delete)
}

type todosRouter struct {
	log     logs.Logging
	service todos.TodoServicePort
}

func (r *todosRouter) GetByID(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return errs.ErrInvalidTodoID
	}

	data, err := r.service.GetByID(ctx.Context(), todos.GetFirstDto{ID: id})
	if err != nil {
		r.log.Error("Call service got error", logs.F{"error": err})
		return err
	}
	return ctx.JSON(data)
}

func (r *todosRouter) Search(ctx *fiber.Ctx) error {
	fmt.Println("Hello1")
	data, err := r.service.Filter(ctx.Context(), todos.FilterTodoDto{})
	if err != nil {
		r.log.Error("Bug when search todos", logs.F{"error": err})
		return err
	}
	return ctx.JSON(data)
}
func (r *todosRouter) Create(ctx *fiber.Ctx) error {
	return errs.ErrNotImplemented
}
func (r *todosRouter) UpdateInfo(ctx *fiber.Ctx) error {
	return errs.ErrNotImplemented
}
func (r *todosRouter) UpdateStatus(ctx *fiber.Ctx) error {
	return errs.ErrNotImplemented
}
func (r *todosRouter) Delete(ctx *fiber.Ctx) error {
	return errs.ErrNotImplemented
}
