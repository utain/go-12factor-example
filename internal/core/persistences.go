package core

import "github.com/utain/go/example/internal/core/todos"

type PersistencesContainer struct {
	TodoPersistencePort todos.TodoPersistencePort
}
