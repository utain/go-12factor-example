package models

import (
	"go-example/log"
	"time"

	"github.com/jinzhu/gorm"
)

//AutoMigrate for migrate database schema
func AutoMigrate(db *gorm.DB) {
	log.Debug("Migrating User model")
	if err := db.AutoMigrate(&User{}, &Product{}, &ProductProps{}).Error; err != nil {
		log.Debugf("Can't automigrate schema %v", err)
	}
	InitData(db)
}

//InitData using for initial data when first time install
func InitData(db *gorm.DB) {
	db.Create(&User{
		Model:     Model{ID: "1"},
		Username:  "Example",
		Email:     "user@example.com",
		Firstname: "Example",
		Lastname:  "E",
	})
	db.Create(&Product{
		Model: Model{ID: "1"}, Name: "MacBook Pro 13 2019", Code: "mbp132019", Price: 2999, Attr: AttrType{
			"with":   "23m",
			"height": "32m",
		},
		Props: []ProductProps{
			ProductProps{Model: Model{ID: "1"}, Key: "Name", Value: "MacDev"},
		}})
	time.Sleep(time.Second)

	prod := &Product{}
	db.Preload("Props").First(prod)
	log.Debug("Product:", prod)
}
