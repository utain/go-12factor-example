package core

import (
	"github.com/utain/go/example/internal/core/todos"
	"github.com/utain/go/example/internal/core/todos/services"
	"github.com/utain/go/example/internal/logs"
)

type ServicesContainer struct {
	Log             logs.Logging
	TodoServicePort todos.TodoServicePort
}

func ServicesRegister(log logs.Logging, persistences PersistencesContainer) ServicesContainer {
	return ServicesContainer{
		Log:             log,
		TodoServicePort: services.TodoService(log, persistences.TodoPersistencePort),
	}
}
