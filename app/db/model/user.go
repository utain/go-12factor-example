package model

import (
	"github.com/jinzhu/gorm"
)

// User model
type User struct {
	gorm.Model
	Username  string
	Email     string
	Password  string
	Firstname string
	Lastname  string
}
