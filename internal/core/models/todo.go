package models

import "github.com/google/uuid"

type TodoStatus uint

const (
	StatusPending TodoStatus = iota
	StatusDoing
	StatusDone
)

func (s TodoStatus) String() string {
	switch s {
	case StatusDoing:
		return "DOING"
	case StatusDone:
		return "DONE"
	default:
		return "PENDING"
	}
}

type Todo struct {
	ID          uuid.UUID  `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      TodoStatus `json:"status"`
}
