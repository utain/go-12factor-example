package gormadapter

import (
	"github.com/utain/go/example/internal/adapters/gormadapter/entities"
	"gorm.io/gorm"
)

func Migration(db *gorm.DB) (err error) {
	return db.AutoMigrate(&entities.TodoEntity{})
}
