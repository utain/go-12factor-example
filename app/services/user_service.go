package services

import (
	"go-example/app/models"

	"github.com/jinzhu/gorm"
)

// NewUserService create userService
func NewUserService(db *gorm.DB) UserService {
	return &userService{db}
}

//UserService interface
type UserService interface {
	GetAllUser(offset int, limit int, search string) *[]models.User
	GetUser(name string) *models.User
}

// userService is a service private
type userService struct {
	db *gorm.DB
}

// GetAllUser return all User
func (p userService) GetAllUser(offset int, limit int, search string) *[]models.User {
	var users []models.User
	p.db.Find(&users)
	return &users
}

// GetUser return only one User
func (p userService) GetUser(name string) *models.User {
	var user models.User
	p.db.First(&user, &models.User{Username: name})
	return &user
}
