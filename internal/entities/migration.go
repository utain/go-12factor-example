package entities

import (
	"go-example/internal/log"

	"gorm.io/gorm"
)

//AutoMigrate for migrate database schema
func AutoMigrate(db *gorm.DB) {
	log.Info("Migrating model")
	if err := db.AutoMigrate(&User{}, &Product{}, &ProductProps{}).Error; err != nil {
		log.Errorf("Can't automigrate schema %v", err)
	}
}
